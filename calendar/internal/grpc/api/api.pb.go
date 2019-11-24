// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	EventAlertBefore_UNKNOWN EventAlertBefore_BeforeType = 0
	EventAlertBefore_DAYS    EventAlertBefore_BeforeType = 1
	EventAlertBefore_HOURS   EventAlertBefore_BeforeType = 2
	EventAlertBefore_MINUTES EventAlertBefore_BeforeType = 3
)

var EventAlertBefore_BeforeType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DAYS",
	2: "HOURS",
	3: "MINUTES",
}

var EventAlertBefore_BeforeType_value = map[string]int32{
	"UNKNOWN": 0,
	"DAYS":    1,
	"HOURS":   2,
	"MINUTES": 3,
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
	Duration             *duration.Duration   `protobuf:"bytes,8,opt,name=duration,proto3" json:"duration,omitempty"`
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

func (m *Event) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
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
	return EventAlertBefore_UNKNOWN
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

type ID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Day struct {
	Day                  *timestamp.Timestamp `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Day) Reset()         { *m = Day{} }
func (m *Day) String() string { return proto.CompactTextString(m) }
func (*Day) ProtoMessage()    {}
func (*Day) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Day) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Day.Unmarshal(m, b)
}
func (m *Day) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Day.Marshal(b, m, deterministic)
}
func (m *Day) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Day.Merge(m, src)
}
func (m *Day) XXX_Size() int {
	return xxx_messageInfo_Day.Size(m)
}
func (m *Day) XXX_DiscardUnknown() {
	xxx_messageInfo_Day.DiscardUnknown(m)
}

var xxx_messageInfo_Day proto.InternalMessageInfo

func (m *Day) GetDay() *timestamp.Timestamp {
	if m != nil {
		return m.Day
	}
	return nil
}

type EventRequest struct {
	OccursAt             *timestamp.Timestamp `protobuf:"bytes,1,opt,name=occurs_at,json=occursAt,proto3" json:"occurs_at,omitempty"`
	Subject              string               `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string               `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Location             string               `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	UserID               string               `protobuf:"bytes,6,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetOccursAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccursAt
	}
	return nil
}

func (m *EventRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EventRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *EventRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *EventRequest) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *EventRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type EventResponse struct {
	// Types that are valid to be assigned to Result:
	//	*EventResponse_Event
	//	*EventResponse_Error
	Result               isEventResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

type isEventResponse_Result interface {
	isEventResponse_Result()
}

type EventResponse_Event struct {
	Event *Event `protobuf:"bytes,1,opt,name=event,proto3,oneof"`
}

type EventResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*EventResponse_Event) isEventResponse_Result() {}

func (*EventResponse_Error) isEventResponse_Result() {}

func (m *EventResponse) GetResult() isEventResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *EventResponse) GetEvent() *Event {
	if x, ok := m.GetResult().(*EventResponse_Event); ok {
		return x.Event
	}
	return nil
}

func (m *EventResponse) GetError() string {
	if x, ok := m.GetResult().(*EventResponse_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventResponse_Event)(nil),
		(*EventResponse_Error)(nil),
	}
}

type EventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventsResponse) Reset()         { *m = EventsResponse{} }
func (m *EventsResponse) String() string { return proto.CompactTextString(m) }
func (*EventsResponse) ProtoMessage()    {}
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *EventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsResponse.Unmarshal(m, b)
}
func (m *EventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsResponse.Marshal(b, m, deterministic)
}
func (m *EventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsResponse.Merge(m, src)
}
func (m *EventsResponse) XXX_Size() int {
	return xxx_messageInfo_EventsResponse.Size(m)
}
func (m *EventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventsResponse proto.InternalMessageInfo

func (m *EventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.EventAlertBefore_BeforeType", EventAlertBefore_BeforeType_name, EventAlertBefore_BeforeType_value)
	proto.RegisterType((*Event)(nil), "api.Event")
	proto.RegisterType((*EventAlertBefore)(nil), "api.Event.alertBefore")
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*ID)(nil), "api.ID")
	proto.RegisterType((*Day)(nil), "api.Day")
	proto.RegisterType((*EventRequest)(nil), "api.EventRequest")
	proto.RegisterType((*EventResponse)(nil), "api.EventResponse")
	proto.RegisterType((*EventsResponse)(nil), "api.EventsResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xad, 0x63, 0x27, 0x71, 0x6e, 0x68, 0x9b, 0x0e, 0x15, 0x32, 0x59, 0xf0, 0xf0, 0x0a, 0xa4,
	0xca, 0x95, 0xd2, 0x02, 0x42, 0xac, 0xda, 0xa6, 0xd0, 0x0a, 0x35, 0x95, 0x9c, 0x5a, 0x15, 0x2b,
	0xe4, 0xc7, 0xb4, 0x35, 0x38, 0x1e, 0x33, 0x1e, 0x57, 0xca, 0x5f, 0xb0, 0xe5, 0x17, 0xf8, 0x00,
	0xbe, 0x87, 0x4f, 0x61, 0x1e, 0x76, 0xe3, 0x3e, 0x90, 0x0b, 0x2b, 0xcf, 0xbd, 0x73, 0xce, 0xb9,
	0x77, 0xe6, 0x1e, 0x0f, 0xf4, 0xfc, 0x2c, 0x76, 0x32, 0x4a, 0x18, 0x41, 0x3a, 0x5f, 0x0e, 0x9f,
	0x9e, 0x13, 0x72, 0x9e, 0xe0, 0x4d, 0x99, 0x0a, 0x8a, 0xb3, 0x4d, 0x16, 0xcf, 0x70, 0xce, 0xfc,
	0x59, 0xa6, 0x50, 0xc3, 0x27, 0x37, 0x01, 0x51, 0x41, 0x7d, 0x16, 0x93, 0x54, 0xed, 0xdb, 0xbf,
	0x0c, 0x68, 0xef, 0x5f, 0xe2, 0x94, 0xa1, 0x15, 0x68, 0xc5, 0x91, 0xa5, 0x3d, 0xd3, 0x5e, 0xf4,
	0x5c, 0xbe, 0x42, 0x6f, 0x01, 0x42, 0x8a, 0x7d, 0x86, 0xa3, 0xcf, 0x3e, 0xb3, 0x5a, 0x3c, 0xdf,
	0x1f, 0x0d, 0x1d, 0x25, 0xe7, 0x54, 0x72, 0xce, 0x49, 0x55, 0xcf, 0xed, 0x95, 0xe8, 0x1d, 0x26,
	0xa8, 0x45, 0x16, 0x55, 0x54, 0xbd, 0x99, 0x5a, 0xa2, 0x15, 0x35, 0xc2, 0x09, 0x2e, 0xa9, 0x46,
	0x33, 0xb5, 0x44, 0x73, 0xea, 0x1b, 0xe8, 0x91, 0x30, 0x2c, 0x68, 0x2e, 0x98, 0xed, 0x46, 0xa6,
	0xa9, 0xc0, 0x9c, 0x68, 0x41, 0x37, 0x2f, 0x82, 0x2f, 0x38, 0x64, 0x56, 0x47, 0x1e, 0xbf, 0x0a,
	0x11, 0x02, 0x23, 0x20, 0xd1, 0xdc, 0xea, 0xca, 0xb4, 0x5c, 0xa3, 0x57, 0x60, 0x56, 0x77, 0x68,
	0x99, 0xb2, 0xca, 0xe3, 0x5b, 0x55, 0xc6, 0x25, 0xc0, 0xbd, 0x82, 0xa2, 0x21, 0x98, 0x09, 0x09,
	0x15, 0xad, 0x27, 0xe5, 0xae, 0x62, 0xf4, 0x08, 0x3a, 0x45, 0x8e, 0xe9, 0xe1, 0xd8, 0x02, 0xb9,
	0x53, 0x46, 0xc3, 0x1f, 0x1a, 0xf4, 0xfd, 0x04, 0x53, 0xb6, 0x8b, 0xcf, 0x08, 0xc5, 0xbc, 0xb4,
	0xc1, 0xe6, 0x19, 0x96, 0x43, 0x5a, 0x19, 0x3d, 0x77, 0x84, 0x19, 0xe4, 0xf0, 0x9c, 0x1a, 0xca,
	0x51, 0x9f, 0x13, 0x0e, 0x74, 0x25, 0x5c, 0xc8, 0x07, 0x32, 0x27, 0xa7, 0xa8, 0xbb, 0x65, 0x64,
	0xbf, 0x03, 0x58, 0x60, 0x51, 0x1f, 0xba, 0xde, 0xe4, 0xe3, 0xe4, 0xf8, 0x74, 0x32, 0x58, 0x42,
	0x26, 0x18, 0xe3, 0x9d, 0x4f, 0xd3, 0x81, 0x86, 0x7a, 0xd0, 0x3e, 0x38, 0xf6, 0xdc, 0xe9, 0xa0,
	0x25, 0x10, 0x47, 0x87, 0x13, 0xef, 0x64, 0x7f, 0x3a, 0xd0, 0xed, 0xef, 0x1a, 0x18, 0x1e, 0x6f,
	0xf3, 0x96, 0x6f, 0xf8, 0x9d, 0xa5, 0xfe, 0x4c, 0xd5, 0xe2, 0x77, 0x26, 0xd6, 0x68, 0x1d, 0xda,
	0x78, 0xe6, 0xc7, 0x09, 0xf7, 0x82, 0xce, 0x93, 0x2a, 0x10, 0x7d, 0xcd, 0x48, 0x10, 0x27, 0x98,
	0xcf, 0x59, 0xa4, 0xcb, 0x08, 0xbd, 0x06, 0x33, 0x88, 0x29, 0xbb, 0x88, 0xfc, 0xf9, 0x7d, 0xe6,
	0x58, 0x61, 0xed, 0x75, 0x68, 0x1d, 0x8e, 0x6f, 0xf6, 0x63, 0x6f, 0x81, 0x3e, 0xf6, 0xe7, 0x68,
	0x03, 0x74, 0xa1, 0xa7, 0x35, 0xea, 0x09, 0x98, 0xfd, 0x5b, 0x83, 0x07, 0xf2, 0x66, 0x5d, 0xfc,
	0xad, 0xe0, 0x1b, 0xd7, 0xcd, 0xa5, 0xfd, 0x9f, 0xb9, 0x5a, 0x77, 0x9b, 0x4b, 0xaf, 0x99, 0xab,
	0xee, 0x12, 0xe3, 0x86, 0x4b, 0xea, 0xc6, 0x6b, 0xdf, 0xdf, 0x78, 0x0b, 0x73, 0x75, 0xea, 0xe6,
	0xb2, 0x3d, 0x58, 0x2e, 0x4f, 0x98, 0x67, 0x24, 0xcd, 0x31, 0xb2, 0xf9, 0x90, 0x44, 0xa2, 0x3c,
	0x1e, 0x2c, 0xec, 0x75, 0xb0, 0xe4, 0xaa, 0x2d, 0x2e, 0xd6, 0xc6, 0x94, 0x12, 0xaa, 0xce, 0x22,
	0xf3, 0x22, 0xdc, 0x35, 0xa1, 0x43, 0x71, 0x5e, 0x24, 0xcc, 0xde, 0x86, 0x15, 0xc9, 0xc9, 0x6b,
	0xba, 0x1d, 0x49, 0xce, 0xb9, 0xb0, 0x7e, 0x5d, 0xd8, 0x2d, 0x77, 0x46, 0x3f, 0x75, 0x58, 0xdd,
	0xe3, 0x26, 0x4e, 0x23, 0x9f, 0x4e, 0x31, 0xbd, 0x8c, 0x43, 0x61, 0x83, 0xfe, 0x9e, 0x7c, 0x52,
	0xd4, 0xfb, 0xb4, 0x56, 0xa3, 0xa9, 0xa1, 0x0c, 0x51, 0x3d, 0xa5, 0xaa, 0xd9, 0x4b, 0xe8, 0x25,
	0x98, 0x1f, 0x30, 0x53, 0xa4, 0xae, 0x44, 0xf0, 0x5f, 0xe9, 0x6e, 0xe8, 0x26, 0x2c, 0x73, 0xa8,
	0xb0, 0xb1, 0xea, 0x79, 0x81, 0x7f, 0xb8, 0xc0, 0xe7, 0x35, 0x02, 0xef, 0xc9, 0x93, 0x6f, 0xd5,
	0x3f, 0xf6, 0xb4, 0x01, 0xfd, 0xb1, 0x7c, 0xa8, 0xee, 0xd5, 0xd6, 0x08, 0x56, 0xab, 0x13, 0xe4,
	0xef, 0x09, 0x15, 0xf6, 0x35, 0x25, 0x90, 0xaf, 0xfe, 0xd6, 0xd9, 0x16, 0x0c, 0xea, 0x9c, 0x53,
	0x8c, 0xbf, 0x36, 0x93, 0xb6, 0x61, 0xad, 0x4e, 0x3a, 0x22, 0x29, 0xbb, 0x68, 0x64, 0x05, 0x1d,
	0x69, 0xb7, 0xad, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x3a, 0x0c, 0x62, 0x8d, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarServiceClient is the client API for CalendarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarServiceClient interface {
	CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	GetEvent(ctx context.Context, in *ID, opts ...grpc.CallOption) (*EventResponse, error)
	GetUserEvents(ctx context.Context, in *ID, opts ...grpc.CallOption) (*EventsResponse, error)
	UpdateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	DeleteEvent(ctx context.Context, in *ID, opts ...grpc.CallOption) (*EventResponse, error)
	GetEventsForDay(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error)
	GetEventsForWeek(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error)
	GetEventsForMonth(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error)
}

type calendarServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalendarServiceClient(cc *grpc.ClientConn) CalendarServiceClient {
	return &calendarServiceClient{cc}
}

func (c *calendarServiceClient) CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/api.CalendarService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEvent(ctx context.Context, in *ID, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/api.CalendarService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetUserEvents(ctx context.Context, in *ID, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/api.CalendarService/GetUserEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) UpdateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/api.CalendarService/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) DeleteEvent(ctx context.Context, in *ID, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/api.CalendarService/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEventsForDay(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/api.CalendarService/GetEventsForDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEventsForWeek(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/api.CalendarService/GetEventsForWeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEventsForMonth(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/api.CalendarService/GetEventsForMonth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServiceServer is the server API for CalendarService service.
type CalendarServiceServer interface {
	CreateEvent(context.Context, *EventRequest) (*EventResponse, error)
	GetEvent(context.Context, *ID) (*EventResponse, error)
	GetUserEvents(context.Context, *ID) (*EventsResponse, error)
	UpdateEvent(context.Context, *EventRequest) (*EventResponse, error)
	DeleteEvent(context.Context, *ID) (*EventResponse, error)
	GetEventsForDay(context.Context, *Day) (*EventsResponse, error)
	GetEventsForWeek(context.Context, *Day) (*EventsResponse, error)
	GetEventsForMonth(context.Context, *Day) (*EventsResponse, error)
}

// UnimplementedCalendarServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServiceServer struct {
}

func (*UnimplementedCalendarServiceServer) CreateEvent(ctx context.Context, req *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) GetEvent(ctx context.Context, req *ID) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) GetUserEvents(ctx context.Context, req *ID) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEvents not implemented")
}
func (*UnimplementedCalendarServiceServer) UpdateEvent(ctx context.Context, req *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) DeleteEvent(ctx context.Context, req *ID) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) GetEventsForDay(ctx context.Context, req *Day) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForDay not implemented")
}
func (*UnimplementedCalendarServiceServer) GetEventsForWeek(ctx context.Context, req *Day) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForWeek not implemented")
}
func (*UnimplementedCalendarServiceServer) GetEventsForMonth(ctx context.Context, req *Day) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForMonth not implemented")
}

func RegisterCalendarServiceServer(s *grpc.Server, srv CalendarServiceServer) {
	s.RegisterService(&_CalendarService_serviceDesc, srv)
}

func _CalendarService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalendarService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).CreateEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalendarService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEvent(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetUserEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetUserEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalendarService/GetUserEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetUserEvents(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalendarService/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).UpdateEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalendarService/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEventsForDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEventsForDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalendarService/GetEventsForDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEventsForDay(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEventsForWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEventsForWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalendarService/GetEventsForWeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEventsForWeek(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEventsForMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEventsForMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalendarService/GetEventsForMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEventsForMonth(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalendarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CalendarService",
	HandlerType: (*CalendarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _CalendarService_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _CalendarService_GetEvent_Handler,
		},
		{
			MethodName: "GetUserEvents",
			Handler:    _CalendarService_GetUserEvents_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _CalendarService_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _CalendarService_DeleteEvent_Handler,
		},
		{
			MethodName: "GetEventsForDay",
			Handler:    _CalendarService_GetEventsForDay_Handler,
		},
		{
			MethodName: "GetEventsForWeek",
			Handler:    _CalendarService_GetEventsForWeek_Handler,
		},
		{
			MethodName: "GetEventsForMonth",
			Handler:    _CalendarService_GetEventsForMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
