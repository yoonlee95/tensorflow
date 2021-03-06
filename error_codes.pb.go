// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: error_codes.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The canonical error codes for TensorFlow APIs.
//
// Warnings:
//
// -   Do not change any numeric assignments.
// -   Changes to this list should only be made if there is a compelling
//     need that can't be satisfied in another way.  Such changes
//     must be approved by at least two OWNERS.
//
// Sometimes multiple error codes may apply.  Services should return
// the most specific error code that applies.  For example, prefer
// OUT_OF_RANGE over FAILED_PRECONDITION if both codes apply.
// Similarly prefer NOT_FOUND or ALREADY_EXISTS over FAILED_PRECONDITION.
type Error_Code int32

const (
	// Not an error; returned on success
	Error_OK Error_Code = 0
	// The operation was cancelled (typically by the caller).
	Error_CANCELLED Error_Code = 1
	// Unknown error.  An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Error_UNKNOWN Error_Code = 2
	// Client specified an invalid argument.  Note that this differs
	// from FAILED_PRECONDITION.  INVALID_ARGUMENT indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	Error_INVALID_ARGUMENT Error_Code = 3
	// Deadline expired before operation could complete.  For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example, a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	Error_DEADLINE_EXCEEDED Error_Code = 4
	// Some requested entity (e.g., file or directory) was not found.
	// For privacy reasons, this code *may* be returned when the client
	// does not have the access right to the entity.
	Error_NOT_FOUND Error_Code = 5
	// Some entity that we attempted to create (e.g., file or directory)
	// already exists.
	Error_ALREADY_EXISTS Error_Code = 6
	// The caller does not have permission to execute the specified
	// operation.  PERMISSION_DENIED must not be used for rejections
	// caused by exhausting some resource (use RESOURCE_EXHAUSTED
	// instead for those errors).  PERMISSION_DENIED must not be
	// used if the caller can not be identified (use UNAUTHENTICATED
	// instead for those errors).
	Error_PERMISSION_DENIED Error_Code = 7
	// The request does not have valid authentication credentials for the
	// operation.
	Error_UNAUTHENTICATED Error_Code = 16
	// Some resource has been exhausted, perhaps a per-user quota, or
	// perhaps the entire file system is out of space.
	Error_RESOURCE_EXHAUSTED Error_Code = 8
	// Operation was rejected because the system is not in a state
	// required for the operation's execution.  For example, directory
	// to be deleted may be non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FAILED_PRECONDITION, ABORTED, and UNAVAILABLE:
	//  (a) Use UNAVAILABLE if the client can retry just the failing call.
	//  (b) Use ABORTED if the client should retry at a higher-level
	//      (e.g., restarting a read-modify-write sequence).
	//  (c) Use FAILED_PRECONDITION if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, FAILED_PRECONDITION
	//      should be returned since the client should not retry unless
	//      they have first fixed up the directory by deleting files from it.
	//  (d) Use FAILED_PRECONDITION if the client performs conditional
	//      REST Get/Update/Delete on a resource and the resource on the
	//      server does not match the condition. E.g., conflicting
	//      read-modify-write on the same resource.
	Error_FAILED_PRECONDITION Error_Code = 9
	// The operation was aborted, typically due to a concurrency issue
	// like sequencer check failures, transaction aborts, etc.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Error_ABORTED Error_Code = 10
	// Operation tried to iterate past the valid input range.  E.g., seeking or
	// reading past end of file.
	//
	// Unlike INVALID_ARGUMENT, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate INVALID_ARGUMENT if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// OUT_OF_RANGE if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between FAILED_PRECONDITION and
	// OUT_OF_RANGE.  We recommend using OUT_OF_RANGE (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an OUT_OF_RANGE error to detect when
	// they are done.
	Error_OUT_OF_RANGE Error_Code = 11
	// Operation is not implemented or not supported/enabled in this service.
	Error_UNIMPLEMENTED Error_Code = 12
	// Internal errors.  Means some invariants expected by underlying
	// system has been broken.  If you see one of these errors,
	// something is very broken.
	Error_INTERNAL Error_Code = 13
	// The service is currently unavailable.  This is a most likely a
	// transient condition and may be corrected by retrying with
	// a backoff.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Error_UNAVAILABLE Error_Code = 14
	// Unrecoverable data loss or corruption.
	Error_DATA_LOSS Error_Code = 15
	// An extra enum entry to prevent people from writing code that
	// fails to compile when a new code is added.
	//
	// Nobody should ever reference this enumeration entry. In particular,
	// if you write C++ code that switches on this enumeration, add a default:
	// case instead of a case that mentions this enumeration entry.
	//
	// Nobody should rely on the value (currently 20) listed here.  It
	// may change in the future.
	Error_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_ Error_Code = 20
)

