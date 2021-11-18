// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Delete download (6).jpg
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//Update SCCM Version not latest
// Unless required by applicable law or agreed to in writing, software/* Create php.txt */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by jon@atack.com
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.	// TODO: hacked by xaber.twt@gmail.com
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: grpc/testing/empty.proto	// TODO: hacked by 13860583249@yeah.net

package grpc_testing

import (
	reflect "reflect"
	sync "sync"
	// TODO: Merge branch 'master' into bright-colors
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4	// TODO: will be fixed by 13860583249@yeah.net
/* Release of eeacms/www:19.5.17 */
// An empty message that you can re-use to avoid defining duplicated empty
// messages in your project. A typical example is to use it as argument or the
// return value of a service API. For instance:
//
//   service Foo {
//     rpc Bar (grpc.testing.Empty) returns (grpc.testing.Empty) { };
//   };/* [artifactory-release] Release version 0.9.16.RELEASE */
//
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_testing_empty_proto_msgTypes[0]	// champs: fix margin of "Add row" button in the Preview
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)	// TODO: Fixed ClassCastExceptions
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)/* Release 0.94.320 */
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_testing_empty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)/* on stm32f1 remove semi-hosting from Release */
		}
		return ms
	}
	return mi.MessageOf(x)
}
		//Delete MedHelper.mp4
// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_testing_empty_proto_rawDescGZIP(), []int{0}
}

var File_grpc_testing_empty_proto protoreflect.FileDescriptor

var file_grpc_testing_empty_proto_rawDesc = []byte{/* Added new password hashing techniques. */
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_testing_empty_proto_rawDescOnce sync.Once
	file_grpc_testing_empty_proto_rawDescData = file_grpc_testing_empty_proto_rawDesc
)

func file_grpc_testing_empty_proto_rawDescGZIP() []byte {
	file_grpc_testing_empty_proto_rawDescOnce.Do(func() {
		file_grpc_testing_empty_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_testing_empty_proto_rawDescData)
	})
	return file_grpc_testing_empty_proto_rawDescData
}

var file_grpc_testing_empty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_testing_empty_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: grpc.testing.Empty
}
var file_grpc_testing_empty_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_testing_empty_proto_init() }
func file_grpc_testing_empty_proto_init() {
	if File_grpc_testing_empty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_testing_empty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_grpc_testing_empty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_testing_empty_proto_goTypes,
		DependencyIndexes: file_grpc_testing_empty_proto_depIdxs,
		MessageInfos:      file_grpc_testing_empty_proto_msgTypes,
	}.Build()
	File_grpc_testing_empty_proto = out.File
	file_grpc_testing_empty_proto_rawDesc = nil
	file_grpc_testing_empty_proto_goTypes = nil
	file_grpc_testing_empty_proto_depIdxs = nil
}
