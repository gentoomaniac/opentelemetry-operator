package manifestutils

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-operator/apis/v1beta1"
	"github.com/stretchr/testify/assert"
)

func TestGetInitialReplicas(t *testing.T) {
	// prepare
	minReplicas := int32(1)
	desiredReplicas := int32(1)
	otelcol := v1beta1.OpenTelemetryCollector{
		Spec: v1beta1.OpenTelemetryCollectorSpec{
			Autoscaler: &v1beta1.AutoscalerSpec{
				MinReplicas: &minReplicas,
			},
		},
	}

	// test
	assert.Equal(t, *GetInitialReplicas(otelcol), minReplicas)

	otelcol.Spec.Replicas = &desiredReplicas
	assert.Equal(t, *GetInitialReplicas(otelcol), minReplicas)

	desiredReplicas = 2
	assert.Equal(t, *GetInitialReplicas(otelcol), desiredReplicas)
}
