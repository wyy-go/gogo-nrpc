// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nrpc.proto

package nrpc

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SubjectRule int32

const (
	SubjectRule_COPY    SubjectRule = 0
	SubjectRule_TOLOWER SubjectRule = 1
)

var SubjectRule_name = map[int32]string{
	0: "COPY",
	1: "TOLOWER",
}

var SubjectRule_value = map[string]int32{
	"COPY":    0,
	"TOLOWER": 1,
}

func (x SubjectRule) String() string {
	return proto.EnumName(SubjectRule_name, int32(x))
}

func (SubjectRule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f9057aa2cf08065b, []int{0}
}

type Error_Type int32

const (
	Error_CLIENT        Error_Type = 0
	Error_SERVER        Error_Type = 1
	Error_EOS           Error_Type = 3
	Error_SERVERTOOBUSY Error_Type = 4
)

var Error_Type_name = map[int32]string{
	0: "CLIENT",
	1: "SERVER",
	3: "EOS",
	4: "SERVERTOOBUSY",
}

var Error_Type_value = map[string]int32{
	"CLIENT":        0,
	"SERVER":        1,
	"EOS":           3,
	"SERVERTOOBUSY": 4,
}

func (x Error_Type) String() string {
	return proto.EnumName(Error_Type_name, int32(x))
}

func (Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f9057aa2cf08065b, []int{0, 0}
}

type Error struct {
	Type     Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=nrpc.Error_Type" json:"type,omitempty"`
	Message  string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MsgCount uint32     `protobuf:"varint,3,opt,name=msgCount,proto3" json:"msgCount,omitempty"`
}

func (m *Error) Reset()      { *m = Error{} }
func (*Error) ProtoMessage() {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9057aa2cf08065b, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetType() Error_Type {
	if m != nil {
		return m.Type
	}
	return Error_CLIENT
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetMsgCount() uint32 {
	if m != nil {
		return m.MsgCount
	}
	return 0
}

type Void struct {
}

func (m *Void) Reset()      { *m = Void{} }
func (*Void) ProtoMessage() {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9057aa2cf08065b, []int{1}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Void.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return m.Size()
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type NoRequest struct {
}

func (m *NoRequest) Reset()      { *m = NoRequest{} }
func (*NoRequest) ProtoMessage() {}
func (*NoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9057aa2cf08065b, []int{2}
}
func (m *NoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoRequest.Merge(m, src)
}
func (m *NoRequest) XXX_Size() int {
	return m.Size()
}
func (m *NoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NoRequest proto.InternalMessageInfo

type NoReply struct {
}

func (m *NoReply) Reset()      { *m = NoReply{} }
func (*NoReply) ProtoMessage() {}
func (*NoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9057aa2cf08065b, []int{3}
}
func (m *NoReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoReply.Merge(m, src)
}
func (m *NoReply) XXX_Size() int {
	return m.Size()
}
func (m *NoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NoReply.DiscardUnknown(m)
}

var xxx_messageInfo_NoReply proto.InternalMessageInfo

type HeartBeat struct {
	Lastbeat bool `protobuf:"varint,1,opt,name=lastbeat,proto3" json:"lastbeat,omitempty"`
}

func (m *HeartBeat) Reset()      { *m = HeartBeat{} }
func (*HeartBeat) ProtoMessage() {}
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9057aa2cf08065b, []int{4}
}
func (m *HeartBeat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeartBeat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeat.Merge(m, src)
}
func (m *HeartBeat) XXX_Size() int {
	return m.Size()
}
func (m *HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeat proto.InternalMessageInfo

func (m *HeartBeat) GetLastbeat() bool {
	if m != nil {
		return m.Lastbeat
	}
	return false
}

var E_PackageSubject = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "nrpc.packageSubject",
	Tag:           "bytes,50000,opt,name=packageSubject",
	Filename:      "nrpc.proto",
}

var E_PackageSubjectParams = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         50001,
	Name:          "nrpc.packageSubjectParams",
	Tag:           "bytes,50001,rep,name=packageSubjectParams",
	Filename:      "nrpc.proto",
}

var E_ServiceSubjectRule = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*SubjectRule)(nil),
	Field:         50002,
	Name:          "nrpc.serviceSubjectRule",
	Tag:           "varint,50002,opt,name=serviceSubjectRule,enum=nrpc.SubjectRule",
	Filename:      "nrpc.proto",
}

