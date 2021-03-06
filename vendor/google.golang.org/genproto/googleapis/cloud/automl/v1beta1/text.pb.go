// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/text.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Dataset metadata for classification.
type TextClassificationDatasetMetadata struct {
	// Required.
	// Type of the classification problem.
	ClassificationType   ClassificationType `protobuf:"varint,1,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1beta1.ClassificationType" json:"classification_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TextClassificationDatasetMetadata) Reset()         { *m = TextClassificationDatasetMetadata{} }
func (m *TextClassificationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationDatasetMetadata) ProtoMessage()    {}
func (*TextClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_15092de8f39c6a71, []int{0}
}
func (m *TextClassificationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Unmarshal(m, b)
}
func (m *TextClassificationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Marshal(b, m, deterministic)
}
func (dst *TextClassificationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationDatasetMetadata.Merge(dst, src)
}
func (m *TextClassificationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Size(m)
}
func (m *TextClassificationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationDatasetMetadata proto.InternalMessageInfo

func (m *TextClassificationDatasetMetadata) GetClassificationType() ClassificationType {
	if m != nil {
		return m.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Model metadata that is specific to text classification.
type TextClassificationModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextClassificationModelMetadata) Reset()         { *m = TextClassificationModelMetadata{} }
func (m *TextClassificationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationModelMetadata) ProtoMessage()    {}
func (*TextClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_15092de8f39c6a71, []int{1}
}
func (m *TextClassificationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationModelMetadata.Unmarshal(m, b)
}
func (m *TextClassificationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationModelMetadata.Marshal(b, m, deterministic)
}
func (dst *TextClassificationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationModelMetadata.Merge(dst, src)
}
func (m *TextClassificationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationModelMetadata.Size(m)
}
func (m *TextClassificationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationModelMetadata proto.InternalMessageInfo

// Dataset metadata that is specific to text extraction
type TextExtractionDatasetMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextExtractionDatasetMetadata) Reset()         { *m = TextExtractionDatasetMetadata{} }
func (m *TextExtractionDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextExtractionDatasetMetadata) ProtoMessage()    {}
func (*TextExtractionDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_15092de8f39c6a71, []int{2}
}
func (m *TextExtractionDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Unmarshal(m, b)
}
func (m *TextExtractionDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Marshal(b, m, deterministic)
}
func (dst *TextExtractionDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextExtractionDatasetMetadata.Merge(dst, src)
}
func (m *TextExtractionDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextExtractionDatasetMetadata.Size(m)
}
func (m *TextExtractionDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextExtractionDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextExtractionDatasetMetadata proto.InternalMessageInfo

// Model metadata that is specific to text extraction.
type TextExtractionModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextExtractionModelMetadata) Reset()         { *m = TextExtractionModelMetadata{} }
func (m *TextExtractionModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextExtractionModelMetadata) ProtoMessage()    {}
func (*TextExtractionModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_15092de8f39c6a71, []int{3}
}
func (m *TextExtractionModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextExtractionModelMetadata.Unmarshal(m, b)
}
func (m *TextExtractionModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextExtractionModelMetadata.Marshal(b, m, deterministic)
}
func (dst *TextExtractionModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextExtractionModelMetadata.Merge(dst, src)
}
func (m *TextExtractionModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextExtractionModelMetadata.Size(m)
}
func (m *TextExtractionModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextExtractionModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextExtractionModelMetadata proto.InternalMessageInfo

// Dataset metadata for text sentiment.
type TextSentimentDatasetMetadata struct {
	// Required.
	// A sentiment is expressed as an integer ordinal, where higher value
	// means a more positive sentiment. The range of sentiments that will be used
	// is between 0 and sentiment_max (inclusive on both ends), and all the values
	// in the range must be represented in the dataset before a model can be
	// created.
	// sentiment_max value must be between 1 and 10 (inclusive).
	SentimentMax         int32    `protobuf:"varint,1,opt,name=sentiment_max,json=sentimentMax,proto3" json:"sentiment_max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentDatasetMetadata) Reset()         { *m = TextSentimentDatasetMetadata{} }
func (m *TextSentimentDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextSentimentDatasetMetadata) ProtoMessage()    {}
func (*TextSentimentDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_15092de8f39c6a71, []int{4}
}
func (m *TextSentimentDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Unmarshal(m, b)
}
func (m *TextSentimentDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Marshal(b, m, deterministic)
}
func (dst *TextSentimentDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentDatasetMetadata.Merge(dst, src)
}
func (m *TextSentimentDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextSentimentDatasetMetadata.Size(m)
}
func (m *TextSentimentDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentDatasetMetadata proto.InternalMessageInfo

func (m *TextSentimentDatasetMetadata) GetSentimentMax() int32 {
	if m != nil {
		return m.SentimentMax
	}
	return 0
}

// Model metadata that is specific to text classification.
type TextSentimentModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentModelMetadata) Reset()         { *m = TextSentimentModelMetadata{} }
func (m *TextSentimentModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextSentimentModelMetadata) ProtoMessage()    {}
func (*TextSentimentModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_15092de8f39c6a71, []int{5}
}
func (m *TextSentimentModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentModelMetadata.Unmarshal(m, b)
}
func (m *TextSentimentModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentModelMetadata.Marshal(b, m, deterministic)
}
func (dst *TextSentimentModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentModelMetadata.Merge(dst, src)
}
func (m *TextSentimentModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextSentimentModelMetadata.Size(m)
}
func (m *TextSentimentModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentModelMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TextClassificationDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TextClassificationDatasetMetadata")
	proto.RegisterType((*TextClassificationModelMetadata)(nil), "google.cloud.automl.v1beta1.TextClassificationModelMetadata")
	proto.RegisterType((*TextExtractionDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TextExtractionDatasetMetadata")
	proto.RegisterType((*TextExtractionModelMetadata)(nil), "google.cloud.automl.v1beta1.TextExtractionModelMetadata")
	proto.RegisterType((*TextSentimentDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TextSentimentDatasetMetadata")
	proto.RegisterType((*TextSentimentModelMetadata)(nil), "google.cloud.automl.v1beta1.TextSentimentModelMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/text.proto", fileDescriptor_text_15092de8f39c6a71)
}

var fileDescriptor_text_15092de8f39c6a71 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x59, 0x41, 0xc1, 0xa0, 0x1e, 0xea, 0x45, 0xfa, 0x87, 0xda, 0x15, 0xc4, 0x53, 0x62,
	0xf5, 0xe8, 0xa9, 0xad, 0xe2, 0x69, 0xa1, 0xd4, 0xe2, 0x41, 0x0a, 0x75, 0xba, 0x1d, 0x97, 0x85,
	0x6c, 0x66, 0xe9, 0x4e, 0x65, 0xfb, 0x01, 0x3c, 0xfb, 0xbd, 0xfc, 0x54, 0xb2, 0x49, 0x2b, 0xc4,
	0x96, 0x1e, 0x93, 0xf7, 0x7b, 0x2f, 0x2f, 0x33, 0xe2, 0x3a, 0x21, 0x4a, 0x34, 0xaa, 0x58, 0xd3,
	0x72, 0xae, 0x60, 0xc9, 0x94, 0x69, 0xf5, 0xd9, 0x9d, 0x21, 0x43, 0x57, 0x31, 0x96, 0x2c, 0xf3,
	0x05, 0x31, 0xd5, 0x1a, 0x8e, 0x93, 0x96, 0x93, 0x8e, 0x93, 0x6b, 0xae, 0xde, 0x5c, 0x87, 0x40,
	0x9e, 0x2a, 0x30, 0x86, 0x18, 0x38, 0x25, 0x53, 0x38, 0x6b, 0xfd, 0x76, 0xdf, 0x13, 0xb1, 0x86,
	0xa2, 0x48, 0x3f, 0xd2, 0xd8, 0x5a, 0x9c, 0x23, 0xfc, 0x0a, 0x44, 0x67, 0x8c, 0x25, 0x0f, 0x3c,
	0xf1, 0x11, 0x18, 0x0a, 0xe4, 0x08, 0x19, 0xe6, 0xc0, 0x50, 0x7b, 0x17, 0xe7, 0xbe, 0x7b, 0xca,
	0xab, 0x1c, 0x2f, 0x82, 0xcb, 0xe0, 0xe6, 0xec, 0x4e, 0xc9, 0x3d, 0x85, 0xa5, 0x1f, 0x3c, 0x5e,
	0xe5, 0x38, 0xaa, 0xc5, 0x5b, 0x77, 0x61, 0x47, 0xb4, 0xb7, 0x6b, 0x44, 0x34, 0x47, 0xbd, 0x29,
	0x11, 0xb6, 0x45, 0xab, 0x42, 0x9e, 0x4a, 0x5e, 0x40, 0xbc, 0xa3, 0x65, 0xd8, 0x12, 0x0d, 0x1f,
	0xf0, 0xfd, 0x03, 0xd1, 0xac, 0xe4, 0x17, 0x34, 0x9c, 0x66, 0x68, 0xf8, 0xff, 0x27, 0xaf, 0xc4,
	0x69, 0xb1, 0xd1, 0xa6, 0x19, 0x94, 0xf6, 0x7b, 0x87, 0xa3, 0x93, 0xbf, 0xcb, 0x08, 0xca, 0xb0,
	0x29, 0xea, 0x5e, 0x88, 0xf7, 0x44, 0xff, 0x3b, 0x10, 0xed, 0x98, 0xb2, 0x7d, 0x03, 0xe9, 0x1f,
	0x57, 0xfe, 0x61, 0x35, 0xfc, 0x61, 0xf0, 0xd6, 0x5b, 0x93, 0x09, 0x69, 0x30, 0x89, 0xa4, 0x45,
	0xa2, 0x12, 0x34, 0x76, 0x35, 0xca, 0x49, 0x90, 0xa7, 0xc5, 0xce, 0x7d, 0x3e, 0xb8, 0xe3, 0xcf,
	0x41, 0xe3, 0xd9, 0x82, 0x93, 0x41, 0x05, 0x4d, 0x7a, 0x4b, 0xa6, 0x48, 0x4f, 0x5e, 0x1d, 0x34,
	0x3b, 0xb2, 0x59, 0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x91, 0x55, 0x3e, 0x63, 0x7d, 0x02,
	0x00, 0x00,
}
