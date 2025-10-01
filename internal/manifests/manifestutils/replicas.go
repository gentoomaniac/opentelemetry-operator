// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package manifestutils

import (
	"github.com/open-telemetry/opentelemetry-operator/apis/v1beta1"
)

func GetInitialReplicas(otelCol v1beta1.OpenTelemetryCollector) *int32 {
	if otelCol.Spec.Autoscaler != nil && otelCol.Spec.Autoscaler.MinReplicas != nil {
		if otelCol.Spec.Replicas != nil && *otelCol.Spec.Replicas > *otelCol.Spec.Autoscaler.MinReplicas {
			return otelCol.Spec.Replicas
		}
		return otelCol.Spec.Autoscaler.MinReplicas
	}

	return otelCol.Spec.Replicas
}
