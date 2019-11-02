// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventAlertBefore_BeforeType int32

const (
	EventAlertBefore_DAYS    EventAlertBefore_BeforeType = 0
	EventAlertBefore_HOURS   EventAlertBefore_BeforeType = 1
	EventAlertBefore_MINUTES EventAlertBefore_BeforeType = 2
)

var EventAlertBefore_BeforeType_name = map[int32]string{
	0: "DAYS",
	1: "HOURS",
	2: "MINUTES",
}

var EventAlertBefore_BeforeType_value = map[string]int32{
	"DAYS":    0,
	"HOURS":   1,
	"MINUTES": 2,
}

func (x EventAlertBefore_BeforeType) String() string {
	return proto.EnumName(EventAlertBefore_BeforeType_name, int32(x))
}

func (EventAlertBefore_BeforeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0, 0, 0}
}

type Event struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	OccursAt             *timestamp.Timestamp `protobuf:"bytes,5,opt,name=occurs_at,json=occursAt,proto3" json:"occurs_at,omitempty"`
	Subject              string               `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string               `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	Duration             int64                `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	Location             string               `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	UserID               string               `protobuf:"bytes,10,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Event) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Event) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Event) GetOccursAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccursAt
	}
	return nil
}

func (m *Event) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Event) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Event) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Event) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Event) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type EventAlertBefore struct {
	Type                 EventAlertBefore_BeforeType `protobuf:"varint,1,opt,name=type,proto3,enum=api.EventAlertBefore_BeforeType" json:"type,omitempty"`
	Before               int64                       `protobuf:"varint,2,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *EventAlertBefore) Reset()         { *m = EventAlertBefore{} }
func (m *EventAlertBefore) String() string { return proto.CompactTextString(m) }
func (*EventAlertBefore) ProtoMessage()    {}
func (*EventAlertBefore) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0, 0}
}

func (m *EventAlertBefore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventAlertBefore.Unmarshal(m, b)
}
func (m *EventAlertBefore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventAlertBefore.Marshal(b, m, deterministic)
}
func (m *EventAlertBefore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAlertBefore.Merge(m, src)
}
func (m *EventAlertBefore) XXX_Size() int {
	return xxx_messageInfo_EventAlertBefore.Size(m)
}
func (m *EventAlertBefore) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAlertBefore.DiscardUnknown(m)
}

var xxx_messageInfo_EventAlertBefore proto.InternalMessageInfo

func (m *EventAlertBefore) GetType() EventAlertBefore_BeforeType {
	if m != nil {
		return m.Type
	}
	return EventAlertBefore_DAYS
}

