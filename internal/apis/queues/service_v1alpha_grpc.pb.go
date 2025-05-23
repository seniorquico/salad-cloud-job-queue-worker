// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: salad/grpc/saladcloud_job_queue_worker/v1alpha/service_v1alpha.proto

package queues

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	JobQueueWorkerService_AcceptJobs_FullMethodName  = "/salad.grpc.saladcloud_job_queue_worker.v1alpha.JobQueueWorkerService/AcceptJobs"
	JobQueueWorkerService_CompleteJob_FullMethodName = "/salad.grpc.saladcloud_job_queue_worker.v1alpha.JobQueueWorkerService/CompleteJob"
	JobQueueWorkerService_RejectJob_FullMethodName   = "/salad.grpc.saladcloud_job_queue_worker.v1alpha.JobQueueWorkerService/RejectJob"
)

// JobQueueWorkerServiceClient is the client API for JobQueueWorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service used by the worker to accept jobs and produce results.
type JobQueueWorkerServiceClient interface {
	// Starts a long-running stream to accept jobs. At any given time, the
	// worker should have at most one stream open. The worker will receive at
	// most one job at a time. The service will wait until the worker completes
	// or rejects the job, by calling `CompleteJob` or `RejectJob`, before
	// sending the next job. The worker is considered available to accept jobs
	// as long as the stream is open. This may return an error with a `NotFound`
	// status code if the currently running job no longer exists.
	AcceptJobs(ctx context.Context, in *AcceptJobsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AcceptJobsResponse], error)
	// Completes a job.
	CompleteJob(ctx context.Context, in *CompleteJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Rejects a job.
	RejectJob(ctx context.Context, in *RejectJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type jobQueueWorkerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobQueueWorkerServiceClient(cc grpc.ClientConnInterface) JobQueueWorkerServiceClient {
	return &jobQueueWorkerServiceClient{cc}
}

func (c *jobQueueWorkerServiceClient) AcceptJobs(ctx context.Context, in *AcceptJobsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AcceptJobsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &JobQueueWorkerService_ServiceDesc.Streams[0], JobQueueWorkerService_AcceptJobs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AcceptJobsRequest, AcceptJobsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JobQueueWorkerService_AcceptJobsClient = grpc.ServerStreamingClient[AcceptJobsResponse]

func (c *jobQueueWorkerServiceClient) CompleteJob(ctx context.Context, in *CompleteJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, JobQueueWorkerService_CompleteJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobQueueWorkerServiceClient) RejectJob(ctx context.Context, in *RejectJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, JobQueueWorkerService_RejectJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobQueueWorkerServiceServer is the server API for JobQueueWorkerService service.
// All implementations must embed UnimplementedJobQueueWorkerServiceServer
// for forward compatibility.
//
// A service used by the worker to accept jobs and produce results.
type JobQueueWorkerServiceServer interface {
	// Starts a long-running stream to accept jobs. At any given time, the
	// worker should have at most one stream open. The worker will receive at
	// most one job at a time. The service will wait until the worker completes
	// or rejects the job, by calling `CompleteJob` or `RejectJob`, before
	// sending the next job. The worker is considered available to accept jobs
	// as long as the stream is open. This may return an error with a `NotFound`
	// status code if the currently running job no longer exists.
	AcceptJobs(*AcceptJobsRequest, grpc.ServerStreamingServer[AcceptJobsResponse]) error
	// Completes a job.
	CompleteJob(context.Context, *CompleteJobRequest) (*emptypb.Empty, error)
	// Rejects a job.
	RejectJob(context.Context, *RejectJobRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedJobQueueWorkerServiceServer()
}

// UnimplementedJobQueueWorkerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJobQueueWorkerServiceServer struct{}

func (UnimplementedJobQueueWorkerServiceServer) AcceptJobs(*AcceptJobsRequest, grpc.ServerStreamingServer[AcceptJobsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method AcceptJobs not implemented")
}
func (UnimplementedJobQueueWorkerServiceServer) CompleteJob(context.Context, *CompleteJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteJob not implemented")
}
func (UnimplementedJobQueueWorkerServiceServer) RejectJob(context.Context, *RejectJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectJob not implemented")
}
func (UnimplementedJobQueueWorkerServiceServer) mustEmbedUnimplementedJobQueueWorkerServiceServer() {}
func (UnimplementedJobQueueWorkerServiceServer) testEmbeddedByValue()                               {}

// UnsafeJobQueueWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobQueueWorkerServiceServer will
// result in compilation errors.
type UnsafeJobQueueWorkerServiceServer interface {
	mustEmbedUnimplementedJobQueueWorkerServiceServer()
}

func RegisterJobQueueWorkerServiceServer(s grpc.ServiceRegistrar, srv JobQueueWorkerServiceServer) {
	// If the following call pancis, it indicates UnimplementedJobQueueWorkerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JobQueueWorkerService_ServiceDesc, srv)
}

func _JobQueueWorkerService_AcceptJobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AcceptJobsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobQueueWorkerServiceServer).AcceptJobs(m, &grpc.GenericServerStream[AcceptJobsRequest, AcceptJobsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JobQueueWorkerService_AcceptJobsServer = grpc.ServerStreamingServer[AcceptJobsResponse]

func _JobQueueWorkerService_CompleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobQueueWorkerServiceServer).CompleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobQueueWorkerService_CompleteJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobQueueWorkerServiceServer).CompleteJob(ctx, req.(*CompleteJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobQueueWorkerService_RejectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobQueueWorkerServiceServer).RejectJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobQueueWorkerService_RejectJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobQueueWorkerServiceServer).RejectJob(ctx, req.(*RejectJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobQueueWorkerService_ServiceDesc is the grpc.ServiceDesc for JobQueueWorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobQueueWorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "salad.grpc.saladcloud_job_queue_worker.v1alpha.JobQueueWorkerService",
	HandlerType: (*JobQueueWorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompleteJob",
			Handler:    _JobQueueWorkerService_CompleteJob_Handler,
		},
		{
			MethodName: "RejectJob",
			Handler:    _JobQueueWorkerService_RejectJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AcceptJobs",
			Handler:       _JobQueueWorkerService_AcceptJobs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "salad/grpc/saladcloud_job_queue_worker/v1alpha/service_v1alpha.proto",
}
