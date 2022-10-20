// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.5
// source: proto/deploymentConfigurations/providers/Kubernetes.proto

package providers

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

// Configuration for the clouddriver microservice.
type Kubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled        bool             `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	PrimaryAccount string           `protobuf:"bytes,2,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	Account        []*KubernetesAcc `protobuf:"bytes,3,rep,name=Account,proto3" json:"Account,omitempty"`
}

func (x *Kubernetes) Reset() {
	*x = Kubernetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Kubernetes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kubernetes) ProtoMessage() {}

func (x *Kubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Kubernetes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kubernetes.ProtoReflect.Descriptor instead.
func (*Kubernetes) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescGZIP(), []int{0}
}

func (x *Kubernetes) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Kubernetes) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

func (x *Kubernetes) GetAccount() []*KubernetesAcc {
	if x != nil {
		return x.Account
	}
	return nil
}

type KubernetesAcc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RequiredGroupMembership   []string `protobuf:"bytes,2,rep,name=requiredGroupMembership,proto3" json:"requiredGroupMembership,omitempty"`
	Permissions               []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	DockerRegistries          []string `protobuf:"bytes,4,rep,name=dockerRegistries,proto3" json:"dockerRegistries,omitempty"`
	ProviderVersion           string   `protobuf:"bytes,5,opt,name=providerVersion,proto3" json:"providerVersion,omitempty"`
	ConfigureImagePullSecrets bool     `protobuf:"varint,6,opt,name=configureImagePullSecrets,proto3" json:"configureImagePullSecrets,omitempty"`
	ServiceAccount            bool     `protobuf:"varint,7,opt,name=serviceAccount,proto3" json:"serviceAccount,omitempty"`
	CacheThreads              int32    `protobuf:"varint,8,opt,name=cacheThreads,proto3" json:"cacheThreads,omitempty"`
	Namespaces                []string `protobuf:"bytes,9,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	KubeconfigFile            string   `protobuf:"bytes,10,opt,name=kubeconfigFile,proto3" json:"kubeconfigFile,omitempty"`
	OmitNamespaces            []string `protobuf:"bytes,11,rep,name=omitNamespaces,proto3" json:"omitNamespaces,omitempty"`
	Kinds                     []string `protobuf:"bytes,12,rep,name=kinds,proto3" json:"kinds,omitempty"`
	OmitKinds                 []string `protobuf:"bytes,13,rep,name=omitKinds,proto3" json:"omitKinds,omitempty"`
	CustomResources           []string `protobuf:"bytes,14,rep,name=customResources,proto3" json:"customResources,omitempty"`
	CachingPolicies           []string `protobuf:"bytes,15,rep,name=cachingPolicies,proto3" json:"cachingPolicies,omitempty"`
	OauthScopes               []string `protobuf:"bytes,16,rep,name=oauthScopes,proto3" json:"oauthScopes,omitempty"`
	OnlySpinnakerManaged      bool     `protobuf:"varint,17,opt,name=onlySpinnakerManaged,proto3" json:"onlySpinnakerManaged,omitempty"`
}

func (x *KubernetesAcc) Reset() {
	*x = KubernetesAcc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Kubernetes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesAcc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesAcc) ProtoMessage() {}

func (x *KubernetesAcc) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Kubernetes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesAcc.ProtoReflect.Descriptor instead.
func (*KubernetesAcc) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesAcc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubernetesAcc) GetRequiredGroupMembership() []string {
	if x != nil {
		return x.RequiredGroupMembership
	}
	return nil
}

func (x *KubernetesAcc) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *KubernetesAcc) GetDockerRegistries() []string {
	if x != nil {
		return x.DockerRegistries
	}
	return nil
}

func (x *KubernetesAcc) GetProviderVersion() string {
	if x != nil {
		return x.ProviderVersion
	}
	return ""
}

func (x *KubernetesAcc) GetConfigureImagePullSecrets() bool {
	if x != nil {
		return x.ConfigureImagePullSecrets
	}
	return false
}

func (x *KubernetesAcc) GetServiceAccount() bool {
	if x != nil {
		return x.ServiceAccount
	}
	return false
}

func (x *KubernetesAcc) GetCacheThreads() int32 {
	if x != nil {
		return x.CacheThreads
	}
	return 0
}

func (x *KubernetesAcc) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *KubernetesAcc) GetKubeconfigFile() string {
	if x != nil {
		return x.KubeconfigFile
	}
	return ""
}

func (x *KubernetesAcc) GetOmitNamespaces() []string {
	if x != nil {
		return x.OmitNamespaces
	}
	return nil
}

func (x *KubernetesAcc) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

func (x *KubernetesAcc) GetOmitKinds() []string {
	if x != nil {
		return x.OmitKinds
	}
	return nil
}

func (x *KubernetesAcc) GetCustomResources() []string {
	if x != nil {
		return x.CustomResources
	}
	return nil
}

func (x *KubernetesAcc) GetCachingPolicies() []string {
	if x != nil {
		return x.CachingPolicies
	}
	return nil
}

func (x *KubernetesAcc) GetOauthScopes() []string {
	if x != nil {
		return x.OauthScopes
	}
	return nil
}

func (x *KubernetesAcc) GetOnlySpinnakerManaged() bool {
	if x != nil {
		return x.OnlySpinnakerManaged
	}
	return false
}

var File_proto_deploymentConfigurations_providers_Kubernetes_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDesc = []byte{
	0x0a, 0x39, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0x88, 0x01, 0x0a,
	0x0a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x63, 0x63, 0x52, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xad, 0x05, 0x0a, 0x0d, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x63, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a,
	0x17, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x3c, 0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x75, 0x62,
	0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x6d, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x6d, 0x69, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x69, 0x6e,
	0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x6f, 0x6d, 0x69, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x6d, 0x69, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0f, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x6f, 0x6e, 0x6c, 0x79, 0x53, 0x70, 0x69, 0x6e, 0x6e,
	0x61, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x14, 0x6f, 0x6e, 0x6c, 0x79, 0x53, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x42, 0x2b, 0x5a, 0x29, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescData = file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDesc
)

func file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDescData
}

var file_proto_deploymentConfigurations_providers_Kubernetes_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_deploymentConfigurations_providers_Kubernetes_proto_goTypes = []interface{}{
	(*Kubernetes)(nil),    // 0: proto.providers.Kubernetes
	(*KubernetesAcc)(nil), // 1: proto.providers.KubernetesAcc
}
var file_proto_deploymentConfigurations_providers_Kubernetes_proto_depIdxs = []int32{
	1, // 0: proto.providers.Kubernetes.Account:type_name -> proto.providers.KubernetesAcc
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_providers_Kubernetes_proto_init() }
func file_proto_deploymentConfigurations_providers_Kubernetes_proto_init() {
	if File_proto_deploymentConfigurations_providers_Kubernetes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_providers_Kubernetes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kubernetes); i {
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
		file_proto_deploymentConfigurations_providers_Kubernetes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesAcc); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_providers_Kubernetes_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_providers_Kubernetes_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_providers_Kubernetes_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_providers_Kubernetes_proto = out.File
	file_proto_deploymentConfigurations_providers_Kubernetes_proto_rawDesc = nil
	file_proto_deploymentConfigurations_providers_Kubernetes_proto_goTypes = nil
	file_proto_deploymentConfigurations_providers_Kubernetes_proto_depIdxs = nil
}
