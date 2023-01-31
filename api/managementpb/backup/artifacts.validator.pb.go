// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/backup/artifacts.proto

package backupv1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *Artifact) Validate() error {
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}

func (this *ListArtifactsRequest) Validate() error {
	return nil
}

func (this *ListArtifactsResponse) Validate() error {
	for _, item := range this.Artifacts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Artifacts", err)
			}
		}
	}
	return nil
}

func (this *DeleteArtifactRequest) Validate() error {
	return nil
}

func (this *DeleteArtifactResponse) Validate() error {
	return nil
}

func (this *PitrTimerange) Validate() error {
	if this.StartTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTimestamp", err)
		}
	}
	if this.EndTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndTimestamp", err)
		}
	}
	return nil
}

func (this *ListPitrTimerangesRequest) Validate() error {
	return nil
}

func (this *ListPitrTimerangesResponse) Validate() error {
	for _, item := range this.Timeranges {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Timeranges", err)
			}
		}
	}
	return nil
}
