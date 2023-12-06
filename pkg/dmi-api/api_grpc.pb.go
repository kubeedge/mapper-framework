//
//Copyright 2023 The KubeEdge Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// To regenerate api.pb.go run hack/generate-dmi.sh

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: api.proto

package v1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DeviceManagerService_MapperRegister_FullMethodName     = "/v1beta1.DeviceManagerService/MapperRegister"
	DeviceManagerService_ReportDeviceStatus_FullMethodName = "/v1beta1.DeviceManagerService/ReportDeviceStatus"
)

// DeviceManagerServiceClient is the client API for DeviceManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceManagerServiceClient interface {
	// MapperRegister registers the information of the mapper to device manager
	// when the mapper is online. Device manager returns the list of devices and device models which
	// this mapper should manage.
	MapperRegister(ctx context.Context, in *MapperRegisterRequest, opts ...grpc.CallOption) (*MapperRegisterResponse, error)
	// ReportDeviceStatus reports the status of devices to device manager.
	// When the mapper collects some properties of a device, it can make them a map of device twins
	// and report it to the device manager through the interface of ReportDeviceStatus.
	ReportDeviceStatus(ctx context.Context, in *ReportDeviceStatusRequest, opts ...grpc.CallOption) (*ReportDeviceStatusResponse, error)
}

type deviceManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceManagerServiceClient(cc grpc.ClientConnInterface) DeviceManagerServiceClient {
	return &deviceManagerServiceClient{cc}
}

