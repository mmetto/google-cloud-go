// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/devtools/artifactregistry/v1/package.proto

package artifactregistrypb

import (
	reflect "reflect"
	sync "sync"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Packages are named collections of versions.
type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the package, for example:
	// "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1".
	// If the package ID part contains slashes, the slashes are escaped.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The display name of the package.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The time when the package was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time when the package was last updated. This includes publishing a new
	// version of the package.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_package_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Package) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Package) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// The request to list packages.
type ListPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the parent resource whose packages will be listed.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of packages to return. Maximum page size is 1,000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous list request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListPackagesRequest) Reset() {
	*x = ListPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPackagesRequest) ProtoMessage() {}

func (x *ListPackagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPackagesRequest.ProtoReflect.Descriptor instead.
func (*ListPackagesRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_package_proto_rawDescGZIP(), []int{1}
}

func (x *ListPackagesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListPackagesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPackagesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response from listing packages.
type ListPackagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The packages returned.
	Packages []*Package `protobuf:"bytes,1,rep,name=packages,proto3" json:"packages,omitempty"`
	// The token to retrieve the next page of packages, or empty if there are no
	// more packages to return.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPackagesResponse) Reset() {
	*x = ListPackagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPackagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPackagesResponse) ProtoMessage() {}

func (x *ListPackagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPackagesResponse.ProtoReflect.Descriptor instead.
func (*ListPackagesResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_package_proto_rawDescGZIP(), []int{2}
}

func (x *ListPackagesResponse) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *ListPackagesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The request to retrieve a package.
type GetPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the package to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPackageRequest) Reset() {
	*x = GetPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageRequest) ProtoMessage() {}

func (x *GetPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageRequest.ProtoReflect.Descriptor instead.
func (*GetPackageRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_package_proto_rawDescGZIP(), []int{3}
}

func (x *GetPackageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request to delete a package.
type DeletePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the package to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeletePackageRequest) Reset() {
	*x = DeletePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackageRequest) ProtoMessage() {}

func (x *DeletePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackageRequest.ProtoReflect.Descriptor instead.
func (*DeletePackageRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1_package_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePackageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_devtools_artifactregistry_v1_package_proto protoreflect.FileDescriptor

var file_google_devtools_artifactregistry_v1_package_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x07, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2f, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0xfa, 0x01, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_artifactregistry_v1_package_proto_rawDescOnce sync.Once
	file_google_devtools_artifactregistry_v1_package_proto_rawDescData = file_google_devtools_artifactregistry_v1_package_proto_rawDesc
)

func file_google_devtools_artifactregistry_v1_package_proto_rawDescGZIP() []byte {
	file_google_devtools_artifactregistry_v1_package_proto_rawDescOnce.Do(func() {
		file_google_devtools_artifactregistry_v1_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_artifactregistry_v1_package_proto_rawDescData)
	})
	return file_google_devtools_artifactregistry_v1_package_proto_rawDescData
}

var file_google_devtools_artifactregistry_v1_package_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_devtools_artifactregistry_v1_package_proto_goTypes = []interface{}{
	(*Package)(nil),              // 0: google.devtools.artifactregistry.v1.Package
	(*ListPackagesRequest)(nil),  // 1: google.devtools.artifactregistry.v1.ListPackagesRequest
	(*ListPackagesResponse)(nil), // 2: google.devtools.artifactregistry.v1.ListPackagesResponse
	(*GetPackageRequest)(nil),    // 3: google.devtools.artifactregistry.v1.GetPackageRequest
	(*DeletePackageRequest)(nil), // 4: google.devtools.artifactregistry.v1.DeletePackageRequest
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_google_devtools_artifactregistry_v1_package_proto_depIdxs = []int32{
	5, // 0: google.devtools.artifactregistry.v1.Package.create_time:type_name -> google.protobuf.Timestamp
	5, // 1: google.devtools.artifactregistry.v1.Package.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.devtools.artifactregistry.v1.ListPackagesResponse.packages:type_name -> google.devtools.artifactregistry.v1.Package
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_devtools_artifactregistry_v1_package_proto_init() }
func file_google_devtools_artifactregistry_v1_package_proto_init() {
	if File_google_devtools_artifactregistry_v1_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_artifactregistry_v1_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_google_devtools_artifactregistry_v1_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPackagesRequest); i {
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
		file_google_devtools_artifactregistry_v1_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPackagesResponse); i {
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
		file_google_devtools_artifactregistry_v1_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageRequest); i {
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
		file_google_devtools_artifactregistry_v1_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackageRequest); i {
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
			RawDescriptor: file_google_devtools_artifactregistry_v1_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_artifactregistry_v1_package_proto_goTypes,
		DependencyIndexes: file_google_devtools_artifactregistry_v1_package_proto_depIdxs,
		MessageInfos:      file_google_devtools_artifactregistry_v1_package_proto_msgTypes,
	}.Build()
	File_google_devtools_artifactregistry_v1_package_proto = out.File
	file_google_devtools_artifactregistry_v1_package_proto_rawDesc = nil
	file_google_devtools_artifactregistry_v1_package_proto_goTypes = nil
	file_google_devtools_artifactregistry_v1_package_proto_depIdxs = nil
}
