// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: config/log.proto

package config

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

type Rotate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	MaxAge     uint32 `protobuf:"varint,2,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	MaxSize    uint32 `protobuf:"varint,3,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	MaxBackups uint32 `protobuf:"varint,4,opt,name=max_backups,json=maxBackups,proto3" json:"max_backups,omitempty"`
	LocalTime  bool   `protobuf:"varint,5,opt,name=local_time,json=localTime,proto3" json:"local_time,omitempty"`
	Compress   bool   `protobuf:"varint,6,opt,name=compress,proto3" json:"compress,omitempty"`
}

func (x *Rotate) Reset() {
	*x = Rotate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rotate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rotate) ProtoMessage() {}

func (x *Rotate) ProtoReflect() protoreflect.Message {
	mi := &file_config_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rotate.ProtoReflect.Descriptor instead.
func (*Rotate) Descriptor() ([]byte, []int) {
	return file_config_log_proto_rawDescGZIP(), []int{0}
}

func (x *Rotate) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Rotate) GetMaxAge() uint32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Rotate) GetMaxSize() uint32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *Rotate) GetMaxBackups() uint32 {
	if x != nil {
		return x.MaxBackups
	}
	return 0
}

func (x *Rotate) GetLocalTime() bool {
	if x != nil {
		return x.LocalTime
	}
	return false
}

func (x *Rotate) GetCompress() bool {
	if x != nil {
		return x.Compress
	}
	return false
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level    string  `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Output   string  `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Dev      bool    `protobuf:"varint,3,opt,name=dev,proto3" json:"dev,omitempty"`
	Rotate   *Rotate `protobuf:"bytes,4,opt,name=rotate,proto3" json:"rotate,omitempty"`
	Provider string  `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"` // zap和zerolog
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_log_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Config) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Config) GetDev() bool {
	if x != nil {
		return x.Dev
	}
	return false
}

func (x *Config) GetRotate() *Rotate {
	if x != nil {
		return x.Rotate
	}
	return nil
}

func (x *Config) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

var File_config_log_proto protoreflect.FileDescriptor

var file_config_log_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xb4, 0x01, 0x0a, 0x06, 0x52,
	0x6f, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61,
	0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61,
	0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x8c, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65,
	0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x64, 0x65, 0x76, 0x12, 0x26, 0x0a, 0x06,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x73, 0x6d, 0x2d, 0x78, 0x79, 0x7a, 0x2f, 0x65, 0x7a, 0x78, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_log_proto_rawDescOnce sync.Once
	file_config_log_proto_rawDescData = file_config_log_proto_rawDesc
)

func file_config_log_proto_rawDescGZIP() []byte {
	file_config_log_proto_rawDescOnce.Do(func() {
		file_config_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_log_proto_rawDescData)
	})
	return file_config_log_proto_rawDescData
}

var file_config_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_log_proto_goTypes = []any{
	(*Rotate)(nil), // 0: config.Rotate
	(*Config)(nil), // 1: config.Config
}
var file_config_log_proto_depIdxs = []int32{
	0, // 0: config.Config.rotate:type_name -> config.Rotate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_log_proto_init() }
func file_config_log_proto_init() {
	if File_config_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_log_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Rotate); i {
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
		file_config_log_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Config); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_log_proto_goTypes,
		DependencyIndexes: file_config_log_proto_depIdxs,
		MessageInfos:      file_config_log_proto_msgTypes,
	}.Build()
	File_config_log_proto = out.File
	file_config_log_proto_rawDesc = nil
	file_config_log_proto_goTypes = nil
	file_config_log_proto_depIdxs = nil
}
