// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tenant.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/descriptorpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateTenantRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.TenantType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantType", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantType))
	}
	for _, item := range this.Workspaces {
		if item == "" {
			return github_com_mwitkow_go_proto_validators.FieldError("Workspaces", fmt.Errorf(`value '%v' must not be an empty string`, item))
		}
	}
	return nil
}
func (this *CreateTenantResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTenantRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.TenantType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantType", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantType))
	}
	if this.Workspace == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Workspace", fmt.Errorf(`value '%v' must not be an empty string`, this.Workspace))
	}
	return nil
}
func (this *GetTenantResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteTenantRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	if this.TenantType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantType", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantType))
	}
	if this.Workspace == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Workspace", fmt.Errorf(`value '%v' must not be an empty string`, this.Workspace))
	}
	return nil
}
func (this *DeleteTenantResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Tenant) Validate() error {
	return nil
}