var E_MethodSubjectRule = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*SubjectRule)(nil),
	Field:         50003,
	Name:          "nrpc.methodSubjectRule",
	Tag:           "varint,50003,opt,name=methodSubjectRule,enum=nrpc.SubjectRule",
	Filename:      "nrpc.proto",
}

var E_ServiceSubject = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51000,
	Name:          "nrpc.serviceSubject",
	Tag:           "bytes,51000,opt,name=serviceSubject",
	Filename:      "nrpc.proto",
}

var E_ServiceSubjectParams = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         51001,
	Name:          "nrpc.serviceSubjectParams",
	Tag:           "bytes,51001,rep,name=serviceSubjectParams",
	Filename:      "nrpc.proto",
}

var E_MethodSubject = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         52000,
	Name:          "nrpc.methodSubject",
	Tag:           "bytes,52000,opt,name=methodSubject",
	Filename:      "nrpc.proto",
}

var E_MethodSubjectParams = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         52001,
	Name:          "nrpc.methodSubjectParams",
	Tag:           "bytes,52001,rep,name=methodSubjectParams",
	Filename:      "nrpc.proto",
}

var E_StreamedReply = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         52002,
	Name:          "nrpc.streamedReply",
	Tag:           "varint,52002,opt,name=streamedReply",
	Filename:      "nrpc.proto",
}

func init() {
	proto.RegisterEnum("nrpc.SubjectRule", SubjectRule_name, SubjectRule_value)
	proto.RegisterEnum("nrpc.Error_Type", Error_Type_name, Error_Type_value)
	proto.RegisterType((*Error)(nil), "nrpc.Error")
	proto.RegisterType((*Void)(nil), "nrpc.Void")
	proto.RegisterType((*NoRequest)(nil), "nrpc.NoRequest")
	proto.RegisterType((*NoReply)(nil), "nrpc.NoReply")
	proto.RegisterType((*HeartBeat)(nil), "nrpc.HeartBeat")
	proto.RegisterExtension(E_PackageSubject)
	proto.RegisterExtension(E_PackageSubjectParams)
	proto.RegisterExtension(E_ServiceSubjectRule)
	proto.RegisterExtension(E_MethodSubjectRule)
	proto.RegisterExtension(E_ServiceSubject)
	proto.RegisterExtension(E_ServiceSubjectParams)
	proto.RegisterExtension(E_MethodSubject)
	proto.RegisterExtension(E_MethodSubjectParams)
	proto.RegisterExtension(E_StreamedReply)
}

func init() { proto.RegisterFile("nrpc.proto", fileDescriptor_f9057aa2cf08065b) }

