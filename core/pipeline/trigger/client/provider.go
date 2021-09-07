// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: trigger.proto

package client

import (
	fmt "fmt"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/trigger/pb"
	grpc1 "google.golang.org/grpc"
	reflect "reflect"
	strings "strings"
)

var dependencies = []string{
	"grpc-client@erda.core.pipeline.trigger",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType              = reflect.TypeOf((*Client)(nil)).Elem()
	triggerServiceClientType = reflect.TypeOf((*pb.TriggerServiceClient)(nil)).Elem()
	triggerServiceServerType = reflect.TypeOf((*pb.TriggerServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.pipeline.trigger-client":
		return p.client
	case "erda.core.pipeline.trigger.TriggerService":
		return &triggerServiceWrapper{client: p.client.TriggerService(), opts: opts}
	case "erda.core.pipeline.trigger.TriggerService.client":
		return p.client.TriggerService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case triggerServiceClientType:
		return p.client.TriggerService()
	case triggerServiceServerType:
		return &triggerServiceWrapper{client: p.client.TriggerService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.pipeline.trigger-client", &servicehub.Spec{
		Services: []string{
			"erda.core.pipeline.trigger.TriggerService",
			"erda.core.pipeline.trigger-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			triggerServiceClientType,
			// server types
			triggerServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}