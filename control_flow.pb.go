// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: control_flow.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing the values in ControlFlowContext.
type ValuesDef struct {
	// Value names that have been seen in this context.
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// Value names referenced by but external to this context.
	ExternalValues map[string]string `protobuf:"bytes,2,rep,name=external_values,json=externalValues" json:"external_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ValuesDef) Reset()                    { *m = ValuesDef{} }
func (m *ValuesDef) String() string            { return proto.CompactTextString(m) }
func (*ValuesDef) ProtoMessage()               {}
func (*ValuesDef) Descriptor() ([]byte, []int) { return fileDescriptorControlFlow, []int{0} }

func (m *ValuesDef) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ValuesDef) GetExternalValues() map[string]string {
	if m != nil {
		return m.ExternalValues
	}
	return nil
}

// Protocol buffer representing a CondContext object.
type CondContextDef struct {
	// Name of the context.
	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
	// Name of the pred tensor.
	PredName string `protobuf:"bytes,2,opt,name=pred_name,json=predName,proto3" json:"pred_name,omitempty"`
	// Name of the pivot tensor.
	PivotName string `protobuf:"bytes,3,opt,name=pivot_name,json=pivotName,proto3" json:"pivot_name,omitempty"`
	// Branch prediction. 0 or 1.
	Branch int32 `protobuf:"varint,4,opt,name=branch,proto3" json:"branch,omitempty"`
	// Values and external values in control flow context.
	ValuesDef *ValuesDef `protobuf:"bytes,5,opt,name=values_def,json=valuesDef" json:"values_def,omitempty"`
}

func (m *CondContextDef) Reset()                    { *m = CondContextDef{} }
func (m *CondContextDef) String() string            { return proto.CompactTextString(m) }
func (*CondContextDef) ProtoMessage()               {}
func (*CondContextDef) Descriptor() ([]byte, []int) { return fileDescriptorControlFlow, []int{1} }

func (m *CondContextDef) GetContextName() string {
	if m != nil {
		return m.ContextName
	}
	return ""
}

func (m *CondContextDef) GetPredName() string {
	if m != nil {
		return m.PredName
	}
	return ""
}

func (m *CondContextDef) GetPivotName() string {
	if m != nil {
		return m.PivotName
	}
	return ""
}

func (m *CondContextDef) GetBranch() int32 {
	if m != nil {
		return m.Branch
	}
	return 0
}

func (m *CondContextDef) GetValuesDef() *ValuesDef {
	if m != nil {
		return m.ValuesDef
	}
	return nil
}

// Protocol buffer representing a WhileContext object.
type WhileContextDef struct {
	// Name of the context.
	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
	// The number of iterations allowed to run in parallel.
	ParallelIterations int32 `protobuf:"varint,2,opt,name=parallel_iterations,json=parallelIterations,proto3" json:"parallel_iterations,omitempty"`
	// Whether backprop is enabled for this while loop.
	BackProp bool `protobuf:"varint,3,opt,name=back_prop,json=backProp,proto3" json:"back_prop,omitempty"`
	// Whether GPU-CPU memory swap is enabled for this loop.
	SwapMemory bool `protobuf:"varint,4,opt,name=swap_memory,json=swapMemory,proto3" json:"swap_memory,omitempty"`
	// Name of the pivot tensor.
	PivotName string `protobuf:"bytes,5,opt,name=pivot_name,json=pivotName,proto3" json:"pivot_name,omitempty"`
	// Name of the pivot_for_pred tensor.
	PivotForPredName string `protobuf:"bytes,6,opt,name=pivot_for_pred_name,json=pivotForPredName,proto3" json:"pivot_for_pred_name,omitempty"`
	// Name of the pivot_for_body tensor.
	PivotForBodyName string `protobuf:"bytes,7,opt,name=pivot_for_body_name,json=pivotForBodyName,proto3" json:"pivot_for_body_name,omitempty"`
	// List of names for exit tensors.
	LoopExitNames []string `protobuf:"bytes,8,rep,name=loop_exit_names,json=loopExitNames" json:"loop_exit_names,omitempty"`
	// List of names for enter tensors.
	LoopEnterNames []string `protobuf:"bytes,10,rep,name=loop_enter_names,json=loopEnterNames" json:"loop_enter_names,omitempty"`
	// Values and external values in control flow context.
	ValuesDef *ValuesDef `protobuf:"bytes,9,opt,name=values_def,json=valuesDef" json:"values_def,omitempty"`
}

func (m *WhileContextDef) Reset()                    { *m = WhileContextDef{} }
func (m *WhileContextDef) String() string            { return proto.CompactTextString(m) }
func (*WhileContextDef) ProtoMessage()               {}
func (*WhileContextDef) Descriptor() ([]byte, []int) { return fileDescriptorControlFlow, []int{2} }

func (m *WhileContextDef) GetContextName() string {
	if m != nil {
		return m.ContextName
	}
	return ""
}

func (m *WhileContextDef) GetParallelIterations() int32 {
	if m != nil {
		return m.ParallelIterations
	}
	return 0
}

func (m *WhileContextDef) GetBackProp() bool {
	if m != nil {
		return m.BackProp
	}
	return false
}

func (m *WhileContextDef) GetSwapMemory() bool {
	if m != nil {
		return m.SwapMemory
	}
	return false
}

func (m *WhileContextDef) GetPivotName() string {
	if m != nil {
		return m.PivotName
	}
	return ""
}

func (m *WhileContextDef) GetPivotForPredName() string {
	if m != nil {
		return m.PivotForPredName
	}
	return ""
}

func (m *WhileContextDef) GetPivotForBodyName() string {
	if m != nil {
		return m.PivotForBodyName
	}
	return ""
}

func (m *WhileContextDef) GetLoopExitNames() []string {
	if m != nil {
		return m.LoopExitNames
	}
	return nil
}

func (m *WhileContextDef) GetLoopEnterNames() []string {
	if m != nil {
		return m.LoopEnterNames
	}
	return nil
}

func (m *WhileContextDef) GetValuesDef() *ValuesDef {
	if m != nil {
		return m.ValuesDef
	}
	return nil
}

func init() {
	proto.RegisterType((*ValuesDef)(nil), "tensorflow.ValuesDef")
	proto.RegisterType((*CondContextDef)(nil), "tensorflow.CondContextDef")
	proto.RegisterType((*WhileContextDef)(nil), "tensorflow.WhileContextDef")
}
func (m *ValuesDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValuesDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.ExternalValues) > 0 {
		for k, _ := range m.ExternalValues {
			dAtA[i] = 0x12
			i++
			v := m.ExternalValues[k]
			mapSize := 1 + len(k) + sovControlFlow(uint64(len(k))) + 1 + len(v) + sovControlFlow(uint64(len(v)))
			i = encodeVarintControlFlow(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintControlFlow(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintControlFlow(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *CondContextDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CondContextDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContextName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(len(m.ContextName)))
		i += copy(dAtA[i:], m.ContextName)
	}
	if len(m.PredName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(len(m.PredName)))
		i += copy(dAtA[i:], m.PredName)
	}
	if len(m.PivotName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(len(m.PivotName)))
		i += copy(dAtA[i:], m.PivotName)
	}
	if m.Branch != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(m.Branch))
	}
	if m.ValuesDef != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(m.ValuesDef.Size()))
		n1, err := m.ValuesDef.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *WhileContextDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhileContextDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContextName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(len(m.ContextName)))
		i += copy(dAtA[i:], m.ContextName)
	}
	if m.ParallelIterations != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(m.ParallelIterations))
	}
	if m.BackProp {
		dAtA[i] = 0x18
		i++
		if m.BackProp {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.SwapMemory {
		dAtA[i] = 0x20
		i++
		if m.SwapMemory {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.PivotName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(len(m.PivotName)))
		i += copy(dAtA[i:], m.PivotName)
	}
	if len(m.PivotForPredName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(len(m.PivotForPredName)))
		i += copy(dAtA[i:], m.PivotForPredName)
	}
	if len(m.PivotForBodyName) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(len(m.PivotForBodyName)))
		i += copy(dAtA[i:], m.PivotForBodyName)
	}
	if len(m.LoopExitNames) > 0 {
		for _, s := range m.LoopExitNames {
			dAtA[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.ValuesDef != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintControlFlow(dAtA, i, uint64(m.ValuesDef.Size()))
		n2, err := m.ValuesDef.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.LoopEnterNames) > 0 {
		for _, s := range m.LoopEnterNames {
			dAtA[i] = 0x52
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64ControlFlow(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ControlFlow(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintControlFlow(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ValuesDef) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovControlFlow(uint64(l))
		}
	}
	if len(m.ExternalValues) > 0 {
		for k, v := range m.ExternalValues {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovControlFlow(uint64(len(k))) + 1 + len(v) + sovControlFlow(uint64(len(v)))
			n += mapEntrySize + 1 + sovControlFlow(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *CondContextDef) Size() (n int) {
	var l int
	_ = l
	l = len(m.ContextName)
	if l > 0 {
		n += 1 + l + sovControlFlow(uint64(l))
	}
	l = len(m.PredName)
	if l > 0 {
		n += 1 + l + sovControlFlow(uint64(l))
	}
	l = len(m.PivotName)
	if l > 0 {
		n += 1 + l + sovControlFlow(uint64(l))
	}
	if m.Branch != 0 {
		n += 1 + sovControlFlow(uint64(m.Branch))
	}
	if m.ValuesDef != nil {
		l = m.ValuesDef.Size()
		n += 1 + l + sovControlFlow(uint64(l))
	}
	return n
}

func (m *WhileContextDef) Size() (n int) {
	var l int
	_ = l
	l = len(m.ContextName)
	if l > 0 {
		n += 1 + l + sovControlFlow(uint64(l))
	}
	if m.ParallelIterations != 0 {
		n += 1 + sovControlFlow(uint64(m.ParallelIterations))
	}
	if m.BackProp {
		n += 2
	}
	if m.SwapMemory {
		n += 2
	}
	l = len(m.PivotName)
	if l > 0 {
		n += 1 + l + sovControlFlow(uint64(l))
	}
	l = len(m.PivotForPredName)
	if l > 0 {
		n += 1 + l + sovControlFlow(uint64(l))
	}
	l = len(m.PivotForBodyName)
	if l > 0 {
		n += 1 + l + sovControlFlow(uint64(l))
	}
	if len(m.LoopExitNames) > 0 {
		for _, s := range m.LoopExitNames {
			l = len(s)
			n += 1 + l + sovControlFlow(uint64(l))
		}
	}
	if m.ValuesDef != nil {
		l = m.ValuesDef.Size()
		n += 1 + l + sovControlFlow(uint64(l))
	}
	if len(m.LoopEnterNames) > 0 {
		for _, s := range m.LoopEnterNames {
			l = len(s)
			n += 1 + l + sovControlFlow(uint64(l))
		}
	}
	return n
}

func sovControlFlow(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozControlFlow(x uint64) (n int) {
	return sovControlFlow(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValuesDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControlFlow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValuesDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValuesDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalValues", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthControlFlow
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.ExternalValues == nil {
				m.ExternalValues = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowControlFlow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowControlFlow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthControlFlow
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.ExternalValues[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.ExternalValues[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControlFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControlFlow
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CondContextDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControlFlow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CondContextDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CondContextDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContextName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContextName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PredName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PredName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PivotName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PivotName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Branch", wireType)
			}
			m.Branch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Branch |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValuesDef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValuesDef == nil {
				m.ValuesDef = &ValuesDef{}
			}
			if err := m.ValuesDef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControlFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControlFlow
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WhileContextDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControlFlow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WhileContextDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhileContextDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContextName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContextName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParallelIterations", wireType)
			}
			m.ParallelIterations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParallelIterations |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackProp", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BackProp = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapMemory", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SwapMemory = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PivotName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PivotName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PivotForPredName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PivotForPredName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PivotForBodyName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PivotForBodyName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoopExitNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LoopExitNames = append(m.LoopExitNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValuesDef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValuesDef == nil {
				m.ValuesDef = &ValuesDef{}
			}
			if err := m.ValuesDef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoopEnterNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControlFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LoopEnterNames = append(m.LoopEnterNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControlFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControlFlow
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipControlFlow(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowControlFlow
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowControlFlow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthControlFlow
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowControlFlow
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipControlFlow(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthControlFlow = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowControlFlow   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("control_flow.proto", fileDescriptorControlFlow) }

var fileDescriptorControlFlow = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0xae, 0x12, 0x3f,
	0x14, 0xc7, 0x7f, 0x85, 0x1f, 0xc8, 0x1c, 0x14, 0xb0, 0xe8, 0xcd, 0x44, 0x23, 0x22, 0x0b, 0x33,
	0x2e, 0xc4, 0xe4, 0xea, 0xc2, 0xb8, 0x13, 0xe4, 0x26, 0x2e, 0x34, 0x64, 0x16, 0xba, 0x9c, 0x14,
	0x28, 0xde, 0x09, 0xa5, 0x67, 0xd2, 0xa9, 0xfc, 0x79, 0x0b, 0x1f, 0xc3, 0x67, 0x30, 0x71, 0xef,
	0xd2, 0x47, 0x30, 0xf8, 0x12, 0x2e, 0x4d, 0x4f, 0x87, 0x7b, 0x85, 0xdc, 0x85, 0xee, 0x7a, 0x3e,
	0xdf, 0xef, 0x7c, 0xdb, 0x73, 0x3a, 0x05, 0x3e, 0x45, 0x6d, 0x0d, 0xaa, 0x64, 0xae, 0x70, 0xdd,
	0xcf, 0x0c, 0x5a, 0xe4, 0x60, 0xa5, 0xce, 0xd1, 0x38, 0xd2, 0xfb, 0xc2, 0x20, 0x78, 0x27, 0xd4,
	0x47, 0x99, 0xbf, 0x92, 0x73, 0x7e, 0x02, 0xd5, 0x15, 0x15, 0x21, 0xeb, 0x96, 0xa3, 0x20, 0x2e,
	0x2a, 0x1e, 0x43, 0x53, 0x6e, 0xac, 0x34, 0x5a, 0xa8, 0xa4, 0x30, 0x94, 0xba, 0xe5, 0xa8, 0x7e,
	0xfa, 0xa8, 0x7f, 0x99, 0xd5, 0xbf, 0xc8, 0xe9, 0x8f, 0x0a, 0xb3, 0x27, 0x23, 0x6d, 0xcd, 0x36,
	0x6e, 0xc8, 0x03, 0x78, 0xe7, 0x25, 0xb4, 0xaf, 0xb0, 0xf1, 0x16, 0x94, 0x17, 0x72, 0x1b, 0xb2,
	0x2e, 0x8b, 0x82, 0xd8, 0x2d, 0xf9, 0x2d, 0xa8, 0xd0, 0x9e, 0x61, 0x89, 0x98, 0x2f, 0x5e, 0x94,
	0x9e, 0xb3, 0xde, 0x57, 0x06, 0x8d, 0x21, 0xea, 0xd9, 0x10, 0xb5, 0x95, 0x1b, 0xeb, 0x3a, 0x78,
	0x00, 0xd7, 0xa7, 0xbe, 0x4a, 0xb4, 0x58, 0xca, 0x22, 0xa7, 0x5e, 0xb0, 0xb7, 0x62, 0x29, 0xf9,
	0x5d, 0x08, 0x32, 0x23, 0x67, 0x5e, 0xf7, 0x99, 0x35, 0x07, 0x48, 0xbc, 0x07, 0x90, 0xa5, 0x2b,
	0x2c, 0xbe, 0x2e, 0x93, 0x1a, 0x10, 0x21, 0xf9, 0x04, 0xaa, 0x13, 0x23, 0xf4, 0xf4, 0x3c, 0xfc,
	0xbf, 0xcb, 0xa2, 0x4a, 0x5c, 0x54, 0xfc, 0x19, 0x80, 0x9f, 0x4b, 0x32, 0x93, 0xf3, 0xb0, 0xd2,
	0x65, 0x51, 0xfd, 0xf4, 0xf6, 0x95, 0xb3, 0x89, 0x83, 0xd5, 0x7e, 0xd9, 0xfb, 0x5c, 0x86, 0xe6,
	0xfb, 0xf3, 0x54, 0xc9, 0x7f, 0x6b, 0xe0, 0x09, 0xb4, 0x33, 0x61, 0x84, 0x52, 0x52, 0x25, 0xa9,
	0x95, 0x46, 0xd8, 0x14, 0x75, 0x4e, 0xad, 0x54, 0x62, 0xbe, 0x97, 0x5e, 0x5f, 0x28, 0xae, 0xe3,
	0x89, 0x98, 0x2e, 0x92, 0xcc, 0x60, 0x46, 0x3d, 0xd5, 0xe2, 0x9a, 0x03, 0x63, 0x83, 0x19, 0xbf,
	0x0f, 0xf5, 0x7c, 0x2d, 0xb2, 0x64, 0x29, 0x97, 0x68, 0xb6, 0xd4, 0x57, 0x2d, 0x06, 0x87, 0xde,
	0x10, 0x39, 0x1a, 0x49, 0xe5, 0x78, 0x24, 0x8f, 0xa1, 0xed, 0xe5, 0x39, 0x9a, 0xe4, 0x72, 0xb0,
	0x55, 0xf2, 0xb5, 0x48, 0x3a, 0x43, 0x33, 0xde, 0x0f, 0xf8, 0xc0, 0x3e, 0xc1, 0xd9, 0xd6, 0xdb,
	0xaf, 0x1d, 0xda, 0x07, 0x38, 0xdb, 0x92, 0xfd, 0x21, 0x34, 0x15, 0x62, 0x96, 0xc8, 0x4d, 0xea,
	0x0f, 0x90, 0x87, 0x35, 0xfa, 0x35, 0x6f, 0x38, 0x3c, 0xda, 0xa4, 0x74, 0x88, 0xfc, 0xe8, 0x02,
	0x82, 0xbf, 0xbb, 0x00, 0x1e, 0x41, 0xcb, 0xa7, 0x6b, 0x2b, 0x4d, 0x11, 0x0f, 0x14, 0xdf, 0xa0,
	0x78, 0x87, 0x29, 0x7f, 0x30, 0xf8, 0xb6, 0xeb, 0xb0, 0xef, 0xbb, 0x0e, 0xfb, 0xb1, 0xeb, 0xb0,
	0x4f, 0x3f, 0x3b, 0xff, 0x41, 0x88, 0xe6, 0xc3, 0x9f, 0x1b, 0xcc, 0x8d, 0x58, 0xca, 0x35, 0x9a,
	0xc5, 0xe0, 0xe6, 0xd0, 0xbf, 0xb9, 0x33, 0x85, 0xeb, 0xb1, 0x7b, 0x71, 0xf9, 0x98, 0xfd, 0x62,
	0x6c, 0x52, 0xa5, 0xe7, 0xf7, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xf1, 0x6f, 0xd5,
	0x94, 0x03, 0x00, 0x00,
}