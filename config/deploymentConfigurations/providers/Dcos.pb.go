// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.14.0
// source: proto/deploymentConfigurations/providers/Dcos.proto

package providers

import (
	permissions "github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/permissions"
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

// Configuration for the dcos provider.
type Dcos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled        bool            `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Accounts       []*DcosAccounts `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	PrimaryAccount string          `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	Clusters       []*DcosCluster  `protobuf:"bytes,4,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *Dcos) Reset() {
	*x = Dcos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dcos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dcos) ProtoMessage() {}

func (x *Dcos) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dcos.ProtoReflect.Descriptor instead.
func (*Dcos) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescGZIP(), []int{0}
}

func (x *Dcos) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Dcos) GetAccounts() []*DcosAccounts {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *Dcos) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

func (x *Dcos) GetClusters() []*DcosCluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type DcosAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                    string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Environment             string                   `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	RequiredGroupMembership []string                 `protobuf:"bytes,3,rep,name=requiredGroupMembership,proto3" json:"requiredGroupMembership,omitempty"`
	Permissions             *permissions.Permissions `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	DockerRegistries        []*DcosDocker            `protobuf:"bytes,5,rep,name=dockerRegistries,proto3" json:"dockerRegistries,omitempty"`
	Clusters                []*DcosAccCluster        `protobuf:"bytes,6,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *DcosAccounts) Reset() {
	*x = DcosAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosAccounts) ProtoMessage() {}

func (x *DcosAccounts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosAccounts.ProtoReflect.Descriptor instead.
func (*DcosAccounts) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescGZIP(), []int{1}
}

func (x *DcosAccounts) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DcosAccounts) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *DcosAccounts) GetRequiredGroupMembership() []string {
	if x != nil {
		return x.RequiredGroupMembership
	}
	return nil
}

func (x *DcosAccounts) GetPermissions() *permissions.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *DcosAccounts) GetDockerRegistries() []*DcosDocker {
	if x != nil {
		return x.DockerRegistries
	}
	return nil
}

func (x *DcosAccounts) GetClusters() []*DcosAccCluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type DcosDocker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string   `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	Namespaces  []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *DcosDocker) Reset() {
	*x = DcosDocker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosDocker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosDocker) ProtoMessage() {}

func (x *DcosDocker) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosDocker.ProtoReflect.Descriptor instead.
func (*DcosDocker) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescGZIP(), []int{2}
}

func (x *DcosDocker) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *DcosDocker) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type DcosAccCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid            string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Password       string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	ServiceKeyFile string `protobuf:"bytes,4,opt,name=serviceKeyFile,proto3" json:"serviceKeyFile,omitempty"`
}

func (x *DcosAccCluster) Reset() {
	*x = DcosAccCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosAccCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosAccCluster) ProtoMessage() {}

func (x *DcosAccCluster) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosAccCluster.ProtoReflect.Descriptor instead.
func (*DcosAccCluster) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescGZIP(), []int{3}
}

func (x *DcosAccCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DcosAccCluster) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *DcosAccCluster) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DcosAccCluster) GetServiceKeyFile() string {
	if x != nil {
		return x.ServiceKeyFile
	}
	return ""
}

type DcosCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DcosUrl               string  `protobuf:"bytes,2,opt,name=dcosUrl,proto3" json:"dcosUrl,omitempty"`
	CaCertFile            string  `protobuf:"bytes,3,opt,name=caCertFile,proto3" json:"caCertFile,omitempty"`
	InsecureSkipTlsVerify bool    `protobuf:"varint,4,opt,name=insecureSkipTlsVerify,proto3" json:"insecureSkipTlsVerify,omitempty"`
	LoadBalancer          *DcosLB `protobuf:"bytes,5,opt,name=loadBalancer,proto3" json:"loadBalancer,omitempty"`
}

func (x *DcosCluster) Reset() {
	*x = DcosCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosCluster) ProtoMessage() {}

func (x *DcosCluster) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosCluster.ProtoReflect.Descriptor instead.
func (*DcosCluster) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescGZIP(), []int{4}
}

func (x *DcosCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DcosCluster) GetDcosUrl() string {
	if x != nil {
		return x.DcosUrl
	}
	return ""
}

func (x *DcosCluster) GetCaCertFile() string {
	if x != nil {
		return x.CaCertFile
	}
	return ""
}

func (x *DcosCluster) GetInsecureSkipTlsVerify() bool {
	if x != nil {
		return x.InsecureSkipTlsVerify
	}
	return false
}

