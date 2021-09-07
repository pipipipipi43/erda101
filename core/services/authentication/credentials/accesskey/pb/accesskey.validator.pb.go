// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accesskey.proto

package pb

import (
	fmt "fmt"
	_ "github.com/erda-project/erda-proto-go/common/pb"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *QueryAccessKeysRequest) Validate() error {
	return nil
}
func (this *QueryAccessKeysResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetAccessKeyRequest) Validate() error {
	return nil
}
func (this *GetAccessKeyResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateAccessKeyRequest) Validate() error {
	return nil
}
func (this *CreateAccessKeyResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateAccessKeyRequest) Validate() error {
	return nil
}
func (this *UpdateAccessKeyResponse) Validate() error {
	return nil
}
func (this *DeleteAccessKeyRequest) Validate() error {
	return nil
}
func (this *DeleteAccessKeyResponse) Validate() error {
	return nil
}
func (this *AccessKeysItem) Validate() error {
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}
func (this *SubjectTypeEnum) Validate() error {
	return nil
}
func (this *StatusEnum) Validate() error {
	return nil
}
