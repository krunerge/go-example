// Code generated by protoc-gen-go.
// source: addressbook.proto
// DO NOT EDIT!

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package tutorial

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// [START messages]
type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0x84, 0x40,
	0x14, 0xc6, 0xd3, 0x75, 0x65, 0xf7, 0x09, 0x8b, 0x3b, 0x2c, 0x21, 0x4b, 0x87, 0x45, 0x3a, 0x08,
	0xc1, 0x1c, 0xb6, 0xa0, 0x73, 0x82, 0x54, 0xd4, 0xa6, 0x0c, 0x45, 0x67, 0xcd, 0xa9, 0x24, 0xf5,
	0x89, 0xa3, 0x50, 0xff, 0x52, 0x7f, 0x63, 0x87, 0xc6, 0xd1, 0x16, 0x89, 0xbd, 0x7d, 0xdf, 0x9b,
	0x9f, 0x9f, 0xf3, 0xbe, 0x81, 0x65, 0x9c, 0xa6, 0x35, 0x17, 0x22, 0x41, 0xfc, 0xa0, 0x55, 0x8d,
	0x0d, 0x92, 0x59, 0xd3, 0x36, 0x58, 0x67, 0x71, 0xee, 0xfe, 0x68, 0x60, 0x46, 0xbc, 0x16, 0x58,
	0x12, 0x02, 0x46, 0x19, 0x17, 0xdc, 0xd1, 0x36, 0x9a, 0x37, 0x67, 0x4a, 0x93, 0x05, 0xe8, 0x59,
	0xea, 0xe8, 0x72, 0x32, 0x65, 0x52, 0x91, 0x15, 0x4c, 0x79, 0x11, 0x67, 0xb9, 0x33, 0x51, 0x50,
	0x6f, 0xc8, 0x05, 0x98, 0xd5, 0x3b, 0x96, 0x5c, 0x38, 0xc6, 0x66, 0xe2, 0x59, 0xdb, 0x13, 0xfa,
	0x97, 0x4f, 0xfb, 0x6c, 0x1a, 0x75, 0xc7, 0x0f, 0x6d, 0x91, 0xf0, 0x9a, 0x0d, 0xec, 0xfa, 0x09,
	0xac, 0xd1, 0x98, 0x1c, 0x83, 0x59, 0x2a, 0x35, 0x5c, 0x60, 0x70, 0x84, 0x82, 0xd1, 0x7c, 0x55,
	0x5c, 0x5d, 0x62, 0xb1, 0x5d, 0x1f, 0x8e, 0x7e, 0x94, 0x04, 0x53, 0x9c, 0x7b, 0x06, 0xf3, 0xfd,
	0x88, 0x00, 0x98, 0xbb, 0xd0, 0xbf, 0xbd, 0x0f, 0xec, 0x23, 0x32, 0x03, 0xe3, 0x26, 0xdc, 0x05,
	0xb6, 0xd6, 0xa9, 0xe7, 0x90, 0xdd, 0xd9, 0xba, 0x7b, 0x09, 0xd6, 0x55, 0xdf, 0x8e, 0x2f, 0xdb,
	0x21, 0x9e, 0x5c, 0x84, 0x63, 0x95, 0x77, 0x25, 0x74, 0x8b, 0xd8, 0xff, 0xff, 0xc6, 0x86, 0x73,
	0x3f, 0x82, 0xd5, 0x0b, 0x16, 0x94, 0x7f, 0xc6, 0x85, 0xb4, 0x7b, 0xcc, 0x5f, 0x8e, 0xe2, 0xa2,
	0xae, 0x6b, 0xf1, 0xad, 0x9f, 0x5e, 0x23, 0xbe, 0x49, 0x48, 0xf9, 0xa4, 0x7d, 0xa5, 0x41, 0xff,
	0x95, 0xa0, 0x23, 0x38, 0x31, 0xd5, 0xd3, 0x9c, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x30, 0xfb,
	0x3f, 0x9d, 0xaf, 0x01, 0x00, 0x00,
}
