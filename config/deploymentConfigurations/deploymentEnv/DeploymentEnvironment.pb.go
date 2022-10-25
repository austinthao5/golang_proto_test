// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.5
// source: proto/deploymentConfigurations/deploymentEnv/DeploymentEnvironment.proto

package deploymentEnv

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
type DeploymentEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size           string        `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	Type           string        `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	AccountName    string        `protobuf:"bytes,3,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ImageVariant   string        `protobuf:"bytes,4,opt,name=imageVariant,proto3" json:"imageVariant,omitempty"`
	BootstrapOnly  string        `protobuf:"bytes,5,opt,name=bootstrapOnly,proto3" json:"bootstrapOnly,omitempty"`
	UpdateVersions string        `protobuf:"bytes,6,opt,name=updateVersions,proto3" json:"updateVersions,omitempty"`
	Consul         *Consul       `protobuf:"bytes,7,opt,name=consul,proto3" json:"consul,omitempty"`
	Vault          *Vault        `protobuf:"bytes,8,opt,name=vault,proto3" json:"vault,omitempty"`
	Location       string        `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	CustomSizing   *CustomSizing `protobuf:"bytes,10,opt,name=customSizing,proto3" json:"customSizing,omitempty"`
}

func (x *DeploymentEnvironment) Reset() {
	*x = DeploymentEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentEnvironment) ProtoMessage() {}

func (x *DeploymentEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentEnvironment.ProtoReflect.Descriptor instead.
func (*DeploymentEnvironment) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentEnvironment) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *DeploymentEnvironment) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DeploymentEnvironment) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *DeploymentEnvironment) GetImageVariant() string {
	if x != nil {
		return x.ImageVariant
	}
	return ""
}

func (x *DeploymentEnvironment) GetBootstrapOnly() string {
	if x != nil {
		return x.BootstrapOnly
	}
	return ""
}

func (x *DeploymentEnvironment) GetUpdateVersions() string {
	if x != nil {
		return x.UpdateVersions
	}
	return ""
}

func (x *DeploymentEnvironment) GetConsul() *Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *DeploymentEnvironment) GetVault() *Vault {
	if x != nil {
		return x.Vault
	}
	return nil
}

func (x *DeploymentEnvironment) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *DeploymentEnvironment) GetCustomSizing() *CustomSizing {
	if x != nil {
		return x.CustomSizing
	}
	return nil
}

type Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Enabled bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Consul) Reset() {
	*x = Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consul) ProtoMessage() {}

func (x *Consul) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consul.ProtoReflect.Descriptor instead.
func (*Consul) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP(), []int{1}
}

func (x *Consul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Consul) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type Vault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Enabled bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Vault) Reset() {
	*x = Vault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vault) ProtoMessage() {}

func (x *Vault) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vault.ProtoReflect.Descriptor instead.
func (*Vault) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP(), []int{2}
}

func (x *Vault) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Vault) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type CustomSizing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpinClouddriver *ServiceSizing `protobuf:"bytes,1,opt,name=spin_clouddriver,json=spinClouddriver,proto3" json:"spin_clouddriver,omitempty"`
	SpinVaultServer *ServiceSizing `protobuf:"bytes,2,opt,name=spin_vault_server,json=spinVaultServer,proto3" json:"spin_vault_server,omitempty"`
}

func (x *CustomSizing) Reset() {
	*x = CustomSizing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomSizing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomSizing) ProtoMessage() {}

func (x *CustomSizing) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomSizing.ProtoReflect.Descriptor instead.
func (*CustomSizing) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP(), []int{3}
}

func (x *CustomSizing) GetSpinClouddriver() *ServiceSizing {
	if x != nil {
		return x.SpinClouddriver
	}
	return nil
}

func (x *CustomSizing) GetSpinVaultServer() *ServiceSizing {
	if x != nil {
		return x.SpinVaultServer
	}
	return nil
}