func (m *EventAlertBefore) GetBefore() int64 {
	if m != nil {
		return m.Before
	}
	return 0
}

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                []string             `protobuf:"bytes,3,rep,name=email,proto3" json:"email,omitempty"`
	Mobile               []string             `protobuf:"bytes,4,rep,name=mobile,proto3" json:"mobile,omitempty"`
	Birthday             *timestamp.Timestamp `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() []string {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *User) GetMobile() []string {
	if m != nil {
		return m.Mobile
	}
	return nil
}

func (m *User) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.EventAlertBefore_BeforeType", EventAlertBefore_BeforeType_name, EventAlertBefore_BeforeType_value)
	proto.RegisterType((*Event)(nil), "api.Event")
	proto.RegisterType((*EventAlertBefore)(nil), "api.Event.alertBefore")
	proto.RegisterType((*User)(nil), "api.User")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x4b, 0xeb, 0x40,
	0x18, 0x7c, 0x69, 0x92, 0x36, 0xf9, 0x0a, 0xa5, 0x2c, 0x8f, 0x47, 0xc8, 0xe5, 0x69, 0x4f, 0x9e,
	0xb6, 0x50, 0x51, 0xf1, 0x58, 0x69, 0xc1, 0x1e, 0x54, 0x48, 0xdb, 0x83, 0x27, 0xd9, 0x24, 0xdb,
	0x1a, 0x49, 0xba, 0x61, 0xb3, 0x11, 0xfa, 0x3f, 0x88, 0xf8, 0x27, 0xbb, 0x3f, 0x92, 0x2a, 0x78,
	0xa8, 0xa7, 0x7c, 0x33, 0xdf, 0xcc, 0xb0, 0xd9, 0x59, 0xf0, 0x49, 0x99, 0xe1, 0x92, 0x33, 0xc1,
	0x90, 0x2d, 0xc7, 0xf0, 0xff, 0x96, 0xb1, 0x6d, 0x4e, 0xc7, 0x9a, 0x8a, 0xeb, 0xcd, 0x58, 0x64,
	0x05, 0xad, 0x04, 0x29, 0x4a, 0xa3, 0x1a, 0xbd, 0x3b, 0xe0, 0xce, 0x5f, 0xe9, 0x4e, 0xa0, 0x01,
	0x74, 0xb2, 0x34, 0xb0, 0x4e, 0xac, 0x33, 0x3f, 0x92, 0x13, 0xba, 0x06, 0x48, 0x38, 0x25, 0x82,
	0xa6, 0x4f, 0x44, 0x04, 0x1d, 0xc9, 0xf7, 0x27, 0x21, 0x36, 0x79, 0xb8, 0xcd, 0xc3, 0xab, 0x36,
	0x2f, 0xf2, 0x1b, 0xf5, 0x54, 0x28, 0x6b, 0x5d, 0xa6, 0xad, 0xd5, 0x3e, 0x6e, 0x6d, 0xd4, 0xc6,
	0x9a, 0xd2, 0x9c, 0x36, 0x56, 0xe7, 0xb8, 0xb5, 0x51, 0x4b, 0xeb, 0x15, 0xf8, 0x2c, 0x49, 0x6a,
	0x5e, 0x29, 0xa7, 0x7b, 0xd4, 0xe9, 0x19, 0xb1, 0x34, 0x06, 0xd0, 0xab, 0xea, 0xf8, 0x85, 0x26,
	0x22, 0xe8, 0xea, 0xdf, 0x6f, 0x21, 0x42, 0xe0, 0xc4, 0x2c, 0xdd, 0x07, 0x3d, 0x4d, 0xeb, 0x19,
	0x85, 0xe0, 0xa5, 0x35, 0x27, 0x22, 0x63, 0xbb, 0xc0, 0x93, 0xbc, 0x1d, 0x1d, 0xb0, 0xda, 0xe5,
	0x2c, 0x31, 0x3b, 0x5f, 0x7b, 0x0e, 0x18, 0xfd, 0x83, 0x6e, 0x5d, 0x51, 0xbe, 0x98, 0x05, 0xa0,
	0x37, 0x0d, 0x0a, 0xdf, 0x2c, 0xe8, 0x93, 0x9c, 0x72, 0x71, 0x43, 0x37, 0x8c, 0x53, 0x74, 0x01,
	0x8e, 0xd8, 0x97, 0x54, 0x37, 0x31, 0x98, 0x9c, 0x62, 0xd5, 0xa8, 0x6e, 0x08, 0x7f, 0x53, 0x61,
	0xf3, 0x59, 0x49, 0x61, 0xa4, 0xe5, 0x2a, 0x3e, 0xd6, 0x9c, 0xae, 0xca, 0x8e, 0x1a, 0x34, 0xc2,
	0x00, 0x5f, 0x5a, 0xe4, 0x81, 0x33, 0x9b, 0x3e, 0x2e, 0x87, 0x7f, 0x90, 0x0f, 0xee, 0xed, 0xc3,
	0x3a, 0x5a, 0x0e, 0x2d, 0xd4, 0x87, 0xde, 0xdd, 0xe2, 0x7e, 0xbd, 0x9a, 0x2f, 0x87, 0x9d, 0xd1,
	0x87, 0x05, 0xce, 0x5a, 0x9e, 0xec, 0xc7, 0x7b, 0x90, 0x77, 0xb1, 0x23, 0x85, 0x89, 0x97, 0x77,
	0xa1, 0x66, 0xf4, 0x17, 0x5c, 0x5a, 0x90, 0x2c, 0x97, 0x1d, 0xdb, 0x92, 0x34, 0x40, 0x1d, 0xa5,
	0x60, 0x71, 0x96, 0x53, 0xd9, 0x9f, 0xa2, 0x1b, 0x84, 0x2e, 0xc1, 0x8b, 0x33, 0x2e, 0x9e, 0x53,
	0xb2, 0xff, 0x4d, 0x3f, 0xad, 0x36, 0xee, 0xea, 0xed, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xee, 0xe0, 0xdc, 0x3f, 0xdd, 0x02, 0x00, 0x00,
}
