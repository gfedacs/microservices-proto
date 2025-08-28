package shipping

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion3 

type CreateShippingRequest struct {
	OrderId              int64           `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Items                []*ShippingItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateShippingRequest) Reset()         { *m = CreateShippingRequest{} }
func (m *CreateShippingRequest) String() string { return proto.CompactTextString(m) }
func (*CreateShippingRequest) ProtoMessage()    {}
func (*CreateShippingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_218d76102d381bfb, []int{0}
}

func (m *CreateShippingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShippingRequest.Unmarshal(m, b)
}
func (m *CreateShippingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShippingRequest.Marshal(b, m, deterministic)
}
func (m *CreateShippingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShippingRequest.Merge(m, src)
}
func (m *CreateShippingRequest) XXX_Size() int {
	return xxx_messageInfo_CreateShippingRequest.Size(m)
}
func (m *CreateShippingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShippingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShippingRequest proto.InternalMessageInfo

func (m *CreateShippingRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *CreateShippingRequest) GetItems() []*ShippingItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ShippingItem struct {
	ProductCode          string   `protobuf:"bytes,1,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShippingItem) Reset()         { *m = ShippingItem{} }
func (m *ShippingItem) String() string { return proto.CompactTextString(m) }
func (*ShippingItem) ProtoMessage()    {}
func (*ShippingItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_218d76102d381bfb, []int{1}
}

func (m *ShippingItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShippingItem.Unmarshal(m, b)
}
func (m *ShippingItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShippingItem.Marshal(b, m, deterministic)
}
func (m *ShippingItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShippingItem.Merge(m, src)
}
func (m *ShippingItem) XXX_Size() int {
	return xxx_messageInfo_ShippingItem.Size(m)
}
func (m *ShippingItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ShippingItem.DiscardUnknown(m)
}

var xxx_messageInfo_ShippingItem proto.InternalMessageInfo

func (m *ShippingItem) GetProductCode() string {
	if m != nil {
		return m.ProductCode
	}
	return ""
}

func (m *ShippingItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type CreateShippingResponse struct {
	ShippingId           int64    `protobuf:"varint,1,opt,name=shipping_id,json=shippingId,proto3" json:"shipping_id,omitempty"`
	DeliveryDays         int32    `protobuf:"varint,2,opt,name=delivery_days,json=deliveryDays,proto3" json:"delivery_days,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShippingResponse) Reset()         { *m = CreateShippingResponse{} }
func (m *CreateShippingResponse) String() string { return proto.CompactTextString(m) }
func (*CreateShippingResponse) ProtoMessage()    {}
func (*CreateShippingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_218d76102d381bfb, []int{2}
}

func (m *CreateShippingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShippingResponse.Unmarshal(m, b)
}
func (m *CreateShippingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShippingResponse.Marshal(b, m, deterministic)
}
func (m *CreateShippingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShippingResponse.Merge(m, src)
}
func (m *CreateShippingResponse) XXX_Size() int {
	return xxx_messageInfo_CreateShippingResponse.Size(m)
}
func (m *CreateShippingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShippingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShippingResponse proto.InternalMessageInfo

func (m *CreateShippingResponse) GetShippingId() int64 {
	if m != nil {
		return m.ShippingId
	}
	return 0
}

func (m *CreateShippingResponse) GetDeliveryDays() int32 {
	if m != nil {
		return m.DeliveryDays
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateShippingRequest)(nil), "CreateShippingRequest")
	proto.RegisterType((*ShippingItem)(nil), "ShippingItem")
	proto.RegisterType((*CreateShippingResponse)(nil), "CreateShippingResponse")
}

func init() {
	proto.RegisterFile("shipping/shipping.proto", fileDescriptor_218d76102d381bfb)
}

var fileDescriptor_218d76102d381bfb = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xdf, 0x6e, 0x6c, 0x6f, 0x7d, 0xda, 0x5d, 0x02, 0x6e, 0x75, 0x17, 0x6b, 0xbd, 0xf4,
	0x62, 0x0b, 0x53, 0xbc, 0x78, 0x73, 0x82, 0x4c, 0xf4, 0x52, 0x0f, 0x82, 0x07, 0x4b, 0xd6, 0x3c,
	0x74, 0x81, 0xb5, 0xe9, 0xf2, 0x63, 0xd0, 0xff, 0x5e, 0xe8, 0x1a, 0x1d, 0xb2, 0x5b, 0x9e, 0x4f,
	0xc8, 0x37, 0x9f, 0x7c, 0x03, 0x33, 0xb5, 0xe1, 0x4d, 0xc3, 0xeb, 0x32, 0xb5, 0x8b, 0xa4, 0x91,
	0x42, 0x8b, 0xe8, 0x03, 0xce, 0x97, 0x12, 0xa9, 0xc6, 0xf7, 0x9e, 0x67, 0xb8, 0x33, 0xa8, 0x34,
	0xb9, 0x00, 0x57, 0x48, 0x86, 0x32, 0xe7, 0x2c, 0x70, 0x42, 0x27, 0x1e, 0x66, 0xff, 0xbb, 0x79,
	0xc5, 0xc8, 0x35, 0x8c, 0xb8, 0xc6, 0x4a, 0x05, 0x83, 0x70, 0x18, 0x7b, 0x8b, 0x49, 0x62, 0xcf,
	0xae, 0x34, 0x56, 0xd9, 0x61, 0x2f, 0x7a, 0x03, 0xff, 0x18, 0x93, 0x2b, 0xf0, 0x1b, 0x29, 0x98,
	0x29, 0x74, 0x5e, 0x08, 0x86, 0x5d, 0xe6, 0x59, 0xe6, 0xf5, 0x6c, 0x29, 0x18, 0x92, 0x39, 0xb8,
	0x3b, 0x43, 0x6b, 0xcd, 0x75, 0x1b, 0x0c, 0x42, 0x27, 0x1e, 0x65, 0x3f, 0x73, 0xf4, 0x05, 0xd3,
	0xbf, 0x9e, 0xaa, 0x11, 0xb5, 0x42, 0x72, 0x09, 0x9e, 0x7d, 0xd3, 0xaf, 0x2b, 0x58, 0xd4, 0xe9,
	0x4e, 0x18, 0x6e, 0xf9, 0x1e, 0x65, 0x9b, 0x33, 0xda, 0xaa, 0x3e, 0xdb, 0xb7, 0xf0, 0x89, 0xb6,
	0x6a, 0xf1, 0x0c, 0xae, 0x4d, 0x26, 0x0f, 0x30, 0x3e, 0xdc, 0x45, 0xa6, 0xc9, 0xc9, 0x72, 0xe6,
	0xb3, 0xe4, 0xb4, 0x4c, 0xf4, 0xef, 0xf1, 0xfe, 0xf3, 0xae, 0xe4, 0x7a, 0x63, 0xd6, 0xe9, 0xab,
	0x29, 0xa8, 0x7a, 0xa1, 0x86, 0xa5, 0x15, 0x2f, 0xa4, 0x50, 0x28, 0xf7, 0xbc, 0x40, 0x75, 0xd3,
	0x15, 0x9f, 0x96, 0x62, 0x4b, 0x8f, 0xbe, 0x63, 0x3d, 0xee, 0xf0, 0xed, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x12, 0x43, 0x7c, 0xad, 0xaa, 0x01, 0x00, 0x00,
}