func (x *DcosCluster) GetLoadBalancer() *DcosLB {
	if x != nil {
		return x.LoadBalancer
	}
	return nil
}

type DcosLB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image                string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	ServiceAccountSecret string `protobuf:"bytes,2,opt,name=serviceAccountSecret,proto3" json:"serviceAccountSecret,omitempty"`
}

func (x *DcosLB) Reset() {
	*x = DcosLB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcosLB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcosLB) ProtoMessage() {}

func (x *DcosLB) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcosLB.ProtoReflect.Descriptor instead.
func (*DcosLB) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescGZIP(), []int{5}
}

func (x *DcosLB) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *DcosLB) GetServiceAccountSecret() string {
	if x != nil {
		return x.ServiceAccountSecret
	}
	return ""
}

var File_proto_deploymentConfigurations_providers_Dcos_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_providers_Dcos_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x44, 0x63, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x04, 0x44, 0x63, 0x6f, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x63, 0x6f, 0x73,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x44,
	0x63, 0x6f, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x22, 0xc6, 0x02, 0x0a, 0x0c, 0x44, 0x63, 0x6f, 0x73, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x17, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x40, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x44, 0x63, 0x6f, 0x73, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x10,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x3b, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x63, 0x6f, 0x73, 0x41, 0x63, 0x63, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x4e, 0x0a,
	0x0a, 0x44, 0x63, 0x6f, 0x73, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x7a, 0x0a,
	0x0e, 0x44, 0x63, 0x6f, 0x73, 0x41, 0x63, 0x63, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x0b, 0x44, 0x63,
	0x6f, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x63, 0x6f, 0x73, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x63, 0x6f, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x43, 0x65, 0x72,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x69, 0x6e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x54, 0x6c, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x53, 0x6b, 0x69, 0x70, 0x54, 0x6c, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x3b, 0x0a,
	0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x63, 0x6f, 0x73, 0x4c, 0x42, 0x52, 0x0c, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x06, 0x44, 0x63,
	0x6f, 0x73, 0x4c, 0x42, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x2b,
	0x5a, 0x29, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescData = file_proto_deploymentConfigurations_providers_Dcos_proto_rawDesc
)

func file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_providers_Dcos_proto_rawDescData
}

var file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_deploymentConfigurations_providers_Dcos_proto_goTypes = []interface{}{
	(*Dcos)(nil),                    // 0: proto.providers.Dcos
	(*DcosAccounts)(nil),            // 1: proto.providers.DcosAccounts
	(*DcosDocker)(nil),              // 2: proto.providers.DcosDocker
	(*DcosAccCluster)(nil),          // 3: proto.providers.DcosAccCluster
	(*DcosCluster)(nil),             // 4: proto.providers.DcosCluster
	(*DcosLB)(nil),                  // 5: proto.providers.DcosLB
	(*permissions.Permissions)(nil), // 6: proto.permissions.Permissions
}
var file_proto_deploymentConfigurations_providers_Dcos_proto_depIdxs = []int32{
	1, // 0: proto.providers.Dcos.accounts:type_name -> proto.providers.DcosAccounts
	4, // 1: proto.providers.Dcos.clusters:type_name -> proto.providers.DcosCluster
	6, // 2: proto.providers.DcosAccounts.permissions:type_name -> proto.permissions.Permissions
	2, // 3: proto.providers.DcosAccounts.dockerRegistries:type_name -> proto.providers.DcosDocker
	3, // 4: proto.providers.DcosAccounts.clusters:type_name -> proto.providers.DcosAccCluster
	5, // 5: proto.providers.DcosCluster.loadBalancer:type_name -> proto.providers.DcosLB
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_providers_Dcos_proto_init() }
func file_proto_deploymentConfigurations_providers_Dcos_proto_init() {
	if File_proto_deploymentConfigurations_providers_Dcos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dcos); i {
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
		file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosAccounts); i {
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
		file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosDocker); i {
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
		file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosAccCluster); i {
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
		file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosCluster); i {
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
		file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcosLB); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_providers_Dcos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_providers_Dcos_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_providers_Dcos_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_providers_Dcos_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_providers_Dcos_proto = out.File
	file_proto_deploymentConfigurations_providers_Dcos_proto_rawDesc = nil
	file_proto_deploymentConfigurations_providers_Dcos_proto_goTypes = nil
	file_proto_deploymentConfigurations_providers_Dcos_proto_depIdxs = nil
}