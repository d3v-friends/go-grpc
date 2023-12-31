// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: share.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IndexDirection int32

const (
	IndexDirection_ASC  IndexDirection = 0
	IndexDirection_DESC IndexDirection = 1
)

// Enum value maps for IndexDirection.
var (
	IndexDirection_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	IndexDirection_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x IndexDirection) Enum() *IndexDirection {
	p := new(IndexDirection)
	*p = x
	return p
}

func (x IndexDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_share_proto_enumTypes[0].Descriptor()
}

func (IndexDirection) Type() protoreflect.EnumType {
	return &file_share_proto_enumTypes[0]
}

func (x IndexDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexDirection.Descriptor instead.
func (IndexDirection) EnumDescriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{0}
}

type VerifyMode int32

const (
	VerifyMode_COMPARE VerifyMode = 0
	VerifyMode_OTP     VerifyMode = 1
)

// Enum value maps for VerifyMode.
var (
	VerifyMode_name = map[int32]string{
		0: "COMPARE",
		1: "OTP",
	}
	VerifyMode_value = map[string]int32{
		"COMPARE": 0,
		"OTP":     1,
	}
)

func (x VerifyMode) Enum() *VerifyMode {
	p := new(VerifyMode)
	*p = x
	return p
}

func (x VerifyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerifyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_share_proto_enumTypes[1].Descriptor()
}

func (VerifyMode) Type() protoreflect.EnumType {
	return &file_share_proto_enumTypes[1]
}

func (x VerifyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerifyMode.Descriptor instead.
func (VerifyMode) EnumDescriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{1}
}

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{0}
}

func (x *Time) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Decimal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Decimal) Reset() {
	*x = Decimal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Decimal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decimal) ProtoMessage() {}

func (x *Decimal) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decimal.ProtoReflect.Descriptor instead.
func (*Decimal) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{1}
}

func (x *Decimal) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ObjectID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ObjectID) Reset() {
	*x = ObjectID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectID) ProtoMessage() {}

func (x *ObjectID) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectID.ProtoReflect.Descriptor instead.
func (*ObjectID) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type IPeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Equal    *Time `protobuf:"bytes,1,opt,name=equal,proto3,oneof" json:"equal,omitempty"`
	NotEqual *Time `protobuf:"bytes,2,opt,name=notEqual,proto3,oneof" json:"notEqual,omitempty"`
	Gt       *Time `protobuf:"bytes,3,opt,name=gt,proto3,oneof" json:"gt,omitempty"`
	Gte      *Time `protobuf:"bytes,4,opt,name=gte,proto3,oneof" json:"gte,omitempty"`
	Lt       *Time `protobuf:"bytes,5,opt,name=lt,proto3,oneof" json:"lt,omitempty"`
	Lte      *Time `protobuf:"bytes,6,opt,name=lte,proto3,oneof" json:"lte,omitempty"`
}

func (x *IPeriod) Reset() {
	*x = IPeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPeriod) ProtoMessage() {}

func (x *IPeriod) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPeriod.ProtoReflect.Descriptor instead.
func (*IPeriod) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{3}
}

func (x *IPeriod) GetEqual() *Time {
	if x != nil {
		return x.Equal
	}
	return nil
}

func (x *IPeriod) GetNotEqual() *Time {
	if x != nil {
		return x.NotEqual
	}
	return nil
}

func (x *IPeriod) GetGt() *Time {
	if x != nil {
		return x.Gt
	}
	return nil
}

func (x *IPeriod) GetGte() *Time {
	if x != nil {
		return x.Gte
	}
	return nil
}

func (x *IPeriod) GetLt() *Time {
	if x != nil {
		return x.Lt
	}
	return nil
}

func (x *IPeriod) GetLte() *Time {
	if x != nil {
		return x.Lte
	}
	return nil
}

type IPager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *IPager) Reset() {
	*x = IPager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPager) ProtoMessage() {}

func (x *IPager) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPager.ProtoReflect.Descriptor instead.
func (*IPager) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{4}
}

func (x *IPager) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *IPager) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_share_proto protoreflect.FileDescriptor

var file_share_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a, 0x07, 0x44,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a, 0x08,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xfc,
	0x01, 0x0a, 0x07, 0x49, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x05, 0x65, 0x71,
	0x75, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x08,
	0x6e, 0x6f, 0x74, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x01, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x45, 0x71, 0x75, 0x61,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x02, 0x67, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x02, 0x52, 0x02, 0x67, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1c, 0x0a, 0x03, 0x67, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x48, 0x03, 0x52, 0x03, 0x67, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a,
	0x0a, 0x02, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x48, 0x04, 0x52, 0x02, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x03, 0x6c, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x05,
	0x52, 0x03, 0x6c, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x71, 0x75,
	0x61, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x6f, 0x74, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x67, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x67, 0x74, 0x65, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x6c, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6c, 0x74, 0x65, 0x22, 0x30, 0x0a,
	0x06, 0x49, 0x50, 0x61, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x2a,
	0x23, 0x0a, 0x0e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45,
	0x53, 0x43, 0x10, 0x01, 0x2a, 0x22, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x45, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x4f, 0x54, 0x50, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_share_proto_rawDescOnce sync.Once
	file_share_proto_rawDescData = file_share_proto_rawDesc
)

func file_share_proto_rawDescGZIP() []byte {
	file_share_proto_rawDescOnce.Do(func() {
		file_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_share_proto_rawDescData)
	})
	return file_share_proto_rawDescData
}

var file_share_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_share_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_share_proto_goTypes = []interface{}{
	(IndexDirection)(0), // 0: IndexDirection
	(VerifyMode)(0),     // 1: VerifyMode
	(*Time)(nil),        // 2: Time
	(*Decimal)(nil),     // 3: Decimal
	(*ObjectID)(nil),    // 4: ObjectID
	(*IPeriod)(nil),     // 5: IPeriod
	(*IPager)(nil),      // 6: IPager
}
var file_share_proto_depIdxs = []int32{
	2, // 0: IPeriod.equal:type_name -> Time
	2, // 1: IPeriod.notEqual:type_name -> Time
	2, // 2: IPeriod.gt:type_name -> Time
	2, // 3: IPeriod.gte:type_name -> Time
	2, // 4: IPeriod.lt:type_name -> Time
	2, // 5: IPeriod.lte:type_name -> Time
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_share_proto_init() }
func file_share_proto_init() {
	if File_share_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_share_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Time); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_share_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Decimal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_share_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_share_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPeriod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_share_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPager); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_share_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_share_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_share_proto_goTypes,
		DependencyIndexes: file_share_proto_depIdxs,
		EnumInfos:         file_share_proto_enumTypes,
		MessageInfos:      file_share_proto_msgTypes,
	}.Build()
	File_share_proto = out.File
	file_share_proto_rawDesc = nil
	file_share_proto_goTypes = nil
	file_share_proto_depIdxs = nil
}
