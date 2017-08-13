package main

import (
	"fmt"
	"os"

	"github.com/rai-project/dlframework/framework/cmd"
	"github.com/rai-project/tensorflow"
	_ "github.com/rai-project/tensorflow/predict"
)

func main() {
	cmd.Init()

	rootCmd, err := cmd.NewRootCommand(tensorflow.FrameworkManifest)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