var fileDescriptor_f9057aa2cf08065b = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0xd8, 0xe4, 0xc7, 0xab, 0x1a, 0x39, 0x47, 0x87, 0x28, 0x42, 0x47, 0x64, 0x55,
	0x22, 0x62, 0x70, 0xa5, 0xb2, 0x99, 0x2d, 0x21, 0x11, 0x95, 0x4a, 0x5d, 0x5d, 0xd2, 0xa2, 0xb2,
	0x20, 0x27, 0x39, 0x4c, 0x20, 0xc9, 0x19, 0xfb, 0x82, 0x94, 0x8d, 0xa9, 0x63, 0xc5, 0xd4, 0x95,
	0x1f, 0x13, 0x7f, 0x42, 0xd9, 0x18, 0x3b, 0x06, 0x58, 0x3a, 0x36, 0xce, 0x3f, 0xc0, 0xc8, 0x88,
	0x7c, 0x4e, 0xaa, 0x9a, 0x04, 0xc2, 0x76, 0xef, 0xdd, 0xfb, 0x7e, 0xde, 0xd7, 0xdf, 0x33, 0xc0,
	0xc0, 0xf7, 0xda, 0xa6, 0xe7, 0x73, 0xc1, 0xb1, 0x16, 0x9d, 0x8b, 0x25, 0x97, 0x73, 0xb7, 0xc7,
	0xb6, 0x64, 0xaf, 0x35, 0x7c, 0xbe, 0xd5, 0x61, 0x41, 0xdb, 0xef, 0x7a, 0x82, 0xfb, 0xf1, 0x5c,
	0x11, 0x5c, 0xee, 0xf2, 0xf8, 0x6c, 0xbc, 0x47, 0x70, 0xb3, 0xe6, 0xfb, 0xdc, 0xc7, 0x9b, 0xa0,
	0x89, 0x91, 0xc7, 0x0a, 0xa8, 0x84, 0xca, 0xb9, 0x6d, 0xdd, 0x94, 0x60, 0x79, 0x65, 0x36, 0x47,
	0x1e, 0xa3, 0xf2, 0x16, 0x17, 0x20, 0xdd, 0x67, 0x41, 0xe0, 0xb8, 0xac, 0x70, 0xa3, 0x84, 0xca,
	0x59, 0x3a, 0x2f, 0x71, 0x11, 0x32, 0xfd, 0xc0, 0xad, 0xf2, 0xe1, 0x40, 0x14, 0xd4, 0x12, 0x2a,
	0xaf, 0xd3, 0xab, 0xda, 0xb0, 0x40, 0x8b, 0x18, 0x18, 0x20, 0x55, 0xdd, 0xdd, 0xa9, 0xed, 0x35,
	0x75, 0x25, 0x3a, 0x37, 0x6a, 0xf4, 0xb0, 0x46, 0x75, 0x84, 0xd3, 0xa0, 0xd6, 0xec, 0x86, 0xae,
	0xe2, 0x3c, 0xac, 0xc7, 0xcd, 0xa6, 0x6d, 0x57, 0x0e, 0x1a, 0x47, 0xba, 0x66, 0xa4, 0x40, 0x3b,
	0xe4, 0xdd, 0x8e, 0xb1, 0x06, 0xd9, 0x3d, 0x4e, 0xd9, 0xeb, 0x21, 0x0b, 0x84, 0x91, 0x85, 0x74,
	0x54, 0x78, 0xbd, 0x91, 0x71, 0x17, 0xb2, 0x8f, 0x98, 0xe3, 0x8b, 0x0a, 0x73, 0x44, 0x64, 0xa2,
	0xe7, 0x04, 0xa2, 0xc5, 0x1c, 0x21, 0x3f, 0x24, 0x43, 0xaf, 0xea, 0x7b, 0x9b, 0xb0, 0xd6, 0x18,
	0xb6, 0x5e, 0xb2, 0xb6, 0xa0, 0xc3, 0x1e, 0xc3, 0x19, 0xd0, 0xaa, 0xf6, 0xfe, 0x91, 0xae, 0xe0,
	0x35, 0x48, 0x37, 0xed, 0x5d, 0xfb, 0x49, 0x64, 0xc5, 0xaa, 0x43, 0xce, 0x73, 0xda, 0xaf, 0x1c,
	0x97, 0xcd, 0x86, 0xf1, 0x6d, 0x33, 0x4e, 0xd4, 0x9c, 0x27, 0x6a, 0xd6, 0xbb, 0x3d, 0x66, 0x7b,
	0xa2, 0xcb, 0x07, 0x41, 0x61, 0x7c, 0xac, 0xca, 0x1c, 0xfe, 0x50, 0x59, 0x14, 0x36, 0x92, 0x9d,
	0x7d, 0xc7, 0x77, 0xfa, 0xc1, 0x0a, 0xda, 0xb7, 0x63, 0xb5, 0xa4, 0x96, 0xb3, 0x74, 0xa9, 0xd6,
	0x72, 0x00, 0x07, 0xcc, 0x7f, 0xd3, 0x6d, 0xb3, 0xeb, 0x1f, 0xf2, 0x6f, 0xe2, 0x77, 0xe9, 0x2f,
	0xb7, 0x9d, 0x8f, 0x1f, 0xf4, 0x9a, 0x90, 0x2e, 0x81, 0x59, 0xcf, 0x20, 0xdf, 0x67, 0xe2, 0x05,
	0xef, 0xfc, 0xff, 0x86, 0x1f, 0x7f, 0xdf, 0xb0, 0xc8, 0xb2, 0x76, 0x20, 0x97, 0x5c, 0x8b, 0xef,
	0x2c, 0xd0, 0x1b, 0xf1, 0xc0, 0x7c, 0xc1, 0xd9, 0xc9, 0x2c, 0xe2, 0xa4, 0xd0, 0x3a, 0x80, 0x8d,
	0x64, 0x67, 0x16, 0xf1, 0x4a, 0xe0, 0x97, 0x93, 0x59, 0xca, 0xcb, 0xe4, 0x56, 0x1d, 0xd6, 0x13,
	0xb6, 0x31, 0x59, 0xe0, 0x3d, 0x96, 0xf7, 0x73, 0xdc, 0x87, 0xd3, 0xd8, 0x5f, 0x52, 0x66, 0x51,
	0xb8, 0x95, 0x68, 0xcc, 0xdc, 0xad, 0xa2, 0x7d, 0x3c, 0x8d, 0xcd, 0x2d, 0x13, 0x47, 0xde, 0x02,
	0xe1, 0x33, 0xa7, 0xcf, 0x3a, 0xf2, 0xef, 0x5f, 0x49, 0xfb, 0x24, 0xbd, 0x65, 0x68, 0x52, 0x56,
	0x79, 0x78, 0x3e, 0x21, 0xe8, 0x62, 0x42, 0x94, 0xcb, 0x09, 0x41, 0x3f, 0x27, 0x04, 0xfd, 0x9a,
	0x10, 0xf4, 0x36, 0x24, 0xe8, 0x73, 0x48, 0xd0, 0x59, 0x48, 0xd0, 0xd7, 0x90, 0xa0, 0xf3, 0x90,
	0xa0, 0x71, 0x48, 0xd0, 0x65, 0x48, 0xd0, 0xbb, 0x29, 0x51, 0xc6, 0x53, 0xa2, 0x5c, 0x4c, 0x89,
	0xf2, 0x34, 0x65, 0x3e, 0x88, 0x9e, 0xba, 0x95, 0x92, 0x4b, 0xef, 0xff, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0xa5, 0x5c, 0x4c, 0x85, 0x04, 0x00, 0x00,
}

