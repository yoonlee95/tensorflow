package predict

// #cgo LDFLAGS: -ltensorflow
// #cgo CFLAGS: -I${SRCDIR}/../../../tensorflow/tensorflow
// #include "tensorflow/c/c_api.h"
import "C"

// #include "tensorflow/core/framework/graph.pb.h"
// #include <stdlib.h>
// #include <string.h>
// void setDefaultDevice(const char * device, TF_Graph * graph_def) {
//  tensorflow::GraphDef def;
// GetGraphDef(graph_def, &def);
// graph::SetDefaultDevice(device, def);
// 	}
import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unsafe"

	"github.com/k0kubun/pp"
	opentracing "github.com/opentracing/opentracing-go"
	olog "github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"github.com/rai-project/config"
	"github.com/rai-project/dlframework"
	"github.com/rai-project/dlframework/framework/agent"
	"github.com/rai-project/dlframework/framework/options"
	common "github.com/rai-project/dlframework/framework/predict"
	"github.com/rai-project/downloadmanager"
	"github.com/rai-project/image"
	"github.com/rai-project/image/types"
	nvidiasmi "github.com/rai-project/nvidia-smi"
	"github.com/rai-project/tensorflow"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	context "golang.org/x/net/context"
)

type ImagePredictor struct {
	common.ImagePredictor
	features    []string
	tfGraph     *tf.Graph
	tfSession   *tf.Session
	inputLayer  string
	outputLayer string
}

func New(model dlframework.ModelManifest, opts ...options.Option) (common.Predictor, error) {
	modelInputs := model.GetInputs()
	if len(modelInputs) != 1 {
		return nil, errors.New("number of inputs not supported")
	}
	firstInputType := modelInputs[0].GetType()
	if strings.ToLower(firstInputType) != "image" {
		return nil, errors.New("input type not supported")
	}

	predictor := new(ImagePredictor)

	return predictor.Load(context.Background(), model, opts...)
}

func (p *ImagePredictor) Load(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) (common.Predictor, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Load")
	defer span.Finish()

	framework, err := model.ResolveFramework()
	if err != nil {
		return nil, err
	}

	workDir, err := model.WorkDir()
	if err != nil {
		return nil, err
	}

	ip := &ImagePredictor{
		ImagePredictor: common.ImagePredictor{
			Base: common.Base{
				Framework: framework,
				Model:     model,
				Options:   options.New(opts...),
			},
			WorkDir: workDir,
		},
	}

	if err = ip.download(ctx); err != nil {
		return nil, err
	}

	if err = ip.loadPredictor(ctx); err != nil {
		return nil, err
	}

	return ip, nil
}

func (p *ImagePredictor) GetPreprocessOptions(ctx context.Context) (common.PreprocessOptions, error) {
	mean, err := p.GetMeanImage()
	if err != nil {
		return common.PreprocessOptions{}, err
	}

	scale, err := p.GetScale()
	if err != nil {
		return common.PreprocessOptions{}, err
	}

	imageDims, err := p.GetImageDimensions()
	if err != nil {
		return common.PreprocessOptions{}, err
	}

	return common.PreprocessOptions{
		MeanImage: mean,
		Scale:     scale,
		Size:      []int{int(imageDims[1]), int(imageDims[2])},
		ColorMode: types.RGBMode,
		Layout:    image.HWCLayout,
	}, nil
}

func (p *ImagePredictor) download(ctx context.Context) error {
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"Download",
		opentracing.Tags{
			"graph_url":           p.GetGraphUrl(),
			"target_graph_file":   p.GetGraphPath(),
			"weights_url":         p.GetWeightsUrl(),
			"target_weights_file": p.GetWeightsPath(),
			"feature_url":         p.GetFeaturesUrl(),
			"target_feature_file": p.GetFeaturesPath(),
		},
	)
	defer span.Finish()

	model := p.Model
	if model.Model.IsArchive {
		baseURL := model.Model.BaseUrl
		span.LogFields(
			olog.String("event", "download model archive"),
		)
		_, err := downloadmanager.DownloadInto(baseURL, p.WorkDir, downloadmanager.Context(ctx))
		if err != nil {
			return errors.Wrapf(err, "failed to download model archive from %v", model.Model.BaseUrl)
		}
		return nil
	}
	checksum := p.GetGraphChecksum()
	if checksum == "" {
		return errors.New("Need graph file checksum in the model manifest")
	}

	span.LogFields(
		olog.String("event", "download graph"),
	)

	if _, err := downloadmanager.DownloadFile(p.GetGraphUrl(), p.GetGraphPath(), downloadmanager.MD5Sum(checksum)); err != nil {
		return err
	}

	checksum = p.GetFeaturesChecksum()
	if checksum == "" {
		return errors.New("Need features file checksum in the model manifest")
	}

	span.LogFields(
		olog.String("event", "download features"),
	)
	if _, err := downloadmanager.DownloadFile(p.GetFeaturesUrl(), p.GetFeaturesPath(), downloadmanager.MD5Sum(checksum)); err != nil {
		return err
	}

	return nil
}

