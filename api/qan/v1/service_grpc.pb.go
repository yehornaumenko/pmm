// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: qan/v1/service.proto

package qanv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	QANService_GetReport_FullMethodName                   = "/qan.v1.QANService/GetReport"
	QANService_GetFilteredMetricsNames_FullMethodName     = "/qan.v1.QANService/GetFilteredMetricsNames"
	QANService_GetMetricsNames_FullMethodName             = "/qan.v1.QANService/GetMetricsNames"
	QANService_GetMetrics_FullMethodName                  = "/qan.v1.QANService/GetMetrics"
	QANService_GetLabels_FullMethodName                   = "/qan.v1.QANService/GetLabels"
	QANService_GetHistogram_FullMethodName                = "/qan.v1.QANService/GetHistogram"
	QANService_ExplainFingerprintByQueryID_FullMethodName = "/qan.v1.QANService/ExplainFingerprintByQueryID"
	QANService_GetQueryPlan_FullMethodName                = "/qan.v1.QANService/GetQueryPlan"
	QANService_QueryExists_FullMethodName                 = "/qan.v1.QANService/QueryExists"
	QANService_SchemaByQueryID_FullMethodName             = "/qan.v1.QANService/SchemaByQueryID"
	QANService_GetQueryExample_FullMethodName             = "/qan.v1.QANService/GetQueryExample"
)

// QANServiceClient is the client API for QANService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// QANService provides an API to interact with PMM Query Analytics.
type QANServiceClient interface {
	// GetReport returns a list of metrics grouped by queryid or other dimensions.
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error)
	// Get provides a map of metrics names.
	GetFilteredMetricsNames(ctx context.Context, in *GetFilteredMetricsNamesRequest, opts ...grpc.CallOption) (*GetFilteredMetricsNamesResponse, error)
	// GetMetricsNames provides a map of metrics names.
	GetMetricsNames(ctx context.Context, in *GetMetricsNamesRequest, opts ...grpc.CallOption) (*GetMetricsNamesResponse, error)
	// GetMetrics returns a map of metrics for specific filtering.
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// GetLabels return a list of labels for object details.
	GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error)
	// GetHistogram returns histogram items for specific filtering.
	GetHistogram(ctx context.Context, in *GetHistogramRequest, opts ...grpc.CallOption) (*GetHistogramResponse, error)
	// ExplainFingerprintByQueryID returns an explain fingerprint for given query ID.
	ExplainFingerprintByQueryID(ctx context.Context, in *ExplainFingerprintByQueryIDRequest, opts ...grpc.CallOption) (*ExplainFingerprintByQueryIDResponse, error)
	// GetQueryPlan returns a query plan and plan id for specific filtering.
	GetQueryPlan(ctx context.Context, in *GetQueryPlanRequest, opts ...grpc.CallOption) (*GetQueryPlanResponse, error)
	// QueryExists checks if query exists in clickhouse.
	QueryExists(ctx context.Context, in *QueryExistsRequest, opts ...grpc.CallOption) (*QueryExistsResponse, error)
	// SchemaByQueryID returns the schema for a given queryID and serviceID.
	SchemaByQueryID(ctx context.Context, in *SchemaByQueryIDRequest, opts ...grpc.CallOption) (*SchemaByQueryIDResponse, error)
	// GetQueryExample returns a list of query examples.
	GetQueryExample(ctx context.Context, in *GetQueryExampleRequest, opts ...grpc.CallOption) (*GetQueryExampleResponse, error)
}

type qANServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQANServiceClient(cc grpc.ClientConnInterface) QANServiceClient {
	return &qANServiceClient{cc}
}

