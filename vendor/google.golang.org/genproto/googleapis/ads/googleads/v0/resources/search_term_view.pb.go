// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/search_term_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A search term view with metrics aggregated by search term at the ad group
// level.
type SearchTermView struct {
	// The resource name of the search term view.
	// Search term view resource names have the form:
	//
	// `customers/{customer_id}/searchTermViews/{campaign_id}_{ad_group_id}_
	// {URL-base64 search term}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The search term.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// The ad group the search term served in.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Indicates whether the search term is currently one of your
	// targeted or excluded keywords.
	Status               enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *SearchTermView) Reset()         { *m = SearchTermView{} }
func (m *SearchTermView) String() string { return proto.CompactTextString(m) }
func (*SearchTermView) ProtoMessage()    {}
func (*SearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_term_view_41238e08ac980d71, []int{0}
}
func (m *SearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermView.Unmarshal(m, b)
}
func (m *SearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermView.Marshal(b, m, deterministic)
}
func (dst *SearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermView.Merge(dst, src)
}
func (m *SearchTermView) XXX_Size() int {
	return xxx_messageInfo_SearchTermView.Size(m)
}
func (m *SearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermView proto.InternalMessageInfo

func (m *SearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *SearchTermView) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *SearchTermView) GetStatus() enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus {
	if m != nil {
		return m.Status
	}
	return enums.SearchTermTargetingStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SearchTermView)(nil), "google.ads.googleads.v0.resources.SearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/search_term_view.proto", fileDescriptor_search_term_view_41238e08ac980d71)
}

var fileDescriptor_search_term_view_41238e08ac980d71 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x25, 0xe9, 0x47, 0x3f, 0x9d, 0x6a, 0x17, 0x71, 0x13, 0x8a, 0x48, 0xab, 0x08, 0x5d, 0x4d,
	0x42, 0x5d, 0x28, 0x88, 0x60, 0x0a, 0x52, 0x70, 0x21, 0x25, 0x2d, 0x59, 0x48, 0x20, 0x4c, 0x9b,
	0xeb, 0x18, 0x68, 0x66, 0xc2, 0x4c, 0xa6, 0x7d, 0x1a, 0x37, 0x2e, 0x7d, 0x14, 0x9f, 0xc0, 0xc7,
	0x91, 0xfc, 0x4c, 0xb4, 0x8b, 0xaa, 0xbb, 0x93, 0xcc, 0xf9, 0x99, 0x73, 0xef, 0xa0, 0x2b, 0xca,
	0x39, 0x5d, 0x81, 0x43, 0x62, 0xe9, 0x54, 0xb0, 0x40, 0x6b, 0xd7, 0x11, 0x20, 0xb9, 0x12, 0x4b,
	0x90, 0x8e, 0x04, 0x22, 0x96, 0xcf, 0x51, 0x0e, 0x22, 0x8d, 0xd6, 0x09, 0x6c, 0x70, 0x26, 0x78,
	0xce, 0xad, 0x41, 0x45, 0xc7, 0x24, 0x96, 0xb8, 0x51, 0xe2, 0xb5, 0x8b, 0x1b, 0x65, 0xef, 0x76,
	0x97, 0x39, 0x30, 0x95, 0x6e, 0x1b, 0xe7, 0x44, 0x50, 0xc8, 0x13, 0x46, 0x23, 0x99, 0x93, 0x5c,
	0xc9, 0x2a, 0xa4, 0x77, 0x52, 0x3b, 0x94, 0x5f, 0x0b, 0xf5, 0xe4, 0x6c, 0x04, 0xc9, 0x32, 0x10,
	0xf5, 0xf9, 0xe9, 0x8b, 0x89, 0xba, 0xb3, 0xd2, 0x66, 0x0e, 0x22, 0x0d, 0x12, 0xd8, 0x58, 0x67,
	0xe8, 0x50, 0xdf, 0x20, 0x62, 0x24, 0x05, 0xdb, 0xe8, 0x1b, 0xc3, 0x7d, 0xff, 0x40, 0xff, 0x7c,
	0x20, 0x29, 0x58, 0x37, 0xa8, 0xf3, 0x2d, 0xdd, 0x36, 0xfb, 0xc6, 0xb0, 0x33, 0x3a, 0xae, 0x7b,
	0x60, 0x9d, 0x86, 0x67, 0xb9, 0x48, 0x18, 0x0d, 0xc8, 0x4a, 0x81, 0x8f, 0x64, 0x93, 0x63, 0x5d,
	0xa2, 0x3d, 0x12, 0x47, 0x54, 0x70, 0x95, 0xd9, 0xad, 0x3f, 0x68, 0xff, 0x93, 0x78, 0x52, 0x90,
	0x2d, 0x86, 0xda, 0x55, 0x3f, 0xfb, 0x5f, 0xdf, 0x18, 0x76, 0x47, 0x01, 0xde, 0x35, 0xc5, 0x72,
	0x44, 0xf8, 0xab, 0xdb, 0x5c, 0x0f, 0x68, 0x56, 0xea, 0xef, 0x98, 0x4a, 0x77, 0x9f, 0xfa, 0x75,
	0xca, 0xf8, 0xc3, 0x40, 0xe7, 0x4b, 0x9e, 0xe2, 0x5f, 0x77, 0x35, 0x3e, 0xda, 0x1e, 0xe3, 0xb4,
	0xa8, 0x31, 0x35, 0x1e, 0xef, 0x6b, 0x25, 0xe5, 0x2b, 0xc2, 0x28, 0xe6, 0x82, 0x3a, 0x14, 0x58,
	0x59, 0x52, 0xaf, 0x34, 0x4b, 0xe4, 0x0f, 0xcf, 0xe7, 0xba, 0x41, 0xaf, 0x66, 0x6b, 0xe2, 0x79,
	0x6f, 0xe6, 0x60, 0x52, 0x59, 0x7a, 0xb1, 0xc4, 0x15, 0x2c, 0x50, 0xe0, 0x62, 0x5f, 0x33, 0xdf,
	0x35, 0x27, 0xf4, 0x62, 0x19, 0x36, 0x9c, 0x30, 0x70, 0xc3, 0x86, 0xb3, 0x68, 0x97, 0x97, 0xb8,
	0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xa4, 0xe9, 0x8f, 0xc2, 0x02, 0x00, 0x00,
}
