// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_getCaptcha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Pinx                 []uint32 `protobuf:"varint,3,rep,packed,name=pinx,proto3" json:"pinx,omitempty"`
	Stride               int64    `protobuf:"varint,4,opt,name=stride,proto3" json:"stride,omitempty"`
	Min                  *Point   `protobuf:"bytes,5,opt,name=min,proto3" json:"min,omitempty"`
	Max                  *Point   `protobuf:"bytes,6,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetPinx() []uint32 {
	if m != nil {
		return m.Pinx
	}
	return nil
}

func (m *Response) GetStride() int64 {
	if m != nil {
		return m.Stride
	}
	return 0
}

func (m *Response) GetMin() *Point {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *Response) GetMax() *Point {
	if m != nil {
		return m.Max
	}
	return nil
}

type Point struct {
	X                    int64    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int64    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.getCaptcha.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.getCaptcha.Response")
	proto.RegisterType((*Point)(nil), "go.micro.srv.getCaptcha.Point")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0xc4, 0x20,
	0x10, 0x86, 0x45, 0x76, 0xbb, 0x3a, 0xea, 0x85, 0x18, 0x25, 0x1a, 0x0d, 0xd6, 0x4b, 0x4f, 0x68,
	0xd6, 0x47, 0x30, 0xc6, 0xab, 0xc1, 0xbb, 0x09, 0xee, 0x4e, 0x2a, 0x89, 0x05, 0x04, 0x6a, 0xd8,
	0x37, 0xf4, 0xb1, 0x4c, 0x69, 0xf5, 0xb6, 0xc6, 0x13, 0xf3, 0x4d, 0xfe, 0x19, 0x7e, 0x7e, 0xe0,
	0xdc, 0x07, 0x97, 0xdc, 0x0d, 0x66, 0xdd, 0xf9, 0x77, 0xfc, 0x39, 0x65, 0xe9, 0xb2, 0xd3, 0xd6,
	0xc9, 0xce, 0xac, 0x82, 0x93, 0x31, 0x7c, 0xca, 0x16, 0xd3, 0xbd, 0xf6, 0x69, 0xf5, 0xa6, 0xeb,
	0x0b, 0x58, 0x28, 0xfc, 0xe8, 0x31, 0x26, 0xc6, 0x60, 0xd6, 0xf7, 0x66, 0xcd, 0x89, 0x20, 0xcd,
	0xbe, 0x2a, 0x75, 0xfd, 0x45, 0x60, 0x4f, 0x61, 0xf4, 0xce, 0x46, 0x64, 0xc7, 0x30, 0xc7, 0x10,
	0xac, 0x9b, 0x14, 0x23, 0xb0, 0x13, 0xa8, 0x30, 0x84, 0x2e, 0xb6, 0x7c, 0xb7, 0xb4, 0x27, 0x1a,
	0xd6, 0x79, 0x63, 0x33, 0xa7, 0x82, 0x36, 0x47, 0xaa, 0xd4, 0x83, 0x36, 0xa6, 0x60, 0xd6, 0xc8,
	0x67, 0x82, 0x34, 0x54, 0x4d, 0xc4, 0x6e, 0x81, 0x76, 0xc6, 0xf2, 0xb9, 0x20, 0xcd, 0xc1, 0xf2,
	0x52, 0x6e, 0x31, 0x2b, 0x9f, 0x9c, 0xb1, 0x49, 0x0d, 0xd2, 0x32, 0xa1, 0x33, 0xaf, 0xfe, 0x39,
	0xa1, 0x73, 0x7d, 0x0d, 0xf3, 0x42, 0xec, 0x10, 0x48, 0x2e, 0x4f, 0xa0, 0x8a, 0xe4, 0x81, 0x36,
	0xc5, 0x39, 0x55, 0x64, 0xb3, 0x7c, 0x81, 0xc5, 0xc3, 0x18, 0x1c, 0x7b, 0x06, 0x78, 0xfc, 0x5d,
	0xc4, 0xc4, 0xd6, 0x2b, 0xa6, 0xf8, 0xce, 0xae, 0xfe, 0x50, 0x8c, 0x01, 0xd6, 0x3b, 0xaf, 0x55,
	0xf9, 0x8e, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x69, 0xc7, 0x77, 0xad, 0x01, 0x00,
	0x00,
}