func (c *qANServiceClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReportResponse)
	err := c.cc.Invoke(ctx, QANService_GetReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) GetFilteredMetricsNames(ctx context.Context, in *GetFilteredMetricsNamesRequest, opts ...grpc.CallOption) (*GetFilteredMetricsNamesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFilteredMetricsNamesResponse)
	err := c.cc.Invoke(ctx, QANService_GetFilteredMetricsNames_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) GetMetricsNames(ctx context.Context, in *GetMetricsNamesRequest, opts ...grpc.CallOption) (*GetMetricsNamesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricsNamesResponse)
	err := c.cc.Invoke(ctx, QANService_GetMetricsNames_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, QANService_GetMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLabelsResponse)
	err := c.cc.Invoke(ctx, QANService_GetLabels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) GetHistogram(ctx context.Context, in *GetHistogramRequest, opts ...grpc.CallOption) (*GetHistogramResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHistogramResponse)
	err := c.cc.Invoke(ctx, QANService_GetHistogram_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) ExplainFingerprintByQueryID(ctx context.Context, in *ExplainFingerprintByQueryIDRequest, opts ...grpc.CallOption) (*ExplainFingerprintByQueryIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExplainFingerprintByQueryIDResponse)
	err := c.cc.Invoke(ctx, QANService_ExplainFingerprintByQueryID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) GetQueryPlan(ctx context.Context, in *GetQueryPlanRequest, opts ...grpc.CallOption) (*GetQueryPlanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQueryPlanResponse)
	err := c.cc.Invoke(ctx, QANService_GetQueryPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) QueryExists(ctx context.Context, in *QueryExistsRequest, opts ...grpc.CallOption) (*QueryExistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryExistsResponse)
	err := c.cc.Invoke(ctx, QANService_QueryExists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) SchemaByQueryID(ctx context.Context, in *SchemaByQueryIDRequest, opts ...grpc.CallOption) (*SchemaByQueryIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SchemaByQueryIDResponse)
	err := c.cc.Invoke(ctx, QANService_SchemaByQueryID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qANServiceClient) GetQueryExample(ctx context.Context, in *GetQueryExampleRequest, opts ...grpc.CallOption) (*GetQueryExampleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQueryExampleResponse)
	err := c.cc.Invoke(ctx, QANService_GetQueryExample_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QANServiceServer is the server API for QANService service.
// All implementations must embed UnimplementedQANServiceServer
// for forward compatibility.
//
// QANService provides an API to interact with PMM Query Analytics.
type QANServiceServer interface {
	// GetReport returns a list of metrics grouped by queryid or other dimensions.
	GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error)
	// Get provides a map of metrics names.
	GetFilteredMetricsNames(context.Context, *GetFilteredMetricsNamesRequest) (*GetFilteredMetricsNamesResponse, error)
	// GetMetricsNames provides a map of metrics names.
	GetMetricsNames(context.Context, *GetMetricsNamesRequest) (*GetMetricsNamesResponse, error)
	// GetMetrics returns a map of metrics for specific filtering.
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// GetLabels return a list of labels for object details.
	GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error)
	// GetHistogram returns histogram items for specific filtering.
	GetHistogram(context.Context, *GetHistogramRequest) (*GetHistogramResponse, error)
	// ExplainFingerprintByQueryID returns an explain fingerprint for given query ID.
	ExplainFingerprintByQueryID(context.Context, *ExplainFingerprintByQueryIDRequest) (*ExplainFingerprintByQueryIDResponse, error)
	// GetQueryPlan returns a query plan and plan id for specific filtering.
	GetQueryPlan(context.Context, *GetQueryPlanRequest) (*GetQueryPlanResponse, error)
	// QueryExists checks if query exists in clickhouse.
	QueryExists(context.Context, *QueryExistsRequest) (*QueryExistsResponse, error)
	// SchemaByQueryID returns the schema for a given queryID and serviceID.
	SchemaByQueryID(context.Context, *SchemaByQueryIDRequest) (*SchemaByQueryIDResponse, error)
	// GetQueryExample returns a list of query examples.
	GetQueryExample(context.Context, *GetQueryExampleRequest) (*GetQueryExampleResponse, error)
	mustEmbedUnimplementedQANServiceServer()
}

// UnimplementedQANServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQANServiceServer struct{}

func (UnimplementedQANServiceServer) GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}

func (UnimplementedQANServiceServer) GetFilteredMetricsNames(context.Context, *GetFilteredMetricsNamesRequest) (*GetFilteredMetricsNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilteredMetricsNames not implemented")
}

func (UnimplementedQANServiceServer) GetMetricsNames(context.Context, *GetMetricsNamesRequest) (*GetMetricsNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricsNames not implemented")
}

func (UnimplementedQANServiceServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}

func (UnimplementedQANServiceServer) GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}

func (UnimplementedQANServiceServer) GetHistogram(context.Context, *GetHistogramRequest) (*GetHistogramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistogram not implemented")
}

