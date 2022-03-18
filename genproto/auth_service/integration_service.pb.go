// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.2
// source: integration_service.proto

package auth_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
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

type CreateIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId        string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId string `protobuf:"bytes,2,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId     string `protobuf:"bytes,3,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	RoleId           string `protobuf:"bytes,4,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	SecretKey        string `protobuf:"bytes,5,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Active           int32  `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	ExpiresAt        string `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Title            string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	IpWhitelist      string `protobuf:"bytes,9,opt,name=ip_whitelist,json=ipWhitelist,proto3" json:"ip_whitelist,omitempty"`
}

func (x *CreateIntegrationRequest) Reset() {
	*x = CreateIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIntegrationRequest) ProtoMessage() {}

func (x *CreateIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIntegrationRequest.ProtoReflect.Descriptor instead.
func (*CreateIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_integration_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateIntegrationRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateIntegrationRequest) GetClientPlatformId() string {
	if x != nil {
		return x.ClientPlatformId
	}
	return ""
}

func (x *CreateIntegrationRequest) GetClientTypeId() string {
	if x != nil {
		return x.ClientTypeId
	}
	return ""
}

func (x *CreateIntegrationRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CreateIntegrationRequest) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *CreateIntegrationRequest) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *CreateIntegrationRequest) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

func (x *CreateIntegrationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateIntegrationRequest) GetIpWhitelist() string {
	if x != nil {
		return x.IpWhitelist
	}
	return ""
}

type IntegrationPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IntegrationPrimaryKey) Reset() {
	*x = IntegrationPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationPrimaryKey) ProtoMessage() {}

func (x *IntegrationPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_integration_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationPrimaryKey.ProtoReflect.Descriptor instead.
func (*IntegrationPrimaryKey) Descriptor() ([]byte, []int) {
	return file_integration_service_proto_rawDescGZIP(), []int{1}
}

func (x *IntegrationPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IntegrationPrimaryKeyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IntegrationPrimaryKeyList) Reset() {
	*x = IntegrationPrimaryKeyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationPrimaryKeyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationPrimaryKeyList) ProtoMessage() {}

func (x *IntegrationPrimaryKeyList) ProtoReflect() protoreflect.Message {
	mi := &file_integration_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationPrimaryKeyList.ProtoReflect.Descriptor instead.
func (*IntegrationPrimaryKeyList) Descriptor() ([]byte, []int) {
	return file_integration_service_proto_rawDescGZIP(), []int{2}
}

func (x *IntegrationPrimaryKeyList) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetIntegrationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit            int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset           int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search           string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	ClientPlatformId string `protobuf:"bytes,4,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId     string `protobuf:"bytes,5,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	ProjectId        string `protobuf:"bytes,6,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetIntegrationListRequest) Reset() {
	*x = GetIntegrationListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIntegrationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIntegrationListRequest) ProtoMessage() {}

func (x *GetIntegrationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIntegrationListRequest.ProtoReflect.Descriptor instead.
func (*GetIntegrationListRequest) Descriptor() ([]byte, []int) {
	return file_integration_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetIntegrationListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetIntegrationListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetIntegrationListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetIntegrationListRequest) GetClientPlatformId() string {
	if x != nil {
		return x.ClientPlatformId
	}
	return ""
}

func (x *GetIntegrationListRequest) GetClientTypeId() string {
	if x != nil {
		return x.ClientTypeId
	}
	return ""
}

func (x *GetIntegrationListRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetIntegrationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count        int32          `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Integrations []*Integration `protobuf:"bytes,2,rep,name=integrations,proto3" json:"integrations,omitempty"`
}

func (x *GetIntegrationListResponse) Reset() {
	*x = GetIntegrationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIntegrationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIntegrationListResponse) ProtoMessage() {}

func (x *GetIntegrationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integration_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIntegrationListResponse.ProtoReflect.Descriptor instead.
func (*GetIntegrationListResponse) Descriptor() ([]byte, []int) {
	return file_integration_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetIntegrationListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetIntegrationListResponse) GetIntegrations() []*Integration {
	if x != nil {
		return x.Integrations
	}
	return nil
}

type UpdateIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId        string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientPlatformId string `protobuf:"bytes,3,opt,name=client_platform_id,json=clientPlatformId,proto3" json:"client_platform_id,omitempty"`
	ClientTypeId     string `protobuf:"bytes,4,opt,name=client_type_id,json=clientTypeId,proto3" json:"client_type_id,omitempty"`
	RoleId           string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Active           int32  `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	ExpiresAt        string `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Name             string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	IpWhitelist      string `protobuf:"bytes,9,opt,name=ip_whitelist,json=ipWhitelist,proto3" json:"ip_whitelist,omitempty"`
}

func (x *UpdateIntegrationRequest) Reset() {
	*x = UpdateIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIntegrationRequest) ProtoMessage() {}

func (x *UpdateIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIntegrationRequest.ProtoReflect.Descriptor instead.
func (*UpdateIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_integration_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateIntegrationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateIntegrationRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *UpdateIntegrationRequest) GetClientPlatformId() string {
	if x != nil {
		return x.ClientPlatformId
	}
	return ""
}

func (x *UpdateIntegrationRequest) GetClientTypeId() string {
	if x != nil {
		return x.ClientTypeId
	}
	return ""
}

func (x *UpdateIntegrationRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UpdateIntegrationRequest) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *UpdateIntegrationRequest) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

func (x *UpdateIntegrationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateIntegrationRequest) GetIpWhitelist() string {
	if x != nil {
		return x.IpWhitelist
	}
	return ""
}

type AddIntegrationRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	RelationId    string `protobuf:"bytes,2,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
}

func (x *AddIntegrationRelationRequest) Reset() {
	*x = AddIntegrationRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddIntegrationRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddIntegrationRelationRequest) ProtoMessage() {}

func (x *AddIntegrationRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddIntegrationRelationRequest.ProtoReflect.Descriptor instead.
func (*AddIntegrationRelationRequest) Descriptor() ([]byte, []int) {
	return file_integration_service_proto_rawDescGZIP(), []int{6}
}

func (x *AddIntegrationRelationRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *AddIntegrationRelationRequest) GetRelationId() string {
	if x != nil {
		return x.RelationId
	}
	return ""
}

type IntegrationRelationPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	RelationId    string `protobuf:"bytes,2,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
}

func (x *IntegrationRelationPrimaryKey) Reset() {
	*x = IntegrationRelationPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationRelationPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationRelationPrimaryKey) ProtoMessage() {}

func (x *IntegrationRelationPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_integration_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationRelationPrimaryKey.ProtoReflect.Descriptor instead.
func (*IntegrationRelationPrimaryKey) Descriptor() ([]byte, []int) {
	return file_integration_service_proto_rawDescGZIP(), []int{7}
}

func (x *IntegrationRelationPrimaryKey) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *IntegrationRelationPrimaryKey) GetRelationId() string {
	if x != nil {
		return x.RelationId
	}
	return ""
}

var File_integration_service_proto protoreflect.FileDescriptor

var file_integration_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb5, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x70, 0x5f, 0x77, 0x68, 0x69, 0x74,
	0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x57,
	0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2d, 0x0a, 0x19, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0xd4, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa4, 0x02, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x70, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x67, 0x0a, 0x1d, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x1d, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x32, 0xcf, 0x04, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a,
	0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integration_service_proto_rawDescOnce sync.Once
	file_integration_service_proto_rawDescData = file_integration_service_proto_rawDesc
)

func file_integration_service_proto_rawDescGZIP() []byte {
	file_integration_service_proto_rawDescOnce.Do(func() {
		file_integration_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_integration_service_proto_rawDescData)
	})
	return file_integration_service_proto_rawDescData
}

var file_integration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_integration_service_proto_goTypes = []interface{}{
	(*CreateIntegrationRequest)(nil),      // 0: auth_service.CreateIntegrationRequest
	(*IntegrationPrimaryKey)(nil),         // 1: auth_service.IntegrationPrimaryKey
	(*IntegrationPrimaryKeyList)(nil),     // 2: auth_service.IntegrationPrimaryKeyList
	(*GetIntegrationListRequest)(nil),     // 3: auth_service.GetIntegrationListRequest
	(*GetIntegrationListResponse)(nil),    // 4: auth_service.GetIntegrationListResponse
	(*UpdateIntegrationRequest)(nil),      // 5: auth_service.UpdateIntegrationRequest
	(*AddIntegrationRelationRequest)(nil), // 6: auth_service.AddIntegrationRelationRequest
	(*IntegrationRelationPrimaryKey)(nil), // 7: auth_service.IntegrationRelationPrimaryKey
	(*Integration)(nil),                   // 8: auth_service.Integration
	(*empty.Empty)(nil),                   // 9: google.protobuf.Empty
}
var file_integration_service_proto_depIdxs = []int32{
	8, // 0: auth_service.GetIntegrationListResponse.integrations:type_name -> auth_service.Integration
	0, // 1: auth_service.IntegrationService.CreateIntegration:input_type -> auth_service.CreateIntegrationRequest
	1, // 2: auth_service.IntegrationService.GetIntegrationByID:input_type -> auth_service.IntegrationPrimaryKey
	2, // 3: auth_service.IntegrationService.GetIntegrationListByIDs:input_type -> auth_service.IntegrationPrimaryKeyList
	3, // 4: auth_service.IntegrationService.GetIntegrationList:input_type -> auth_service.GetIntegrationListRequest
	5, // 5: auth_service.IntegrationService.UpdateIntegration:input_type -> auth_service.UpdateIntegrationRequest
	1, // 6: auth_service.IntegrationService.DeleteIntegration:input_type -> auth_service.IntegrationPrimaryKey
	8, // 7: auth_service.IntegrationService.CreateIntegration:output_type -> auth_service.Integration
	8, // 8: auth_service.IntegrationService.GetIntegrationByID:output_type -> auth_service.Integration
	4, // 9: auth_service.IntegrationService.GetIntegrationListByIDs:output_type -> auth_service.GetIntegrationListResponse
	4, // 10: auth_service.IntegrationService.GetIntegrationList:output_type -> auth_service.GetIntegrationListResponse
	8, // 11: auth_service.IntegrationService.UpdateIntegration:output_type -> auth_service.Integration
	9, // 12: auth_service.IntegrationService.DeleteIntegration:output_type -> google.protobuf.Empty
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_integration_service_proto_init() }
func file_integration_service_proto_init() {
	if File_integration_service_proto != nil {
		return
	}
	file_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_integration_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIntegrationRequest); i {
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
		file_integration_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationPrimaryKey); i {
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
		file_integration_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationPrimaryKeyList); i {
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
		file_integration_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIntegrationListRequest); i {
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
		file_integration_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIntegrationListResponse); i {
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
		file_integration_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIntegrationRequest); i {
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
		file_integration_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddIntegrationRelationRequest); i {
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
		file_integration_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationRelationPrimaryKey); i {
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
			RawDescriptor: file_integration_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integration_service_proto_goTypes,
		DependencyIndexes: file_integration_service_proto_depIdxs,
		MessageInfos:      file_integration_service_proto_msgTypes,
	}.Build()
	File_integration_service_proto = out.File
	file_integration_service_proto_rawDesc = nil
	file_integration_service_proto_goTypes = nil
	file_integration_service_proto_depIdxs = nil
}