func (c *deviceManagerServiceClient) MapperRegister(ctx context.Context, in *MapperRegisterRequest, opts ...grpc.CallOption) (*MapperRegisterResponse, error) {
	out := new(MapperRegisterResponse)
	err := c.cc.Invoke(ctx, DeviceManagerService_MapperRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManagerServiceClient) ReportDeviceStatus(ctx context.Context, in *ReportDeviceStatusRequest, opts ...grpc.CallOption) (*ReportDeviceStatusResponse, error) {
	out := new(ReportDeviceStatusResponse)
	err := c.cc.Invoke(ctx, DeviceManagerService_ReportDeviceStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceManagerServiceServer is the server API for DeviceManagerService service.
// All implementations must embed UnimplementedDeviceManagerServiceServer
// for forward compatibility
type DeviceManagerServiceServer interface {
	// MapperRegister registers the information of the mapper to device manager
	// when the mapper is online. Device manager returns the list of devices and device models which
	// this mapper should manage.
	MapperRegister(context.Context, *MapperRegisterRequest) (*MapperRegisterResponse, error)
	// ReportDeviceStatus reports the status of devices to device manager.
	// When the mapper collects some properties of a device, it can make them a map of device twins
	// and report it to the device manager through the interface of ReportDeviceStatus.
	ReportDeviceStatus(context.Context, *ReportDeviceStatusRequest) (*ReportDeviceStatusResponse, error)
	mustEmbedUnimplementedDeviceManagerServiceServer()
}

// UnimplementedDeviceManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceManagerServiceServer struct {
}

func (UnimplementedDeviceManagerServiceServer) MapperRegister(context.Context, *MapperRegisterRequest) (*MapperRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapperRegister not implemented")
}
func (UnimplementedDeviceManagerServiceServer) ReportDeviceStatus(context.Context, *ReportDeviceStatusRequest) (*ReportDeviceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportDeviceStatus not implemented")
}
func (UnimplementedDeviceManagerServiceServer) mustEmbedUnimplementedDeviceManagerServiceServer() {}

// UnsafeDeviceManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceManagerServiceServer will
// result in compilation errors.
type UnsafeDeviceManagerServiceServer interface {
	mustEmbedUnimplementedDeviceManagerServiceServer()
}

func RegisterDeviceManagerServiceServer(s grpc.ServiceRegistrar, srv DeviceManagerServiceServer) {
	s.RegisterService(&DeviceManagerService_ServiceDesc, srv)
}

func _DeviceManagerService_MapperRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapperRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServiceServer).MapperRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceManagerService_MapperRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServiceServer).MapperRegister(ctx, req.(*MapperRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManagerService_ReportDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportDeviceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServiceServer).ReportDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceManagerService_ReportDeviceStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServiceServer).ReportDeviceStatus(ctx, req.(*ReportDeviceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceManagerService_ServiceDesc is the grpc.ServiceDesc for DeviceManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1beta1.DeviceManagerService",
	HandlerType: (*DeviceManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MapperRegister",
			Handler:    _DeviceManagerService_MapperRegister_Handler,
		},
		{
			MethodName: "ReportDeviceStatus",
			Handler:    _DeviceManagerService_ReportDeviceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	DeviceMapperService_RegisterDevice_FullMethodName    = "/v1beta1.DeviceMapperService/RegisterDevice"
	DeviceMapperService_RemoveDevice_FullMethodName      = "/v1beta1.DeviceMapperService/RemoveDevice"
	DeviceMapperService_UpdateDevice_FullMethodName      = "/v1beta1.DeviceMapperService/UpdateDevice"
	DeviceMapperService_CreateDeviceModel_FullMethodName = "/v1beta1.DeviceMapperService/CreateDeviceModel"
	DeviceMapperService_RemoveDeviceModel_FullMethodName = "/v1beta1.DeviceMapperService/RemoveDeviceModel"
	DeviceMapperService_UpdateDeviceModel_FullMethodName = "/v1beta1.DeviceMapperService/UpdateDeviceModel"
	DeviceMapperService_GetDevice_FullMethodName         = "/v1beta1.DeviceMapperService/GetDevice"
)

// DeviceMapperServiceClient is the client API for DeviceMapperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceMapperServiceClient interface {
	// RegisterDevice registers a device to the device mapper.
	// Device manager registers a device instance with the information of device
	// to the mapper through the interface of RegisterDevice.
	// When the mapper gets the request of register with device information,
	// it should add the device to the device list and connect to the real physical device via the specific protocol.
	RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*RegisterDeviceResponse, error)
	// RemoveDevice unregisters a device to the device mapper.
	// Device manager unregisters a device instance with the name of device
	// to the mapper through the interface of RemoveDevice.
	// When the mapper gets the request of unregister with device name,
	// it should remove the device from the device list and disconnect to the real physical device.
	RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error)
	// UpdateDevice updates a device to the device mapper
	// Device manager updates the information of a device used by the mapper
	// through the interface of UpdateDevice.
	// The information of a device includes the meta data and the status data of a device.
	// When the mapper gets the request of updating with the information of a device,
	// it should update the device of the device list and connect to the real physical device via the updated information.
	UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error)
	// CreateDeviceModel creates a device model to the device mapper.
	// Device manager sends the information of device model to the mapper
	// through the interface of CreateDeviceModel.
	// When the mapper gets the request of creating with the information of device model,
	// it should create a new device model to the list of device models.
	CreateDeviceModel(ctx context.Context, in *CreateDeviceModelRequest, opts ...grpc.CallOption) (*CreateDeviceModelResponse, error)
	// RemoveDeviceModel remove a device model to the device mapper.
	// Device manager sends the name of device model to the mapper
	// through the interface of RemoveDeviceModel.
	// When the mapper gets the request of removing with the name of device model,
	// it should remove the device model to the list of device models.
	RemoveDeviceModel(ctx context.Context, in *RemoveDeviceModelRequest, opts ...grpc.CallOption) (*RemoveDeviceModelResponse, error)
	// UpdateDeviceModel update a device model to the device mapper.
	// Device manager sends the information of device model to the mapper
	// through the interface of UpdateDeviceModel.
	// When the mapper gets the request of updating with the information of device model,
	// it should update the device model to the list of device models.
	UpdateDeviceModel(ctx context.Context, in *UpdateDeviceModelRequest, opts ...grpc.CallOption) (*UpdateDeviceModelResponse, error)
	// GetDevice get the information of a device from the device mapper.
	// Device sends the request of querying device information with the device name to the mapper
	// through the interface of GetDevice.
	// When the mapper gets the request of querying with the device name,
	// it should return the device information.
	GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error)
}

type deviceMapperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceMapperServiceClient(cc grpc.ClientConnInterface) DeviceMapperServiceClient {
	return &deviceMapperServiceClient{cc}
}

