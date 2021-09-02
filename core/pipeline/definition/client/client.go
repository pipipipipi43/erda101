// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: definition.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/definition/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// DefinitionService definition.proto
	DefinitionService() pb.DefinitionServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		definitionService: pb.NewDefinitionServiceClient(cc),
	}
}

type serviceClients struct {
	definitionService pb.DefinitionServiceClient
}

func (c *serviceClients) DefinitionService() pb.DefinitionServiceClient {
	return c.definitionService
}

type definitionServiceWrapper struct {
	client pb.DefinitionServiceClient
	opts   []grpc1.CallOption
}

func (s *definitionServiceWrapper) Process(ctx context.Context, req *pb.PipelineDefinitionProcessRequest) (*pb.PipelineDefinitionProcessResponse, error) {
	return s.client.Process(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *definitionServiceWrapper) Version(ctx context.Context, req *pb.PipelineDefinitionProcessVersionRequest) (*pb.PipelineDefinitionProcessVersionResponse, error) {
	return s.client.Version(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}