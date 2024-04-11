// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: app/tweet.proto

package app

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tweet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// usinque id of the tweet
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// unique userid of the user creating a tweet
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// content of the tweet
	Tweet string `protobuf:"bytes,3,opt,name=tweet,proto3" json:"tweet,omitempty"`
	// time when this tweet was created
	CreateTime string `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Tweet) Reset() {
	*x = Tweet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_tweet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tweet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tweet) ProtoMessage() {}

func (x *Tweet) ProtoReflect() protoreflect.Message {
	mi := &file_app_tweet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tweet.ProtoReflect.Descriptor instead.
func (*Tweet) Descriptor() ([]byte, []int) {
	return file_app_tweet_proto_rawDescGZIP(), []int{0}
}

func (x *Tweet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tweet) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Tweet) GetTweet() string {
	if x != nil {
		return x.Tweet
	}
	return ""
}

func (x *Tweet) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type CreateTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// content of the tweet
	Tweet string `protobuf:"bytes,1,opt,name=tweet,proto3" json:"tweet,omitempty"`
	// unique userid of the user creating a tweet
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *CreateTweetRequest) Reset() {
	*x = CreateTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_tweet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTweetRequest) ProtoMessage() {}

func (x *CreateTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_tweet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTweetRequest.ProtoReflect.Descriptor instead.
func (*CreateTweetRequest) Descriptor() ([]byte, []int) {
	return file_app_tweet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTweetRequest) GetTweet() string {
	if x != nil {
		return x.Tweet
	}
	return ""
}

func (x *CreateTweetRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type ListTweetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique userid of the user to list its tweets
	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *ListTweetsRequest) Reset() {
	*x = ListTweetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_tweet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTweetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTweetsRequest) ProtoMessage() {}

func (x *ListTweetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_tweet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTweetsRequest.ProtoReflect.Descriptor instead.
func (*ListTweetsRequest) Descriptor() ([]byte, []int) {
	return file_app_tweet_proto_rawDescGZIP(), []int{2}
}

func (x *ListTweetsRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type ListTweetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of tweets
	Tweets []*Tweet `protobuf:"bytes,1,rep,name=tweets,proto3" json:"tweets,omitempty"`
}

func (x *ListTweetsResponse) Reset() {
	*x = ListTweetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_tweet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTweetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTweetsResponse) ProtoMessage() {}

func (x *ListTweetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_tweet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTweetsResponse.ProtoReflect.Descriptor instead.
func (*ListTweetsResponse) Descriptor() ([]byte, []int) {
	return file_app_tweet_proto_rawDescGZIP(), []int{3}
}

func (x *ListTweetsResponse) GetTweets() []*Tweet {
	if x != nil {
		return x.Tweets
	}
	return nil
}

type DeleteTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the tweet to be deleted
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTweetRequest) Reset() {
	*x = DeleteTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_tweet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTweetRequest) ProtoMessage() {}

func (x *DeleteTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_tweet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTweetRequest.ProtoReflect.Descriptor instead.
func (*DeleteTweetRequest) Descriptor() ([]byte, []int) {
	return file_app_tweet_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTweetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_app_tweet_proto protoreflect.FileDescriptor

var file_app_tweet_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x05, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x42, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x22, 0x2b, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x3a,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x77, 0x65,
	0x65, 0x74, 0x52, 0x06, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x32, 0xd5, 0x01, 0x0a, 0x0c, 0x54, 0x77, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74,
	0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x77,
	0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x76, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0a, 0x54, 0x77, 0x65, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x32, 0x34, 0x37, 0x30, 0x34, 0x2f, 0x67, 0x6f, 0x2d, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0xa2, 0x02,
	0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x05, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_tweet_proto_rawDescOnce sync.Once
	file_app_tweet_proto_rawDescData = file_app_tweet_proto_rawDesc
)

func file_app_tweet_proto_rawDescGZIP() []byte {
	file_app_tweet_proto_rawDescOnce.Do(func() {
		file_app_tweet_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_tweet_proto_rawDescData)
	})
	return file_app_tweet_proto_rawDescData
}

var file_app_tweet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_tweet_proto_goTypes = []interface{}{
	(*Tweet)(nil),              // 0: proto.Tweet
	(*CreateTweetRequest)(nil), // 1: proto.CreateTweetRequest
	(*ListTweetsRequest)(nil),  // 2: proto.ListTweetsRequest
	(*ListTweetsResponse)(nil), // 3: proto.ListTweetsResponse
	(*DeleteTweetRequest)(nil), // 4: proto.DeleteTweetRequest
	(*emptypb.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_app_tweet_proto_depIdxs = []int32{
	0, // 0: proto.ListTweetsResponse.tweets:type_name -> proto.Tweet
	1, // 1: proto.TweetService.CreateTweet:input_type -> proto.CreateTweetRequest
	2, // 2: proto.TweetService.ListTweets:input_type -> proto.ListTweetsRequest
	4, // 3: proto.TweetService.DeleteTweet:input_type -> proto.DeleteTweetRequest
	5, // 4: proto.TweetService.CreateTweet:output_type -> google.protobuf.Empty
	3, // 5: proto.TweetService.ListTweets:output_type -> proto.ListTweetsResponse
	5, // 6: proto.TweetService.DeleteTweet:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_tweet_proto_init() }
func file_app_tweet_proto_init() {
	if File_app_tweet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_tweet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tweet); i {
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
		file_app_tweet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTweetRequest); i {
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
		file_app_tweet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTweetsRequest); i {
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
		file_app_tweet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTweetsResponse); i {
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
		file_app_tweet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTweetRequest); i {
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
			RawDescriptor: file_app_tweet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_tweet_proto_goTypes,
		DependencyIndexes: file_app_tweet_proto_depIdxs,
		MessageInfos:      file_app_tweet_proto_msgTypes,
	}.Build()
	File_app_tweet_proto = out.File
	file_app_tweet_proto_rawDesc = nil
	file_app_tweet_proto_goTypes = nil
	file_app_tweet_proto_depIdxs = nil
}
