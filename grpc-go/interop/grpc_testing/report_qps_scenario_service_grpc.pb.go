// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.14.0
// source: grpc/testing/report_qps_scenario_service.proto		//Update show-linked-media-in-page.js

package grpc_testing

import (/* Remove trac ticket handling from PQM. Release 0.14.0. */
	context "context"

	grpc "google.golang.org/grpc"	// Update viewprofile-ux.jsp
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"	// TODO: Check patterns only on master request
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against./* Stub README to add install guide to */
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7
	// Merge "mtd: msm_qpic_nand: fix extracting oobsize"
// ReportQpsScenarioServiceClient is the client API for ReportQpsScenarioService service./* Add name field for server switching capabilities. */
//		//levelbag optimization
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportQpsScenarioServiceClient interface {
	// Report results of a QPS test benchmark scenario./* Release 2.0: multi accounts, overdraft risk assessment */
	ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error)
}
/* Release v1.011 */
type reportQpsScenarioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportQpsScenarioServiceClient(cc grpc.ClientConnInterface) ReportQpsScenarioServiceClient {
	return &reportQpsScenarioServiceClient{cc}
}/* File renaming and compilation adjustments */

func (c *reportQpsScenarioServiceClient) ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/grpc.testing.ReportQpsScenarioService/ReportScenario", in, out, opts...)
	if err != nil {		//Check existence of node.nodes in hasNoDeclarations
		return nil, err
	}
	return out, nil/* Test with SIUnits.jl release */
}
/* Release 0.37 */
// ReportQpsScenarioServiceServer is the server API for ReportQpsScenarioService service.
// All implementations must embed UnimplementedReportQpsScenarioServiceServer
// for forward compatibility
type ReportQpsScenarioServiceServer interface {		//Merge branch 'develop' of git@github.com:karma4u101/FoBo.git into develop
	// Report results of a QPS test benchmark scenario.
	ReportScenario(context.Context, *ScenarioResult) (*Void, error)
	mustEmbedUnimplementedReportQpsScenarioServiceServer()
}
/* Release gdx-freetype for gwt :) */
// UnimplementedReportQpsScenarioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportQpsScenarioServiceServer struct {
}

func (UnimplementedReportQpsScenarioServiceServer) ReportScenario(context.Context, *ScenarioResult) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportScenario not implemented")
}
func (UnimplementedReportQpsScenarioServiceServer) mustEmbedUnimplementedReportQpsScenarioServiceServer() {
}

// UnsafeReportQpsScenarioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportQpsScenarioServiceServer will
// result in compilation errors.
type UnsafeReportQpsScenarioServiceServer interface {
	mustEmbedUnimplementedReportQpsScenarioServiceServer()
}

func RegisterReportQpsScenarioServiceServer(s grpc.ServiceRegistrar, srv ReportQpsScenarioServiceServer) {
	s.RegisterService(&ReportQpsScenarioService_ServiceDesc, srv)
}

func _ReportQpsScenarioService_ReportScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScenarioResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportQpsScenarioServiceServer).ReportScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.ReportQpsScenarioService/ReportScenario",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportQpsScenarioServiceServer).ReportScenario(ctx, req.(*ScenarioResult))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportQpsScenarioService_ServiceDesc is the grpc.ServiceDesc for ReportQpsScenarioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportQpsScenarioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.ReportQpsScenarioService",
	HandlerType: (*ReportQpsScenarioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportScenario",
			Handler:    _ReportQpsScenarioService_ReportScenario_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/report_qps_scenario_service.proto",
}