type ServiceSizing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replicas int32   `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Requests *Sizing `protobuf:"bytes,2,opt,name=requests,proto3" json:"requests,omitempty"`
	Limits   *Sizing `protobuf:"bytes,3,opt,name=limits,proto3" json:"limits,omitempty"`
}

func (x *ServiceSizing) Reset() {
	*x = ServiceSizing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSizing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSizing) ProtoMessage() {}

func (x *ServiceSizing) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSizing.ProtoReflect.Descriptor instead.
func (*ServiceSizing) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceSizing) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *ServiceSizing) GetRequests() *Sizing {
	if x != nil {
		return x.Requests
	}
	return nil
}

func (x *ServiceSizing) GetLimits() *Sizing {
	if x != nil {
		return x.Limits
	}
	return nil
}

type Sizing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memory string `protobuf:"bytes,1,opt,name=memory,proto3" json:"memory,omitempty"`
	Cpu    string `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
}

func (x *Sizing) Reset() {
	*x = Sizing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sizing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sizing) ProtoMessage() {}

func (x *Sizing) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sizing.ProtoReflect.Descriptor instead.
func (*Sizing) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP(), []int{5}
}

func (x *Sizing) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *Sizing) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

var File_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDesc = []byte{
	0x0a, 0x48, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x2f, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xb5, 0x03, 0x0a, 0x15, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x4f, 0x6e, 0x6c, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3b, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4d, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x69, 0x7a, 0x69, 0x6e,
	0x67, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x22,
	0x3c, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x3b, 0x0a,
	0x05, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x0c, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x12, 0x55, 0x0a, 0x10, 0x73,
	0x70, 0x69, 0x6e, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x69, 0x6e,
	0x67, 0x52, 0x0f, 0x73, 0x70, 0x69, 0x6e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x56, 0x0a, 0x11, 0x73, 0x70, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x73, 0x70, 0x69, 0x6e, 0x56,
	0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x3f, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x52,
	0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x06, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x42, 0x2f, 0x5a, 0x2d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescData = file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDesc
)

func file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDescData
}

var file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_goTypes = []interface{}{
	(*DeploymentEnvironment)(nil), // 0: proto.deploymentEnvironment.DeploymentEnvironment
	(*Consul)(nil),                // 1: proto.deploymentEnvironment.Consul
	(*Vault)(nil),                 // 2: proto.deploymentEnvironment.Vault
	(*CustomSizing)(nil),          // 3: proto.deploymentEnvironment.CustomSizing
	(*ServiceSizing)(nil),         // 4: proto.deploymentEnvironment.ServiceSizing
	(*Sizing)(nil),                // 5: proto.deploymentEnvironment.Sizing
}
var file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_depIdxs = []int32{
	1, // 0: proto.deploymentEnvironment.DeploymentEnvironment.consul:type_name -> proto.deploymentEnvironment.Consul
	2, // 1: proto.deploymentEnvironment.DeploymentEnvironment.vault:type_name -> proto.deploymentEnvironment.Vault
	3, // 2: proto.deploymentEnvironment.DeploymentEnvironment.customSizing:type_name -> proto.deploymentEnvironment.CustomSizing
	4, // 3: proto.deploymentEnvironment.CustomSizing.spin_clouddriver:type_name -> proto.deploymentEnvironment.ServiceSizing
	4, // 4: proto.deploymentEnvironment.CustomSizing.spin_vault_server:type_name -> proto.deploymentEnvironment.ServiceSizing
	5, // 5: proto.deploymentEnvironment.ServiceSizing.requests:type_name -> proto.deploymentEnvironment.Sizing
	5, // 6: proto.deploymentEnvironment.ServiceSizing.limits:type_name -> proto.deploymentEnvironment.Sizing
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_init() }
func file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_init() {
	if File_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentEnvironment); i {
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
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consul); i {
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
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vault); i {
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
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomSizing); i {
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
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceSizing); i {
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
		file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sizing); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto = out.File
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_rawDesc = nil
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_goTypes = nil
	file_proto_deploymentConfigurations_deploymentEnv_DeploymentEnvironment_proto_depIdxs = nil
}
