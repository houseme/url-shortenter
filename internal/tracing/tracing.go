// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package tracing

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"github.com/houseme/url-shortenter/utility/env"
)

const (
	tracerHostnameTagKey = "hostname"
)

// InitJaeger initializes and registers jaeger to global TracerProvider.
//
// The output parameter `flush` is used for waiting exported trace spans to be uploaded,
// which is useful if your program is ending, and you do not want to lose recent spans.
func InitJaeger(serviceName, endpoint, version, environment, hostIP string) (*trace.TracerProvider, error) {
	var endpointOption jaeger.EndpointOption
	if strings.HasPrefix(endpoint, "http") {
		// HTTP.
		endpointOption = jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint))
	} else {
		// UDP.
		endpointOption = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(endpoint))
	}
	exp, err := jaeger.New(endpointOption)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName+"-"+environment),
			semconv.ServiceVersionKey.String(version),
			semconv.DeploymentEnvironmentKey.String(environment),
			semconv.HostNameKey.String(hostIP),
			attribute.String(tracerHostnameTagKey, hostIP),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

// InitTracer initializes and registers jaeger to global TracerProvider.
func InitTracer(ctx context.Context, serviceName string) {
	if appEnv, err := env.New(ctx); err != nil {
		g.Log().Fatal(ctx, err)
	} else {
		_, err = InitJaeger(serviceName, appEnv.JaegerEndpoint(ctx), appEnv.Version(ctx), appEnv.Environment(ctx), appEnv.HostIP(ctx))
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
	}
}
