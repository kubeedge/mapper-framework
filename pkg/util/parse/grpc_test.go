package parse

import (
	"testing"

	dmiapi "github.com/kubeedge/api/apis/dmi/v1beta1"
	"github.com/kubeedge/mapper-framework/pkg/common"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestBuildPropertiesFromGrpcSetsOtelPushMethod(t *testing.T) {
	device := &dmiapi.Device{
		Spec: &dmiapi.DeviceSpec{
			Protocol: &dmiapi.ProtocolConfig{ProtocolName: "test"},
			Properties: []*dmiapi.DeviceProperty{
				{
					Name: "temperature",
					Visitors: &dmiapi.VisitorConfig{
						ConfigData: &dmiapi.CustomizedValue{Data: map[string]*anypb.Any{}},
					},
					PushMethod: &dmiapi.PushMethod{
						Otel: &dmiapi.PushMethodOTEL{EndpointURL: "http://localhost:4318/v1/metrics"},
					},
				},
			},
		},
	}

	properties := buildPropertiesFromGrpc(device)
	if len(properties) != 1 {
		t.Fatalf("expected one property, got %d", len(properties))
	}
	if properties[0].PushMethod.MethodName != common.PushMethodOTEL {
		t.Fatalf("expected OTEL push method name %q, got %q", common.PushMethodOTEL, properties[0].PushMethod.MethodName)
	}
	if properties[0].PushMethod.DBMethod.DBMethodName != "" {
		t.Fatalf("expected no database method name, got %q", properties[0].PushMethod.DBMethod.DBMethodName)
	}
}
