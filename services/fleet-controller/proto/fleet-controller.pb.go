// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleet-controller.proto

package fleet_controller

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type BookingRequest struct {
	VehicleId            string   `protobuf:"bytes,1,opt,name=VehicleId,proto3" json:"VehicleId,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingRequest) Reset()         { *m = BookingRequest{} }
func (m *BookingRequest) String() string { return proto.CompactTextString(m) }
func (*BookingRequest) ProtoMessage()    {}
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{1}
}
func (m *BookingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingRequest.Unmarshal(m, b)
}
func (m *BookingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingRequest.Marshal(b, m, deterministic)
}
func (dst *BookingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingRequest.Merge(dst, src)
}
func (m *BookingRequest) XXX_Size() int {
	return xxx_messageInfo_BookingRequest.Size(m)
}
func (m *BookingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BookingRequest proto.InternalMessageInfo

func (m *BookingRequest) GetVehicleId() string {
	if m != nil {
		return m.VehicleId
	}
	return ""
}

func (m *BookingRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type BookingResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	ReservedTime         uint32   `protobuf:"varint,2,opt,name=ReservedTime,proto3" json:"ReservedTime,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingResponse) Reset()         { *m = BookingResponse{} }
func (m *BookingResponse) String() string { return proto.CompactTextString(m) }
func (*BookingResponse) ProtoMessage()    {}
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{2}
}
func (m *BookingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingResponse.Unmarshal(m, b)
}
func (m *BookingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingResponse.Marshal(b, m, deterministic)
}
func (dst *BookingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingResponse.Merge(dst, src)
}
func (m *BookingResponse) XXX_Size() int {
	return xxx_messageInfo_BookingResponse.Size(m)
}
func (m *BookingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BookingResponse proto.InternalMessageInfo

func (m *BookingResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *BookingResponse) GetReservedTime() uint32 {
	if m != nil {
		return m.ReservedTime
	}
	return 0
}

func (m *BookingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UnbookingRequest struct {
	VehicleId            string   `protobuf:"bytes,1,opt,name=VehicleId,proto3" json:"VehicleId,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnbookingRequest) Reset()         { *m = UnbookingRequest{} }
func (m *UnbookingRequest) String() string { return proto.CompactTextString(m) }
func (*UnbookingRequest) ProtoMessage()    {}
func (*UnbookingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{3}
}
func (m *UnbookingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnbookingRequest.Unmarshal(m, b)
}
func (m *UnbookingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnbookingRequest.Marshal(b, m, deterministic)
}
func (dst *UnbookingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbookingRequest.Merge(dst, src)
}
func (m *UnbookingRequest) XXX_Size() int {
	return xxx_messageInfo_UnbookingRequest.Size(m)
}
func (m *UnbookingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbookingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnbookingRequest proto.InternalMessageInfo

func (m *UnbookingRequest) GetVehicleId() string {
	if m != nil {
		return m.VehicleId
	}
	return ""
}

func (m *UnbookingRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type UnbookingResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnbookingResponse) Reset()         { *m = UnbookingResponse{} }
func (m *UnbookingResponse) String() string { return proto.CompactTextString(m) }
func (*UnbookingResponse) ProtoMessage()    {}
func (*UnbookingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{4}
}
func (m *UnbookingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnbookingResponse.Unmarshal(m, b)
}
func (m *UnbookingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnbookingResponse.Marshal(b, m, deterministic)
}
func (dst *UnbookingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbookingResponse.Merge(dst, src)
}
func (m *UnbookingResponse) XXX_Size() int {
	return xxx_messageInfo_UnbookingResponse.Size(m)
}
func (m *UnbookingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbookingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnbookingResponse proto.InternalMessageInfo

func (m *UnbookingResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *UnbookingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type BeginRideRequest struct {
	CustomerId           string   `protobuf:"bytes,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeginRideRequest) Reset()         { *m = BeginRideRequest{} }
func (m *BeginRideRequest) String() string { return proto.CompactTextString(m) }
func (*BeginRideRequest) ProtoMessage()    {}
func (*BeginRideRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{5}
}
func (m *BeginRideRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeginRideRequest.Unmarshal(m, b)
}
func (m *BeginRideRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeginRideRequest.Marshal(b, m, deterministic)
}
func (dst *BeginRideRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeginRideRequest.Merge(dst, src)
}
func (m *BeginRideRequest) XXX_Size() int {
	return xxx_messageInfo_BeginRideRequest.Size(m)
}
func (m *BeginRideRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BeginRideRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BeginRideRequest proto.InternalMessageInfo

func (m *BeginRideRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type BeginRideResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeginRideResponse) Reset()         { *m = BeginRideResponse{} }
func (m *BeginRideResponse) String() string { return proto.CompactTextString(m) }
func (*BeginRideResponse) ProtoMessage()    {}
func (*BeginRideResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{6}
}
func (m *BeginRideResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeginRideResponse.Unmarshal(m, b)
}
func (m *BeginRideResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeginRideResponse.Marshal(b, m, deterministic)
}
func (dst *BeginRideResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeginRideResponse.Merge(dst, src)
}
func (m *BeginRideResponse) XXX_Size() int {
	return xxx_messageInfo_BeginRideResponse.Size(m)
}
func (m *BeginRideResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BeginRideResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BeginRideResponse proto.InternalMessageInfo

func (m *BeginRideResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *BeginRideResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EndRideRequest struct {
	CustomerId           string   `protobuf:"bytes,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndRideRequest) Reset()         { *m = EndRideRequest{} }
func (m *EndRideRequest) String() string { return proto.CompactTextString(m) }
func (*EndRideRequest) ProtoMessage()    {}
func (*EndRideRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{7}
}
func (m *EndRideRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndRideRequest.Unmarshal(m, b)
}
func (m *EndRideRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndRideRequest.Marshal(b, m, deterministic)
}
func (dst *EndRideRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndRideRequest.Merge(dst, src)
}
func (m *EndRideRequest) XXX_Size() int {
	return xxx_messageInfo_EndRideRequest.Size(m)
}
func (m *EndRideRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndRideRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndRideRequest proto.InternalMessageInfo

func (m *EndRideRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type EndRideResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndRideResponse) Reset()         { *m = EndRideResponse{} }
func (m *EndRideResponse) String() string { return proto.CompactTextString(m) }
func (*EndRideResponse) ProtoMessage()    {}
func (*EndRideResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{8}
}
func (m *EndRideResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndRideResponse.Unmarshal(m, b)
}
func (m *EndRideResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndRideResponse.Marshal(b, m, deterministic)
}
func (dst *EndRideResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndRideResponse.Merge(dst, src)
}
func (m *EndRideResponse) XXX_Size() int {
	return xxx_messageInfo_EndRideResponse.Size(m)
}
func (m *EndRideResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndRideResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndRideResponse proto.InternalMessageInfo

func (m *EndRideResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *EndRideResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RetrieveReservationsResponse struct {
	Reservations         []*RetrieveReservationsResponse_Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
	Error                string                                      `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *RetrieveReservationsResponse) Reset()         { *m = RetrieveReservationsResponse{} }
func (m *RetrieveReservationsResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveReservationsResponse) ProtoMessage()    {}
func (*RetrieveReservationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{9}
}
func (m *RetrieveReservationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveReservationsResponse.Unmarshal(m, b)
}
func (m *RetrieveReservationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveReservationsResponse.Marshal(b, m, deterministic)
}
func (dst *RetrieveReservationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveReservationsResponse.Merge(dst, src)
}
func (m *RetrieveReservationsResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveReservationsResponse.Size(m)
}
func (m *RetrieveReservationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveReservationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveReservationsResponse proto.InternalMessageInfo

func (m *RetrieveReservationsResponse) GetReservations() []*RetrieveReservationsResponse_Reservation {
	if m != nil {
		return m.Reservations
	}
	return nil
}

func (m *RetrieveReservationsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RetrieveReservationsResponse_Reservation struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Vehicle              string   `protobuf:"bytes,3,opt,name=Vehicle,proto3" json:"Vehicle,omitempty"`
	Customer             string   `protobuf:"bytes,4,opt,name=Customer,proto3" json:"Customer,omitempty"`
	Status               uint32   `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveReservationsResponse_Reservation) Reset() {
	*m = RetrieveReservationsResponse_Reservation{}
}
func (m *RetrieveReservationsResponse_Reservation) String() string { return proto.CompactTextString(m) }
func (*RetrieveReservationsResponse_Reservation) ProtoMessage()    {}
func (*RetrieveReservationsResponse_Reservation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{9, 0}
}
func (m *RetrieveReservationsResponse_Reservation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveReservationsResponse_Reservation.Unmarshal(m, b)
}
func (m *RetrieveReservationsResponse_Reservation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveReservationsResponse_Reservation.Marshal(b, m, deterministic)
}
func (dst *RetrieveReservationsResponse_Reservation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveReservationsResponse_Reservation.Merge(dst, src)
}
func (m *RetrieveReservationsResponse_Reservation) XXX_Size() int {
	return xxx_messageInfo_RetrieveReservationsResponse_Reservation.Size(m)
}
func (m *RetrieveReservationsResponse_Reservation) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveReservationsResponse_Reservation.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveReservationsResponse_Reservation proto.InternalMessageInfo

func (m *RetrieveReservationsResponse_Reservation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RetrieveReservationsResponse_Reservation) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *RetrieveReservationsResponse_Reservation) GetVehicle() string {
	if m != nil {
		return m.Vehicle
	}
	return ""
}

func (m *RetrieveReservationsResponse_Reservation) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *RetrieveReservationsResponse_Reservation) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type RetrieveUnbilledBookingsResponse struct {
	Bookings             []*Booking `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RetrieveUnbilledBookingsResponse) Reset()         { *m = RetrieveUnbilledBookingsResponse{} }
func (m *RetrieveUnbilledBookingsResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveUnbilledBookingsResponse) ProtoMessage()    {}
func (*RetrieveUnbilledBookingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{10}
}
func (m *RetrieveUnbilledBookingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveUnbilledBookingsResponse.Unmarshal(m, b)
}
func (m *RetrieveUnbilledBookingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveUnbilledBookingsResponse.Marshal(b, m, deterministic)
}
func (dst *RetrieveUnbilledBookingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveUnbilledBookingsResponse.Merge(dst, src)
}
func (m *RetrieveUnbilledBookingsResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveUnbilledBookingsResponse.Size(m)
}
func (m *RetrieveUnbilledBookingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveUnbilledBookingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveUnbilledBookingsResponse proto.InternalMessageInfo

func (m *RetrieveUnbilledBookingsResponse) GetBookings() []*Booking {
	if m != nil {
		return m.Bookings
	}
	return nil
}

type RetrieveBilledBookingsByCustomerRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveBilledBookingsByCustomerRequest) Reset() {
	*m = RetrieveBilledBookingsByCustomerRequest{}
}
func (m *RetrieveBilledBookingsByCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveBilledBookingsByCustomerRequest) ProtoMessage()    {}
func (*RetrieveBilledBookingsByCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{11}
}
func (m *RetrieveBilledBookingsByCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveBilledBookingsByCustomerRequest.Unmarshal(m, b)
}
func (m *RetrieveBilledBookingsByCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveBilledBookingsByCustomerRequest.Marshal(b, m, deterministic)
}
func (dst *RetrieveBilledBookingsByCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveBilledBookingsByCustomerRequest.Merge(dst, src)
}
func (m *RetrieveBilledBookingsByCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveBilledBookingsByCustomerRequest.Size(m)
}
func (m *RetrieveBilledBookingsByCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveBilledBookingsByCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveBilledBookingsByCustomerRequest proto.InternalMessageInfo

func (m *RetrieveBilledBookingsByCustomerRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type RetrieveBilledBookingsByCustomerResponse struct {
	Bookings             []*Booking `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RetrieveBilledBookingsByCustomerResponse) Reset() {
	*m = RetrieveBilledBookingsByCustomerResponse{}
}
func (m *RetrieveBilledBookingsByCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveBilledBookingsByCustomerResponse) ProtoMessage()    {}
func (*RetrieveBilledBookingsByCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{12}
}
func (m *RetrieveBilledBookingsByCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveBilledBookingsByCustomerResponse.Unmarshal(m, b)
}
func (m *RetrieveBilledBookingsByCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveBilledBookingsByCustomerResponse.Marshal(b, m, deterministic)
}
func (dst *RetrieveBilledBookingsByCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveBilledBookingsByCustomerResponse.Merge(dst, src)
}
func (m *RetrieveBilledBookingsByCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveBilledBookingsByCustomerResponse.Size(m)
}
func (m *RetrieveBilledBookingsByCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveBilledBookingsByCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveBilledBookingsByCustomerResponse proto.InternalMessageInfo

func (m *RetrieveBilledBookingsByCustomerResponse) GetBookings() []*Booking {
	if m != nil {
		return m.Bookings
	}
	return nil
}

type AddInvoiceToBookingRequest struct {
	BookingId            uint32   `protobuf:"varint,1,opt,name=BookingId,proto3" json:"BookingId,omitempty"`
	InvoiceId            string   `protobuf:"bytes,2,opt,name=InvoiceId,proto3" json:"InvoiceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddInvoiceToBookingRequest) Reset()         { *m = AddInvoiceToBookingRequest{} }
func (m *AddInvoiceToBookingRequest) String() string { return proto.CompactTextString(m) }
func (*AddInvoiceToBookingRequest) ProtoMessage()    {}
func (*AddInvoiceToBookingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{13}
}
func (m *AddInvoiceToBookingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddInvoiceToBookingRequest.Unmarshal(m, b)
}
func (m *AddInvoiceToBookingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddInvoiceToBookingRequest.Marshal(b, m, deterministic)
}
func (dst *AddInvoiceToBookingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddInvoiceToBookingRequest.Merge(dst, src)
}
func (m *AddInvoiceToBookingRequest) XXX_Size() int {
	return xxx_messageInfo_AddInvoiceToBookingRequest.Size(m)
}
func (m *AddInvoiceToBookingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddInvoiceToBookingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddInvoiceToBookingRequest proto.InternalMessageInfo

func (m *AddInvoiceToBookingRequest) GetBookingId() uint32 {
	if m != nil {
		return m.BookingId
	}
	return 0
}

func (m *AddInvoiceToBookingRequest) GetInvoiceId() string {
	if m != nil {
		return m.InvoiceId
	}
	return ""
}

type AddInvoiceToBookingResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddInvoiceToBookingResponse) Reset()         { *m = AddInvoiceToBookingResponse{} }
func (m *AddInvoiceToBookingResponse) String() string { return proto.CompactTextString(m) }
func (*AddInvoiceToBookingResponse) ProtoMessage()    {}
func (*AddInvoiceToBookingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{14}
}
func (m *AddInvoiceToBookingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddInvoiceToBookingResponse.Unmarshal(m, b)
}
func (m *AddInvoiceToBookingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddInvoiceToBookingResponse.Marshal(b, m, deterministic)
}
func (dst *AddInvoiceToBookingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddInvoiceToBookingResponse.Merge(dst, src)
}
func (m *AddInvoiceToBookingResponse) XXX_Size() int {
	return xxx_messageInfo_AddInvoiceToBookingResponse.Size(m)
}
func (m *AddInvoiceToBookingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddInvoiceToBookingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddInvoiceToBookingResponse proto.InternalMessageInfo

func (m *AddInvoiceToBookingResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

type Booking struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Customer             string   `protobuf:"bytes,3,opt,name=Customer,proto3" json:"Customer,omitempty"`
	DistanceDriven       uint32   `protobuf:"varint,4,opt,name=DistanceDriven,proto3" json:"DistanceDriven,omitempty"`
	TimeDriven           uint32   `protobuf:"varint,5,opt,name=TimeDriven,proto3" json:"TimeDriven,omitempty"`
	Invoice              string   `protobuf:"bytes,6,opt,name=Invoice,proto3" json:"Invoice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Booking) Reset()         { *m = Booking{} }
func (m *Booking) String() string { return proto.CompactTextString(m) }
func (*Booking) ProtoMessage()    {}
func (*Booking) Descriptor() ([]byte, []int) {
	return fileDescriptor_fleet_controller_7315f0aa947faac5, []int{15}
}
func (m *Booking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Booking.Unmarshal(m, b)
}
func (m *Booking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Booking.Marshal(b, m, deterministic)
}
func (dst *Booking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Booking.Merge(dst, src)
}
func (m *Booking) XXX_Size() int {
	return xxx_messageInfo_Booking.Size(m)
}
func (m *Booking) XXX_DiscardUnknown() {
	xxx_messageInfo_Booking.DiscardUnknown(m)
}

var xxx_messageInfo_Booking proto.InternalMessageInfo

func (m *Booking) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Booking) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Booking) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *Booking) GetDistanceDriven() uint32 {
	if m != nil {
		return m.DistanceDriven
	}
	return 0
}

func (m *Booking) GetTimeDriven() uint32 {
	if m != nil {
		return m.TimeDriven
	}
	return 0
}

func (m *Booking) GetInvoice() string {
	if m != nil {
		return m.Invoice
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*BookingRequest)(nil), "BookingRequest")
	proto.RegisterType((*BookingResponse)(nil), "BookingResponse")
	proto.RegisterType((*UnbookingRequest)(nil), "UnbookingRequest")
	proto.RegisterType((*UnbookingResponse)(nil), "UnbookingResponse")
	proto.RegisterType((*BeginRideRequest)(nil), "BeginRideRequest")
	proto.RegisterType((*BeginRideResponse)(nil), "BeginRideResponse")
	proto.RegisterType((*EndRideRequest)(nil), "EndRideRequest")
	proto.RegisterType((*EndRideResponse)(nil), "EndRideResponse")
	proto.RegisterType((*RetrieveReservationsResponse)(nil), "RetrieveReservationsResponse")
	proto.RegisterType((*RetrieveReservationsResponse_Reservation)(nil), "RetrieveReservationsResponse.Reservation")
	proto.RegisterType((*RetrieveUnbilledBookingsResponse)(nil), "RetrieveUnbilledBookingsResponse")
	proto.RegisterType((*RetrieveBilledBookingsByCustomerRequest)(nil), "RetrieveBilledBookingsByCustomerRequest")
	proto.RegisterType((*RetrieveBilledBookingsByCustomerResponse)(nil), "RetrieveBilledBookingsByCustomerResponse")
	proto.RegisterType((*AddInvoiceToBookingRequest)(nil), "AddInvoiceToBookingRequest")
	proto.RegisterType((*AddInvoiceToBookingResponse)(nil), "AddInvoiceToBookingResponse")
	proto.RegisterType((*Booking)(nil), "Booking")
}

func init() {
	proto.RegisterFile("fleet-controller.proto", fileDescriptor_fleet_controller_7315f0aa947faac5)
}

var fileDescriptor_fleet_controller_7315f0aa947faac5 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x75, 0x12, 0x48, 0xc2, 0x05, 0x92, 0x70, 0x3f, 0x84, 0x2c, 0xc3, 0x57, 0xd1, 0x51, 0xd5,
	0x06, 0x55, 0x1d, 0xb5, 0xb4, 0xdb, 0xaa, 0x22, 0x40, 0x69, 0x16, 0xad, 0xd0, 0x00, 0x55, 0xb7,
	0x21, 0xbe, 0x50, 0x0b, 0xe3, 0xa1, 0xe3, 0x49, 0x24, 0x1e, 0xa1, 0xab, 0xbe, 0x4a, 0x37, 0x7d,
	0xbf, 0xca, 0xf6, 0xf8, 0x27, 0x26, 0x40, 0x10, 0x5d, 0xce, 0x19, 0xdf, 0x9f, 0xb9, 0xf7, 0x9c,
	0x63, 0x58, 0x3b, 0xf3, 0x89, 0xf4, 0xab, 0xa1, 0x0c, 0xb4, 0x92, 0xbe, 0x4f, 0x8a, 0x5f, 0x29,
	0xa9, 0x25, 0x6b, 0xc0, 0xfc, 0xfe, 0xe5, 0x95, 0xbe, 0x66, 0x5f, 0xa0, 0xd5, 0x93, 0xf2, 0xc2,
	0x0b, 0xce, 0x05, 0xfd, 0x18, 0x51, 0xa8, 0x71, 0x03, 0x16, 0xbe, 0xd2, 0x77, 0x6f, 0xe8, 0x53,
	0xdf, 0xb5, 0x2b, 0x9b, 0x95, 0xee, 0x82, 0xc8, 0x01, 0x7c, 0x02, 0xb0, 0x3b, 0x0a, 0xb5, 0xbc,
	0x24, 0xd5, 0x77, 0xed, 0x6a, 0x7c, 0x5d, 0x40, 0xd8, 0x05, 0xb4, 0xb3, 0x7c, 0xe1, 0x95, 0x0c,
	0x42, 0x8a, 0x42, 0x8e, 0x46, 0xc3, 0x21, 0x85, 0xe1, 0xd9, 0xc8, 0x8f, 0x33, 0x36, 0x45, 0x01,
	0x41, 0x06, 0x4b, 0x82, 0x42, 0x52, 0x63, 0x72, 0x8f, 0xbd, 0x4b, 0x8a, 0x93, 0x2e, 0x8b, 0x09,
	0x0c, 0x57, 0x61, 0x7e, 0x5f, 0x29, 0xa9, 0xec, 0x5a, 0x5c, 0x31, 0x39, 0xb0, 0x43, 0xe8, 0x9c,
	0x04, 0xa7, 0xff, 0xb2, 0xfd, 0x3e, 0xac, 0x14, 0x32, 0xce, 0xf8, 0x80, 0xac, 0xb9, 0x6a, 0xb1,
	0xb9, 0x6d, 0xe8, 0xf4, 0xe8, 0xdc, 0x0b, 0x84, 0xe7, 0x52, 0xda, 0xdc, 0x0c, 0xe5, 0x0b, 0x31,
	0x8f, 0x2a, 0xff, 0x1a, 0x5a, 0xfb, 0x81, 0xfb, 0x90, 0xe2, 0x07, 0xd0, 0xce, 0x22, 0x1e, 0x55,
	0xfa, 0x57, 0x15, 0x36, 0x04, 0x69, 0xe5, 0xd1, 0x98, 0x92, 0x2d, 0x0e, 0xb4, 0x27, 0x83, 0x30,
	0x4b, 0xfb, 0x19, 0x96, 0x54, 0x01, 0xb7, 0x2b, 0x9b, 0xb5, 0xee, 0xe2, 0xf6, 0x16, 0xbf, 0x2b,
	0x88, 0x17, 0x40, 0x31, 0x11, 0x3e, 0xbd, 0x0b, 0xe7, 0x67, 0x05, 0x16, 0x0b, 0x31, 0xd8, 0x82,
	0xaa, 0x61, 0xc4, 0xb2, 0xa8, 0xf6, 0xdd, 0x88, 0x28, 0xbb, 0x8a, 0x06, 0x9a, 0xdc, 0x1d, 0x1d,
	0x47, 0xd6, 0x44, 0x0e, 0xa0, 0x0d, 0x0d, 0xc3, 0x1a, 0x43, 0xb9, 0xf4, 0x88, 0x0e, 0x34, 0xd3,
	0xa1, 0xd9, 0x73, 0xf1, 0x55, 0x76, 0xc6, 0x35, 0xa8, 0x1f, 0xe9, 0x81, 0x1e, 0x85, 0xf6, 0x7c,
	0x5c, 0xc7, 0x9c, 0xd8, 0x27, 0xd8, 0x4c, 0xdf, 0x76, 0x12, 0x9c, 0x7a, 0xbe, 0x4f, 0xae, 0x51,
	0x49, 0x3e, 0x94, 0x67, 0xd0, 0x34, 0xc4, 0x4b, 0x07, 0xd2, 0xe4, 0xa9, 0x94, 0xb2, 0x1b, 0xb6,
	0x03, 0x2f, 0xd2, 0x4c, 0xbd, 0x89, 0x3c, 0xbd, 0xeb, 0xb4, 0x8b, 0x74, 0xdf, 0x6b, 0x50, 0x3f,
	0x09, 0xe3, 0x5d, 0x27, 0x32, 0x30, 0x27, 0x76, 0x08, 0xdd, 0xfb, 0x53, 0x3c, 0xa8, 0xa9, 0x6f,
	0xe0, 0xec, 0xb8, 0x6e, 0x3f, 0x18, 0x4b, 0x6f, 0x48, 0xc7, 0xf2, 0xa6, 0xa1, 0x18, 0x24, 0x9b,
	0x7f, 0x0e, 0x44, 0xb7, 0x26, 0x30, 0x23, 0x65, 0x0e, 0xb0, 0xf7, 0xb0, 0x3e, 0x35, 0xf3, 0x6c,
	0xfc, 0x64, 0x7f, 0x2a, 0xd0, 0x30, 0x31, 0x0f, 0xdc, 0x7f, 0x71, 0xcb, 0xb5, 0xd2, 0x96, 0x9f,
	0x43, 0x6b, 0xcf, 0x0b, 0xf5, 0x20, 0x18, 0xd2, 0x9e, 0xf2, 0xc6, 0x14, 0xc4, 0x3c, 0x58, 0x16,
	0x25, 0x34, 0xea, 0x2e, 0x32, 0x2f, 0xf3, 0x4d, 0xc2, 0x88, 0x02, 0x12, 0x71, 0xcc, 0xbc, 0xcc,
	0xae, 0x27, 0x1c, 0x33, 0xc7, 0xed, 0xdf, 0x73, 0xd0, 0xfe, 0x18, 0x39, 0xf7, 0x6e, 0x66, 0xdc,
	0xf8, 0x12, 0xe6, 0xa2, 0xa7, 0x60, 0x9b, 0x4f, 0xce, 0xd7, 0xe9, 0xf0, 0xd2, 0x58, 0x98, 0x85,
	0x6f, 0xa0, 0x9e, 0xf8, 0x18, 0xae, 0xf0, 0xb2, 0x45, 0x3a, 0xc8, 0x6f, 0x78, 0x1c, 0xb3, 0xf0,
	0x1d, 0x2c, 0x64, 0xde, 0x83, 0x2b, 0xbc, 0xec, 0x5d, 0x0e, 0xf2, 0x1b, 0xd6, 0xc4, 0x2c, 0xe4,
	0xd0, 0x30, 0xa6, 0x81, 0x6d, 0x3e, 0x69, 0x38, 0x4e, 0x87, 0x97, 0xfc, 0x84, 0x59, 0xf8, 0x01,
	0x56, 0xa7, 0xa9, 0x1c, 0xeb, 0x3c, 0xfe, 0x1f, 0x39, 0xff, 0xdf, 0x69, 0x02, 0xcc, 0xc2, 0x03,
	0xb0, 0x6f, 0x93, 0x52, 0x96, 0xe4, 0x29, 0xbf, 0x4f, 0x6d, 0xcc, 0xc2, 0xeb, 0x5c, 0x93, 0xb7,
	0xc9, 0x00, 0xbb, 0x7c, 0x46, 0xb1, 0x39, 0x5b, 0x7c, 0x56, 0x4d, 0x31, 0x0b, 0x05, 0xfc, 0x37,
	0x85, 0xd5, 0xb8, 0xce, 0x6f, 0x57, 0x91, 0xb3, 0xc1, 0xef, 0x10, 0x02, 0xb3, 0x4e, 0xeb, 0xf1,
	0x8f, 0xfd, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x90, 0x13, 0x40, 0xf2, 0x07, 0x00,
	0x00,
}