func (p ImagePredictor) GetInputLayerName(reader io.Reader) (string, error) {
	model := p.Model
	modelInputs := model.GetInputs()
	typeParameters := modelInputs[0].GetParameters()
	name, err := p.GetLayerName(typeParameters)
	if err != nil {
		graphDef, err := tensorflow.FromCheckpoint(reader)
		if err != nil {
			return "", errors.Wrap(err, "failed to read metagraph from checkpoint")
		}
		nodes := graphDef.GetNode()
		if nodes == nil {
			return "", errors.New("failed to read graph nodes")
		}
		// get the first node which has no input
		for _, n := range nodes {
			if len(n.GetInput()) == 0 {
				return n.GetName(), nil
			}
		}
		return "", errors.New("cannot determin the name of the input layer")
	}
	return name, nil
}

func (p ImagePredictor) GetOutputLayerName(reader io.Reader) (string, error) {
	model := p.Model
	modelOutput := model.GetOutput()
	typeParameters := modelOutput.GetParameters()
	name, err := p.GetLayerName(typeParameters)
	if err != nil {
		graphDef, err := tensorflow.FromCheckpoint(reader)
		if err != nil {
			return "", errors.Wrap(err, "failed to read metagraph from checkpoint")
		}
		nodes := graphDef.GetNode()
		if nodes == nil {
			return "", errors.New("failed to read graph nodes")
		}
		if len(nodes) == 0 {
			return "", errors.New("cannot determin the name of the output layer")
		}
		// get the last node
		return nodes[len(nodes)-1].GetName(), nil
	}
	return name, nil
}

type Graph struct {
	c *C.TF_Graph
}

func (p *ImagePredictor) loadPredictor(ctx context.Context) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoadPredictor")
	defer span.Finish()

	if p.tfSession != nil {
		return nil
	}

	span.LogFields(
		olog.String("event", "read features"),
	)

	var features []string
	f, err := os.Open(p.GetFeaturesPath())
	if err != nil {
		return errors.Wrapf(err, "cannot read %s", p.GetFeaturesPath())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		features = append(features, line)
	}
	p.features = features

	span.LogFields(
		olog.String("event", "read graph"),
	)
	model, err := ioutil.ReadFile(p.GetGraphPath())
	if err != nil {
		return errors.Wrapf(err, "cannot read %s", p.GetGraphPath())
	}

	// Construct an in-memory graph from the serialized form.
	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		return errors.Wrap(err, "unable to create tensorflow model graph")
	}

	modelReader := bytes.NewReader(model)
	p.inputLayer, err = p.GetInputLayerName(modelReader)
	if err != nil {
		return errors.Wrap(err, "failed to get input layer name")
	}
	p.outputLayer, err = p.GetOutputLayerName(modelReader)
	if err != nil {
		return errors.Wrap(err, "failed to get input layer name")
	}

	// Create a session for inference over graph.
	sessionOpts := &tf.SessionOptions{}
	if p.Options.UsesGPU() && false {
		if false {
		g := (*Graph)(unsafe.Pointer(graph))
		_ = g 
		}
		//C.setDefaultDevice("/device:GPU:0", g.c)
		sessionOpts = &tf.SessionOptions{Target: "GPU"}
		_ = nvidiasmi.GPUCount
		sessionConfig := tensorflow.ConfigProto{
			DeviceCount: map[string]int32{
				"GPU": int32(1), //int32(nvidiasmi.GPUCount),
			},
			LogDevicePlacement: true,
			GpuOptions: &tensorflow.GPUOptions{
				ForceGpuCompatible: true,
			},
		}
		if buf, err := sessionConfig.Marshal(); err == nil {
			sessionOpts.Config = buf
			pp.Println("buf = ", string(buf))
		}
	}
	session, err := tf.NewSession(graph, sessionOpts)
	if err != nil {
		return errors.Wrap(err, "unable to create tensorflow session")
	}

	p.tfGraph = graph
	p.tfSession = session

	return nil
}