func (UnimplementedQANServiceServer) ExplainFingerprintByQueryID(context.Context, *ExplainFingerprintByQueryIDRequest) (*ExplainFingerprintByQueryIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExplainFingerprintByQueryID not implemented")
}

func (UnimplementedQANServiceServer) GetQueryPlan(context.Context, *GetQueryPlanRequest) (*GetQueryPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryPlan not implemented")
}

func (UnimplementedQANServiceServer) QueryExists(context.Context, *QueryExistsRequest) (*QueryExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryExists not implemented")
}

func (UnimplementedQANServiceServer) SchemaByQueryID(context.Context, *SchemaByQueryIDRequest) (*SchemaByQueryIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchemaByQueryID not implemented")
}

func (UnimplementedQANServiceServer) GetQueryExample(context.Context, *GetQueryExampleRequest) (*GetQueryExampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryExample not implemented")
}
func (UnimplementedQANServiceServer) mustEmbedUnimplementedQANServiceServer() {}
func (UnimplementedQANServiceServer) testEmbeddedByValue()                    {}

// UnsafeQANServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QANServiceServer will
// result in compilation errors.
type UnsafeQANServiceServer interface {
	mustEmbedUnimplementedQANServiceServer()
}

func RegisterQANServiceServer(s grpc.ServiceRegistrar, srv QANServiceServer) {
	// If the following call pancis, it indicates UnimplementedQANServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QANService_ServiceDesc, srv)
}

func _QANService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_GetReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_GetFilteredMetricsNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilteredMetricsNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).GetFilteredMetricsNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_GetFilteredMetricsNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).GetFilteredMetricsNames(ctx, req.(*GetFilteredMetricsNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_GetMetricsNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).GetMetricsNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_GetMetricsNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).GetMetricsNames(ctx, req.(*GetMetricsNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_GetLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).GetLabels(ctx, req.(*GetLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_GetHistogram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistogramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).GetHistogram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_GetHistogram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).GetHistogram(ctx, req.(*GetHistogramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_ExplainFingerprintByQueryID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExplainFingerprintByQueryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).ExplainFingerprintByQueryID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_ExplainFingerprintByQueryID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).ExplainFingerprintByQueryID(ctx, req.(*ExplainFingerprintByQueryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_GetQueryPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).GetQueryPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_GetQueryPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).GetQueryPlan(ctx, req.(*GetQueryPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_QueryExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).QueryExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_QueryExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).QueryExists(ctx, req.(*QueryExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_SchemaByQueryID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaByQueryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).SchemaByQueryID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_SchemaByQueryID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).SchemaByQueryID(ctx, req.(*SchemaByQueryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QANService_GetQueryExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QANServiceServer).GetQueryExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QANService_GetQueryExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QANServiceServer).GetQueryExample(ctx, req.(*GetQueryExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QANService_ServiceDesc is the grpc.ServiceDesc for QANService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QANService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qan.v1.QANService",
	HandlerType: (*QANServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReport",
			Handler:    _QANService_GetReport_Handler,
		},
		{
			MethodName: "GetFilteredMetricsNames",
			Handler:    _QANService_GetFilteredMetricsNames_Handler,
		},
		{
			MethodName: "GetMetricsNames",
			Handler:    _QANService_GetMetricsNames_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _QANService_GetMetrics_Handler,
		},
		{
			MethodName: "GetLabels",
			Handler:    _QANService_GetLabels_Handler,
		},
		{
			MethodName: "GetHistogram",
			Handler:    _QANService_GetHistogram_Handler,
		},
		{
			MethodName: "ExplainFingerprintByQueryID",
			Handler:    _QANService_ExplainFingerprintByQueryID_Handler,
		},
		{
			MethodName: "GetQueryPlan",
			Handler:    _QANService_GetQueryPlan_Handler,
		},
		{
			MethodName: "QueryExists",
			Handler:    _QANService_QueryExists_Handler,
		},
		{
			MethodName: "SchemaByQueryID",
			Handler:    _QANService_SchemaByQueryID_Handler,
		},
		{
			MethodName: "GetQueryExample",
			Handler:    _QANService_GetQueryExample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qan/v1/service.proto",
}
