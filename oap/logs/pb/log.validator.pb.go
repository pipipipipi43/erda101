// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log.proto

package pb

import (
	fmt "fmt"
	_ "github.com/erda-project/erda-proto-go/oap/common/pb"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/structpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Log) Validate() error {
	if this.Relations != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relations); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relations", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
