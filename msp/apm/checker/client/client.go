// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: checker.proto, checker_v1.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/checker/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// CheckerService checker.proto
	CheckerService() pb.CheckerServiceClient
	// CheckerV1Service checker_v1.proto
	CheckerV1Service() pb.CheckerV1ServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		checkerService:   pb.NewCheckerServiceClient(cc),
		checkerV1Service: pb.NewCheckerV1ServiceClient(cc),
	}
}

type serviceClients struct {
	checkerService   pb.CheckerServiceClient
	checkerV1Service pb.CheckerV1ServiceClient
}

func (c *serviceClients) CheckerService() pb.CheckerServiceClient {
	return c.checkerService
}

func (c *serviceClients) CheckerV1Service() pb.CheckerV1ServiceClient {
	return c.checkerV1Service
}

type checkerServiceWrapper struct {
	client pb.CheckerServiceClient
	opts   []grpc1.CallOption
}

func (s *checkerServiceWrapper) CreateChecker(ctx context.Context, req *pb.CreateCheckerRequest) (*pb.CreateCheckerResponse, error) {
	return s.client.CreateChecker(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerServiceWrapper) UpdateChecker(ctx context.Context, req *pb.UpdateCheckerRequest) (*pb.UpdateCheckerResponse, error) {
	return s.client.UpdateChecker(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerServiceWrapper) DeleteChecker(ctx context.Context, req *pb.UpdateCheckerRequest) (*pb.UpdateCheckerResponse, error) {
	return s.client.DeleteChecker(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerServiceWrapper) ListCheckers(ctx context.Context, req *pb.ListCheckersRequest) (*pb.ListCheckersResponse, error) {
	return s.client.ListCheckers(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerServiceWrapper) DescribeCheckers(ctx context.Context, req *pb.DescribeCheckersRequest) (*pb.DescribeCheckersResponse, error) {
	return s.client.DescribeCheckers(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerServiceWrapper) DescribeChecker(ctx context.Context, req *pb.DescribeCheckerRequest) (*pb.DescribeCheckerResponse, error) {
	return s.client.DescribeChecker(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

type checkerV1ServiceWrapper struct {
	client pb.CheckerV1ServiceClient
	opts   []grpc1.CallOption
}

func (s *checkerV1ServiceWrapper) CreateCheckerV1(ctx context.Context, req *pb.CreateCheckerV1Request) (*pb.CreateCheckerV1Response, error) {
	return s.client.CreateCheckerV1(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerV1ServiceWrapper) UpdateCheckerV1(ctx context.Context, req *pb.UpdateCheckerV1Request) (*pb.UpdateCheckerV1Response, error) {
	return s.client.UpdateCheckerV1(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerV1ServiceWrapper) DeleteCheckerV1(ctx context.Context, req *pb.DeleteCheckerV1Request) (*pb.DeleteCheckerV1Response, error) {
	return s.client.DeleteCheckerV1(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerV1ServiceWrapper) GetCheckerV1(ctx context.Context, req *pb.GetCheckerV1Request) (*pb.GetCheckerV1Response, error) {
	return s.client.GetCheckerV1(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerV1ServiceWrapper) DescribeCheckersV1(ctx context.Context, req *pb.DescribeCheckersV1Request) (*pb.DescribeCheckersV1Response, error) {
	return s.client.DescribeCheckersV1(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerV1ServiceWrapper) DescribeCheckerV1(ctx context.Context, req *pb.DescribeCheckerV1Request) (*pb.DescribeCheckerV1Response, error) {
	return s.client.DescribeCheckerV1(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerV1ServiceWrapper) GetCheckerStatusV1(ctx context.Context, req *pb.GetCheckerStatusV1Request) (*pb.GetCheckerStatusV1Response, error) {
	return s.client.GetCheckerStatusV1(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *checkerV1ServiceWrapper) GetCheckerIssuesV1(ctx context.Context, req *pb.GetCheckerIssuesV1Request) (*pb.GetCheckerIssuesV1Response, error) {
	return s.client.GetCheckerIssuesV1(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
