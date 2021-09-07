// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: expression.proto

package client

import (
	fmt "fmt"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/expression/pb"
	grpc1 "google.golang.org/grpc"
	reflect "reflect"
	strings "strings"
)

var dependencies = []string{
	"grpc-client@erda.msp.apm.expression",
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
	clientsType                 = reflect.TypeOf((*Client)(nil)).Elem()
	expressionServiceClientType = reflect.TypeOf((*pb.ExpressionServiceClient)(nil)).Elem()
	expressionServiceServerType = reflect.TypeOf((*pb.ExpressionServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.msp.apm.expression-client":
		return p.client
	case "erda.msp.apm.expression.ExpressionService":
		return &expressionServiceWrapper{client: p.client.ExpressionService(), opts: opts}
	case "erda.msp.apm.expression.ExpressionService.client":
		return p.client.ExpressionService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case expressionServiceClientType:
		return p.client.ExpressionService()
	case expressionServiceServerType:
		return &expressionServiceWrapper{client: p.client.ExpressionService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.apm.expression-client", &servicehub.Spec{
		Services: []string{
			"erda.msp.apm.expression.ExpressionService",
			"erda.msp.apm.expression-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			expressionServiceClientType,
			// server types
			expressionServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