var Error_Code_name = map[int32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	16: "UNAUTHENTICATED",
	8:  "RESOURCE_EXHAUSTED",
	9:  "FAILED_PRECONDITION",
	10: "ABORTED",
	11: "OUT_OF_RANGE",
	12: "UNIMPLEMENTED",
	13: "INTERNAL",
	14: "UNAVAILABLE",
	15: "DATA_LOSS",
	20: "DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_",
}
var Error_Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     16,
	"RESOURCE_EXHAUSTED":  8,
	"FAILED_PRECONDITION": 9,
	"ABORTED":             10,
	"OUT_OF_RANGE":        11,
	"UNIMPLEMENTED":       12,
	"INTERNAL":            13,
	"UNAVAILABLE":         14,
	"DATA_LOSS":           15,
	"DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_": 20,
}

func (x Error_Code) String() string {
	return proto.EnumName(Error_Code_name, int32(x))
}
func (Error_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptorErrorCodes, []int{0, 0} }

type Error struct {
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorErrorCodes, []int{0} }

func init() {
	proto.RegisterType((*Error)(nil), "tensorflow.error")
	proto.RegisterEnum("tensorflow.Error_Code", Error_Code_name, Error_Code_value)
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64ErrorCodes(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32ErrorCodes(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintErrorCodes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Error) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovErrorCodes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozErrorCodes(x uint64) (n int) {
	return sovErrorCodes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrorCodes
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
			return fmt.Errorf("proto: error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipErrorCodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrorCodes
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
func skipErrorCodes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrorCodes
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
					return 0, ErrIntOverflowErrorCodes
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
					return 0, ErrIntOverflowErrorCodes
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
				return 0, ErrInvalidLengthErrorCodes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowErrorCodes
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
				next, err := skipErrorCodes(dAtA[start:])
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
	ErrInvalidLengthErrorCodes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrorCodes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("error_codes.proto", fileDescriptorErrorCodes) }

var fileDescriptorErrorCodes = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0x3b, 0x6e, 0x1b, 0x31,
	0x10, 0x86, 0xbd, 0x96, 0x2d, 0xdb, 0x94, 0x64, 0x8d, 0x68, 0xe7, 0x51, 0xa9, 0xf0, 0x01, 0xd4,
	0xe4, 0x04, 0xa3, 0xe5, 0xac, 0x44, 0x88, 0x1a, 0x2e, 0xf8, 0x90, 0x9d, 0x8a, 0x48, 0x62, 0x39,
	0x45, 0x1e, 0x0c, 0x56, 0x06, 0x7c, 0x81, 0xd4, 0x41, 0x8e, 0x95, 0x32, 0x47, 0x08, 0x94, 0x4b,
	0xa4, 0x0c, 0xa8, 0xc6, 0xed, 0xe0, 0x9f, 0x99, 0xff, 0xc3, 0x27, 0x26, 0xdb, 0xae, 0xcb, 0x5d,
	0xfa, 0x90, 0xef, 0xb7, 0xbb, 0xd9, 0xb7, 0x2e, 0x3f, 0x66, 0x29, 0x1e, 0xb7, 0x5f, 0x77, 0xb9,
	0x7b, 0xf8, 0x9c, 0x9f, 0x6e, 0x7e, 0xf4, 0xc4, 0xe9, 0x21, 0x71, 0xf3, 0xbd, 0x27, 0x4e, 0xea,
	0x7c, 0xbf, 0x95, 0x7d, 0x71, 0x6c, 0x57, 0x70, 0x24, 0x47, 0xe2, 0xa2, 0x46, 0xae, 0xc9, 0x18,
	0x52, 0x50, 0xc9, 0x81, 0x38, 0x8b, 0xbc, 0x62, 0x7b, 0xcb, 0x70, 0x2c, 0xaf, 0x05, 0x68, 0xde,
	0xa0, 0xd1, 0x2a, 0xa1, 0x5b, 0xc4, 0x35, 0x71, 0x80, 0x9e, 0x7c, 0x21, 0x26, 0x8a, 0x50, 0x19,
	0xcd, 0x94, 0xe8, 0xae, 0x26, 0x52, 0xa4, 0xe0, 0xa4, 0x1c, 0x62, 0x1b, 0x52, 0x63, 0x23, 0x2b,
	0x38, 0x95, 0x52, 0x5c, 0xa2, 0x71, 0x84, 0xea, 0x6d, 0xa2, 0x3b, 0xed, 0x83, 0x87, 0x7e, 0xd9,
	0x6c, 0xc9, 0xad, 0xb5, 0xf7, 0xda, 0x72, 0x52, 0xc4, 0x9a, 0x14, 0x9c, 0xc9, 0x2b, 0x31, 0x8e,
	0x8c, 0x31, 0x2c, 0x89, 0x83, 0xae, 0x31, 0x90, 0x02, 0x90, 0x2f, 0x85, 0x74, 0xe4, 0x6d, 0x74,
	0x75, 0xf9, 0xb2, 0xc4, 0xe8, 0xcb, 0xfc, 0x5c, 0xbe, 0x12, 0x57, 0x0d, 0x6a, 0x43, 0x2a, 0xb5,
	0x8e, 0x6a, 0xcb, 0x4a, 0x07, 0x6d, 0x19, 0x2e, 0x4a, 0x73, 0x9c, 0x5b, 0x57, 0x52, 0x42, 0x82,
	0x18, 0xda, 0x18, 0x92, 0x6d, 0x92, 0x43, 0x5e, 0x10, 0x0c, 0xe4, 0x44, 0x8c, 0x22, 0xeb, 0x75,
	0x6b, 0xa8, 0x60, 0x90, 0x82, 0xa1, 0x1c, 0x8a, 0x73, 0xcd, 0x81, 0x1c, 0xa3, 0x81, 0x91, 0x1c,
	0x8b, 0x41, 0x64, 0xdc, 0xa0, 0x36, 0x38, 0x37, 0x04, 0x97, 0x05, 0x48, 0x61, 0xc0, 0x64, 0xac,
	0xf7, 0x30, 0x96, 0x2b, 0xb1, 0x50, 0x36, 0x15, 0xc4, 0xe8, 0x29, 0x39, 0xf2, 0xe4, 0x36, 0xa4,
	0x52, 0x63, 0x5d, 0x6a, 0x62, 0x88, 0xae, 0xd4, 0x6c, 0x91, 0x0f, 0x68, 0x25, 0xa1, 0xa8, 0xc1,
	0x68, 0x42, 0xd2, 0x9c, 0xfc, 0xad, 0x0e, 0xf5, 0x32, 0x69, 0xf6, 0x81, 0x50, 0x25, 0xb8, 0x9e,
	0xe3, 0xaf, 0xfd, 0xb4, 0xfa, 0xbd, 0x9f, 0x56, 0x7f, 0xf6, 0xd3, 0xea, 0xe7, 0xdf, 0xe9, 0x91,
	0x78, 0x9d, 0xbb, 0x8f, 0xb3, 0x67, 0x65, 0xb3, 0x87, 0xee, 0xdd, 0x97, 0xed, 0x53, 0xee, 0x3e,
	0xcd, 0x81, 0x8a, 0xb9, 0x22, 0x6d, 0xd7, 0x16, 0xb3, 0xbb, 0xb6, 0xfa, 0x57, 0x55, 0xef, 0xfb,
	0x07, 0xcd, 0x6f, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x31, 0x88, 0x0b, 0xce, 0xfb, 0x01, 0x00,
	0x00,
}