func (this *Error) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Error)
	if !ok {
		that2, ok := that.(Error)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Error")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Error but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Error but is not nil && this == nil")
	}
	if this.Type != that1.Type {
		return fmt.Errorf("Type this(%v) Not Equal that(%v)", this.Type, that1.Type)
	}
	if this.Message != that1.Message {
		return fmt.Errorf("Message this(%v) Not Equal that(%v)", this.Message, that1.Message)
	}
	if this.MsgCount != that1.MsgCount {
		return fmt.Errorf("MsgCount this(%v) Not Equal that(%v)", this.MsgCount, that1.MsgCount)
	}
	return nil
}
func (this *Error) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Error)
	if !ok {
		that2, ok := that.(Error)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if this.MsgCount != that1.MsgCount {
		return false
	}
	return true
}
func (this *Void) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Void)
	if !ok {
		that2, ok := that.(Void)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Void")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Void but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Void but is not nil && this == nil")
	}
	return nil
}
func (this *Void) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Void)
	if !ok {
		that2, ok := that.(Void)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *NoRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NoRequest)
	if !ok {
		that2, ok := that.(NoRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NoRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NoRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NoRequest but is not nil && this == nil")
	}
	return nil
}
func (this *NoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NoRequest)
	if !ok {
		that2, ok := that.(NoRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *NoReply) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NoReply)
	if !ok {
		that2, ok := that.(NoReply)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NoReply")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NoReply but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NoReply but is not nil && this == nil")
	}
	return nil
}
func (this *NoReply) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NoReply)
	if !ok {
		that2, ok := that.(NoReply)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *HeartBeat) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*HeartBeat)
	if !ok {
		that2, ok := that.(HeartBeat)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *HeartBeat")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *HeartBeat but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *HeartBeat but is not nil && this == nil")
	}
	if this.Lastbeat != that1.Lastbeat {
		return fmt.Errorf("Lastbeat this(%v) Not Equal that(%v)", this.Lastbeat, that1.Lastbeat)
	}
	return nil
}
func (this *HeartBeat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeartBeat)
	if !ok {
		that2, ok := that.(HeartBeat)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Lastbeat != that1.Lastbeat {
		return false
	}
	return true
}
func (this *Error) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&nrpc.Error{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "MsgCount: "+fmt.Sprintf("%#v", this.MsgCount)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Void) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&nrpc.Void{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NoRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&nrpc.NoRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NoReply) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&nrpc.NoReply{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HeartBeat) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&nrpc.HeartBeat{")
	s = append(s, "Lastbeat: "+fmt.Sprintf("%#v", this.Lastbeat)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringNrpc(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MsgCount != 0 {
		i = encodeVarintNrpc(dAtA, i, uint64(m.MsgCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintNrpc(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintNrpc(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Void) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Void) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Void) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *NoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *NoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *HeartBeat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeartBeat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeartBeat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Lastbeat {
		i--
		if m.Lastbeat {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNrpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovNrpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedError(r randyNrpc, easy bool) *Error {
	this := &Error{}
	this.Type = Error_Type([]int32{0, 1, 3, 4}[r.Intn(4)])
	this.Message = string(randStringNrpc(r))
	this.MsgCount = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedVoid(r randyNrpc, easy bool) *Void {
	this := &Void{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNoRequest(r randyNrpc, easy bool) *NoRequest {
	this := &NoRequest{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNoReply(r randyNrpc, easy bool) *NoReply {
	this := &NoReply{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedHeartBeat(r randyNrpc, easy bool) *HeartBeat {
	this := &HeartBeat{}
	this.Lastbeat = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyNrpc interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneNrpc(r randyNrpc) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringNrpc(r randyNrpc) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneNrpc(r)
	}
	return string(tmps)
}
func randUnrecognizedNrpc(r randyNrpc, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldNrpc(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldNrpc(dAtA []byte, r randyNrpc, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateNrpc(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateNrpc(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateNrpc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateNrpc(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateNrpc(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateNrpc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateNrpc(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovNrpc(uint64(m.Type))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovNrpc(uint64(l))
	}
	if m.MsgCount != 0 {
		n += 1 + sovNrpc(uint64(m.MsgCount))
	}
	return n
}

func (m *Void) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *NoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *NoReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *HeartBeat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Lastbeat {
		n += 2
	}
	return n
}

func sovNrpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNrpc(x uint64) (n int) {
	return sovNrpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Error) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Error{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`MsgCount:` + fmt.Sprintf("%v", this.MsgCount) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Void) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Void{`,
		`}`,
	}, "")
	return s
}
func (this *NoRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NoRequest{`,
		`}`,
	}, "")
	return s
}
func (this *NoReply) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NoReply{`,
		`}`,
	}, "")
	return s
}
func (this *HeartBeat) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HeartBeat{`,
		`Lastbeat:` + fmt.Sprintf("%v", this.Lastbeat) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNrpc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Error_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgCount", wireType)
			}
			m.MsgCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNrpc
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
func (m *Void) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Void: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Void: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNrpc
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
func (m *NoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNrpc
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
func (m *NoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNrpc
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
func (m *HeartBeat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeartBeat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeartBeat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lastbeat", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Lastbeat = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNrpc
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
func skipNrpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNrpc
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
					return 0, ErrIntOverflowNrpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNrpc
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
			if length < 0 {
				return 0, ErrInvalidLengthNrpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNrpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNrpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNrpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNrpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNrpc = fmt.Errorf("proto: unexpected end of group")
)