func (c *deviceMapperServiceClient) RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*RegisterDeviceResponse, error) {
	out := new(RegisterDeviceResponse)
	err := c.cc.Invoke(ctx, DeviceMapperService_RegisterDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMapperServiceClient) RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error) {
	out := new(RemoveDeviceResponse)
	err := c.cc.Invoke(ctx, DeviceMapperService_RemoveDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMapperServiceClient) UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error) {
	out := new(UpdateDeviceResponse)
	err := c.cc.Invoke(ctx, DeviceMapperService_UpdateDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMapperServiceClient) CreateDeviceModel(ctx context.Context, in *CreateDeviceModelRequest, opts ...grpc.CallOption) (*CreateDeviceModelResponse, error) {
	out := new(CreateDeviceModelResponse)
	err := c.cc.Invoke(ctx, DeviceMapperService_CreateDeviceModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMapperServiceClient) RemoveDeviceModel(ctx context.Context, in *RemoveDeviceModelRequest, opts ...grpc.CallOption) (*RemoveDeviceModelResponse, error) {
	out := new(RemoveDeviceModelResponse)
	err := c.cc.Invoke(ctx, DeviceMapperService_RemoveDeviceModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMapperServiceClient) UpdateDeviceModel(ctx context.Context, in *UpdateDeviceModelRequest, opts ...grpc.CallOption) (*UpdateDeviceModelResponse, error) {
	out := new(UpdateDeviceModelResponse)
	err := c.cc.Invoke(ctx, DeviceMapperService_UpdateDeviceModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMapperServiceClient) GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error) {
	out := new(GetDeviceResponse)
	err := c.cc.Invoke(ctx, DeviceMapperService_GetDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceMapperServiceServer is the server API for DeviceMapperService service.
// All implementations must embed UnimplementedDeviceMapperServiceServer
// for forward compatibility
type DeviceMapperServiceServer interface {
	// RegisterDevice registers a device to the device mapper.
	// Device manager registers a device instance with the information of device
	// to the mapper through the interface of RegisterDevice.
	// When the mapper gets the request of register with device information,
	// it should add the device to the device list and connect to the real physical device via the specific protocol.
	RegisterDevice(context.Context, *RegisterDeviceRequest) (*RegisterDeviceResponse, error)
	// RemoveDevice unregisters a device to the device mapper.
	// Device manager unregisters a device instance with the name of device
	// to the mapper through the interface of RemoveDevice.
	// When the mapper gets the request of unregister with device name,
	// it should remove the device from the device list and disconnect to the real physical device.
	RemoveDevice(context.Context, *RemoveDeviceRequest) (*RemoveDeviceResponse, error)
	// UpdateDevice updates a device to the device mapper
	// Device manager updates the information of a device used by the mapper
	// through the interface of UpdateDevice.
	// The information of a device includes the meta data and the status data of a device.
	// When the mapper gets the request of updating with the information of a device,
	// it should update the device of the device list and connect to the real physical device via the updated information.
	UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error)
	// CreateDeviceModel creates a device model to the device mapper.
	// Device manager sends the information of device model to the mapper
	// through the interface of CreateDeviceModel.
	// When the mapper gets the request of creating with the information of device model,
	// it should create a new device model to the list of device models.
	CreateDeviceModel(context.Context, *CreateDeviceModelRequest) (*CreateDeviceModelResponse, error)
	// RemoveDeviceModel remove a device model to the device mapper.
	// Device manager sends the name of device model to the mapper
	// through the interface of RemoveDeviceModel.
	// When the mapper gets the request of removing with the name of device model,
	// it should remove the device model to the list of device models.
	RemoveDeviceModel(context.Context, *RemoveDeviceModelRequest) (*RemoveDeviceModelResponse, error)
	// UpdateDeviceModel update a device model to the device mapper.
	// Device manager sends the information of device model to the mapper
	// through the interface of UpdateDeviceModel.
	// When the mapper gets the request of updating with the information of device model,
	// it should update the device model to the list of device models.
	UpdateDeviceModel(context.Context, *UpdateDeviceModelRequest) (*UpdateDeviceModelResponse, error)
	// GetDevice get the information of a device from the device mapper.
	// Device sends the request of querying device information with the device name to the mapper
	// through the interface of GetDevice.
	// When the mapper gets the request of querying with the device name,
	// it should return the device information.
	GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	mustEmbedUnimplementedDeviceMapperServiceServer()
}

// UnimplementedDeviceMapperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceMapperServiceServer struct {
}

func (UnimplementedDeviceMapperServiceServer) RegisterDevice(context.Context, *RegisterDeviceRequest) (*RegisterDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevice not implemented")
}
func (UnimplementedDeviceMapperServiceServer) RemoveDevice(context.Context, *RemoveDeviceRequest) (*RemoveDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDevice not implemented")
}
func (UnimplementedDeviceMapperServiceServer) UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (UnimplementedDeviceMapperServiceServer) CreateDeviceModel(context.Context, *CreateDeviceModelRequest) (*CreateDeviceModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceModel not implemented")
}
func (UnimplementedDeviceMapperServiceServer) RemoveDeviceModel(context.Context, *RemoveDeviceModelRequest) (*RemoveDeviceModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDeviceModel not implemented")
}
func (UnimplementedDeviceMapperServiceServer) UpdateDeviceModel(context.Context, *UpdateDeviceModelRequest) (*UpdateDeviceModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceModel not implemented")
}
func (UnimplementedDeviceMapperServiceServer) GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (UnimplementedDeviceMapperServiceServer) mustEmbedUnimplementedDeviceMapperServiceServer() {}

// UnsafeDeviceMapperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceMapperServiceServer will
// result in compilation errors.
type UnsafeDeviceMapperServiceServer interface {
	mustEmbedUnimplementedDeviceMapperServiceServer()
}

func RegisterDeviceMapperServiceServer(s grpc.ServiceRegistrar, srv DeviceMapperServiceServer) {
	s.RegisterService(&DeviceMapperService_ServiceDesc, srv)
}

func _DeviceMapperService_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMapperServiceServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceMapperService_RegisterDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMapperServiceServer).RegisterDevice(ctx, req.(*RegisterDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMapperService_RemoveDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMapperServiceServer).RemoveDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceMapperService_RemoveDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMapperServiceServer).RemoveDevice(ctx, req.(*RemoveDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMapperService_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMapperServiceServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceMapperService_UpdateDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMapperServiceServer).UpdateDevice(ctx, req.(*UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMapperService_CreateDeviceModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMapperServiceServer).CreateDeviceModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceMapperService_CreateDeviceModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMapperServiceServer).CreateDeviceModel(ctx, req.(*CreateDeviceModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMapperService_RemoveDeviceModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMapperServiceServer).RemoveDeviceModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceMapperService_RemoveDeviceModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMapperServiceServer).RemoveDeviceModel(ctx, req.(*RemoveDeviceModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMapperService_UpdateDeviceModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMapperServiceServer).UpdateDeviceModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceMapperService_UpdateDeviceModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMapperServiceServer).UpdateDeviceModel(ctx, req.(*UpdateDeviceModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMapperService_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMapperServiceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceMapperService_GetDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMapperServiceServer).GetDevice(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceMapperService_ServiceDesc is the grpc.ServiceDesc for DeviceMapperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceMapperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1beta1.DeviceMapperService",
	HandlerType: (*DeviceMapperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevice",
			Handler:    _DeviceMapperService_RegisterDevice_Handler,
		},
		{
			MethodName: "RemoveDevice",
			Handler:    _DeviceMapperService_RemoveDevice_Handler,
		},
		{
			MethodName: "UpdateDevice",
			Handler:    _DeviceMapperService_UpdateDevice_Handler,
		},
		{
			MethodName: "CreateDeviceModel",
			Handler:    _DeviceMapperService_CreateDeviceModel_Handler,
		},
		{
			MethodName: "RemoveDeviceModel",
			Handler:    _DeviceMapperService_RemoveDeviceModel_Handler,
		},
		{
			MethodName: "UpdateDeviceModel",
			Handler:    _DeviceMapperService_UpdateDeviceModel_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _DeviceMapperService_GetDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
