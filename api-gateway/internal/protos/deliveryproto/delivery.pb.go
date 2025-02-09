// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: protos/deliveryproto/delivery.proto

package deliveryproto

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

type Delivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId string `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Status  string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Delivery) Reset() {
	*x = Delivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_deliveryproto_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delivery) ProtoMessage() {}

func (x *Delivery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_deliveryproto_delivery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delivery.ProtoReflect.Descriptor instead.
func (*Delivery) Descriptor() ([]byte, []int) {
	return file_protos_deliveryproto_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *Delivery) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Delivery) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Delivery) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Delivery) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateDeliveryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CreateDeliveryReq) Reset() {
	*x = CreateDeliveryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_deliveryproto_delivery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeliveryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeliveryReq) ProtoMessage() {}

func (x *CreateDeliveryReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_deliveryproto_delivery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeliveryReq.ProtoReflect.Descriptor instead.
func (*CreateDeliveryReq) Descriptor() ([]byte, []int) {
	return file_protos_deliveryproto_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeliveryReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateDeliveryReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateDeliveryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateDeliveryRes) Reset() {
	*x = CreateDeliveryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_deliveryproto_delivery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeliveryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeliveryRes) ProtoMessage() {}

func (x *CreateDeliveryRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_deliveryproto_delivery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeliveryRes.ProtoReflect.Descriptor instead.
func (*CreateDeliveryRes) Descriptor() ([]byte, []int) {
	return file_protos_deliveryproto_delivery_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeliveryRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateDeliveryRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetDeliveryStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryId string `protobuf:"bytes,1,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"`
}

func (x *GetDeliveryStatusReq) Reset() {
	*x = GetDeliveryStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_deliveryproto_delivery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeliveryStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeliveryStatusReq) ProtoMessage() {}

func (x *GetDeliveryStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_deliveryproto_delivery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeliveryStatusReq.ProtoReflect.Descriptor instead.
func (*GetDeliveryStatusReq) Descriptor() ([]byte, []int) {
	return file_protos_deliveryproto_delivery_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeliveryStatusReq) GetDeliveryId() string {
	if x != nil {
		return x.DeliveryId
	}
	return ""
}

type GetDeliveryStatusRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryId string `protobuf:"bytes,1,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetDeliveryStatusRes) Reset() {
	*x = GetDeliveryStatusRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_deliveryproto_delivery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeliveryStatusRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeliveryStatusRes) ProtoMessage() {}

func (x *GetDeliveryStatusRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_deliveryproto_delivery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeliveryStatusRes.ProtoReflect.Descriptor instead.
func (*GetDeliveryStatusRes) Descriptor() ([]byte, []int) {
	return file_protos_deliveryproto_delivery_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeliveryStatusRes) GetDeliveryId() string {
	if x != nil {
		return x.DeliveryId
	}
	return ""
}

func (x *GetDeliveryStatusRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateDeliveryStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryId string `protobuf:"bytes,1,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateDeliveryStatusReq) Reset() {
	*x = UpdateDeliveryStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_deliveryproto_delivery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeliveryStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeliveryStatusReq) ProtoMessage() {}

func (x *UpdateDeliveryStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_deliveryproto_delivery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeliveryStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateDeliveryStatusReq) Descriptor() ([]byte, []int) {
	return file_protos_deliveryproto_delivery_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDeliveryStatusReq) GetDeliveryId() string {
	if x != nil {
		return x.DeliveryId
	}
	return ""
}

func (x *UpdateDeliveryStatusReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateDeliveryStatusRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateDeliveryStatusRes) Reset() {
	*x = UpdateDeliveryStatusRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_deliveryproto_delivery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeliveryStatusRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeliveryStatusRes) ProtoMessage() {}

func (x *UpdateDeliveryStatusRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_deliveryproto_delivery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeliveryStatusRes.ProtoReflect.Descriptor instead.
func (*UpdateDeliveryStatusRes) Descriptor() ([]byte, []int) {
	return file_protos_deliveryproto_delivery_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDeliveryStatusRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_deliveryproto_delivery_proto protoreflect.FileDescriptor

var file_protos_deliveryproto_delivery_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x48, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x37, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x64, 0x22,
	0x4f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x52, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x33, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xce, 0x01, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12,
	0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x4a,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_deliveryproto_delivery_proto_rawDescOnce sync.Once
	file_protos_deliveryproto_delivery_proto_rawDescData = file_protos_deliveryproto_delivery_proto_rawDesc
)

func file_protos_deliveryproto_delivery_proto_rawDescGZIP() []byte {
	file_protos_deliveryproto_delivery_proto_rawDescOnce.Do(func() {
		file_protos_deliveryproto_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_deliveryproto_delivery_proto_rawDescData)
	})
	return file_protos_deliveryproto_delivery_proto_rawDescData
}

var file_protos_deliveryproto_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_deliveryproto_delivery_proto_goTypes = []any{
	(*Delivery)(nil),                // 0: Delivery
	(*CreateDeliveryReq)(nil),       // 1: CreateDeliveryReq
	(*CreateDeliveryRes)(nil),       // 2: CreateDeliveryRes
	(*GetDeliveryStatusReq)(nil),    // 3: GetDeliveryStatusReq
	(*GetDeliveryStatusRes)(nil),    // 4: GetDeliveryStatusRes
	(*UpdateDeliveryStatusReq)(nil), // 5: UpdateDeliveryStatusReq
	(*UpdateDeliveryStatusRes)(nil), // 6: UpdateDeliveryStatusRes
}
var file_protos_deliveryproto_delivery_proto_depIdxs = []int32{
	1, // 0: DeliveryService.CreateDelivery:input_type -> CreateDeliveryReq
	3, // 1: DeliveryService.GetDeliveryStatus:input_type -> GetDeliveryStatusReq
	5, // 2: DeliveryService.UpdateDeliveryStatus:input_type -> UpdateDeliveryStatusReq
	2, // 3: DeliveryService.CreateDelivery:output_type -> CreateDeliveryRes
	0, // 4: DeliveryService.GetDeliveryStatus:output_type -> Delivery
	6, // 5: DeliveryService.UpdateDeliveryStatus:output_type -> UpdateDeliveryStatusRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_deliveryproto_delivery_proto_init() }
func file_protos_deliveryproto_delivery_proto_init() {
	if File_protos_deliveryproto_delivery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_deliveryproto_delivery_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Delivery); i {
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
		file_protos_deliveryproto_delivery_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDeliveryReq); i {
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
		file_protos_deliveryproto_delivery_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDeliveryRes); i {
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
		file_protos_deliveryproto_delivery_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetDeliveryStatusReq); i {
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
		file_protos_deliveryproto_delivery_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetDeliveryStatusRes); i {
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
		file_protos_deliveryproto_delivery_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateDeliveryStatusReq); i {
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
		file_protos_deliveryproto_delivery_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateDeliveryStatusRes); i {
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
			RawDescriptor: file_protos_deliveryproto_delivery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_deliveryproto_delivery_proto_goTypes,
		DependencyIndexes: file_protos_deliveryproto_delivery_proto_depIdxs,
		MessageInfos:      file_protos_deliveryproto_delivery_proto_msgTypes,
	}.Build()
	File_protos_deliveryproto_delivery_proto = out.File
	file_protos_deliveryproto_delivery_proto_rawDesc = nil
	file_protos_deliveryproto_delivery_proto_goTypes = nil
	file_protos_deliveryproto_delivery_proto_depIdxs = nil
}