func zeros(height, width, channels int) [][][]float32 {
	rows := make([][][]float32, height)
	for ii := range rows {
		columns := make([][]float32, width)
		for jj := range columns {
			columns[jj] = make([]float32, channels)
		}
		rows[ii] = columns
	}
	return rows
}

// Needs NHWC
func (p *ImagePredictor) makeTensorFromImageData(ctx context.Context, data0 [][]float32) (*tf.Tensor, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "makeTensorFromImageData")
	defer span.Finish()

	imageDims, err := p.GetImageDimensions()
	if err != nil {
		return nil, err
	}
	channels, height, width := int(imageDims[0]), int(imageDims[1]), int(imageDims[2])
	batchSize := int(p.BatchSize())
	if batchSize == 0 {
		batchSize = 1
	}

	makeImage := func(arry []float32) [][][]float32 {
		rows := make([][][]float32, height)
		for ii := range rows {
			columns := make([][]float32, width)
			for jj := range columns {
				offset := channels * (width*ii + jj)
				columns[jj] = arry[offset : offset+3]
			}
			rows[ii] = columns
		}
		return rows
	}

	// convert to a 4D tensor (batch, height, width, channels)
	data := make([][][][]float32, batchSize)
	for ii, e := range data0 {
		data[ii] = makeImage(e)
	}
	// perform padding
	if len(data0) < batchSize {
		z := zeros(height, width, channels)
		for ii := len(data0); ii < batchSize; ii++ {
			data[ii] = z
		}
	}

	tensor, err := tf.NewTensor(data)
	if err != nil {
		return nil, err
	}

	return tensor, nil
}

type operation struct {
	c *C.TF_Operation
	// A reference to the Graph to prevent it from
	// being GCed while the Operation is still alive.
	g *Graph
}

func (p *ImagePredictor) Predict(ctx context.Context, data [][]float32, opts ...options.Option) ([]dlframework.Features, error) {
	span := opentracing.SpanFromContext(ctx)
	_ = span

	session := p.tfSession
	graph := p.tfGraph

	options := options.New(opts...)

	batchSize := int(options.BatchSize())
	if batchSize == 0 {
		batchSize = 1
	}

	tensor, err := p.makeTensorFromImageData(ctx, data)
	if err != nil {
		return nil, errors.New("cannot make tensor from image data")
	}
 //
 //if options.UsesGPU() {
 //	op := (*operation)(unsafe.Pointer(graph.Operation(p.inputLayer)))
 //	cstr := C.CString("gpu:0")
 //	C.TF_SetDevice(op.c, cstr)
 //	C.free(unsafe.Pointer(cstr))
 //}
	fetches, err := session.Run(
		map[tf.Output]*tf.Tensor{
			graph.Operation(p.inputLayer).Output(0): tensor,
		},
		[]tf.Output{
			graph.Operation(p.outputLayer).Output(0),
		},
		nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to perform inference")
	}

	probabilities := fetches[0].Value().([][]float32)

	var output []dlframework.Features
	for _, prob := range probabilities {
		rprobs := make([]*dlframework.Feature, len(prob))
		for j, v := range prob {
			rprobs[j] = &dlframework.Feature{
				Index:       int64(j),
				Name:        "<> " + p.features[j],
				Probability: v,
			}
		}
		output = append(output, rprobs)
	}

	return output, nil
}

func (p *ImagePredictor) Reset(ctx context.Context) error {

	return nil
}

func (p *ImagePredictor) Close() error {
	if p.tfSession != nil {
		p.tfSession.Close()
	}
	return nil
}

func init() {
	config.AfterInit(func() {
		framework := tensorflow.FrameworkManifest
		agent.AddPredictor(framework, &ImagePredictor{
			ImagePredictor: common.ImagePredictor{
				Base: common.Base{
					Framework: framework,
				},
			},
		})
	})
}
