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
	"github.com/gogf/gf/v2/net/gipv4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"

	"github.com/houseme/url-shortenter/utility/env"
)

const (
	tracerHostnameTagKey = "hostname"
	traceJaeger          = "jaeger"
	traceOtlpGRPC        = "otlpgrpc"
)

// InitJaeger initializes and registers jaeger to global TracerProvider.
//
// The output parameter `flush` is used for waiting exported trace spans to be uploaded,
// which is useful if your program is ending, and you do not want to lose recent spans.
func InitJaeger(serviceName, endpoint, version, environment string) (*trace.TracerProvider, error) {
	var (
		intranetIPArray, err = gipv4.GetIntranetIpArray()
		hostIP               = "NoHostIpFound"
	)

	if err != nil {
		return nil, err
	}

	if len(intranetIPArray) == 0 {
		if intranetIPArray, err = gipv4.GetIpArray(); err != nil {
			return nil, err
		}
	}
	if len(intranetIPArray) > 0 {
		hostIP = intranetIPArray[0]
	}

	var (
		endpointOption jaeger.EndpointOption
		exp            *jaeger.Exporter
	)
	if strings.HasPrefix(endpoint, "http") {
		// HTTP.
		endpointOption = jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint))
	} else {
		// UDP.
		endpointOption = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(endpoint))
	}
	if exp, err = jaeger.New(endpointOption); err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName+"-"+environment),
			semconv.ServiceVersion(version),
			semconv.DeploymentEnvironment(environment),
			semconv.HostName(hostIP),
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
		if appEnv.TraceType(ctx) == traceJaeger {
			_, err = InitJaeger(serviceName, appEnv.JaegerEndpoint(ctx), appEnv.Version(ctx), appEnv.Environment(ctx))
		}
		if appEnv.TraceType(ctx) == traceOtlpGRPC {
			_, err = Init(serviceName, appEnv.Endpoint(ctx), appEnv.TraceToken(ctx), appEnv.Version(ctx), appEnv.Environment(ctx))
		}
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
	}
}

// Init initializes and registers `otlpgrpc` to global TracerProvider.
//
// The output parameter `Shutdown` is used for waiting exported tracing spans to be uploaded,
// which is useful if your program is ending, and you do not want to lose recent spans.
func Init(serviceName, endpoint, traceToken, version, environment string) (*trace.TracerProvider, error) {
	// Try retrieving host ip for tracing info.
	var (
		intranetIPArray, err = gipv4.GetIntranetIpArray()
		hostIP               = "NoHostIpFound"
	)

	if err != nil {
		return nil, err
	}

	if len(intranetIPArray) == 0 {
		if intranetIPArray, err = gipv4.GetIpArray(); err != nil {
			return nil, err
		}
	}
	if len(intranetIPArray) > 0 {
		hostIP = intranetIPArray[0]
	}

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint), // Replace the otel Agent Addr with the access point obtained in the prerequisite。
		otlptracegrpc.WithHeaders(map[string]string{"Authentication": traceToken}),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	ctx := context.Background()
	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// The name of the service displayed on the traceback end。
			semconv.ServiceNameKey.String(serviceName+"-"+environment),
			semconv.ServiceVersionKey.String(version),
			semconv.DeploymentEnvironmentKey.String(environment),
			semconv.HostNameKey.String(hostIP),
			attribute.String(tracerHostnameTagKey, hostIP),
		),
	)
	if err != nil {
		return nil, err
	}

	bsp := trace.NewBatchSpanProcessor(traceExp)
	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithSpanProcessor(bsp),
	)

	// Set the global propagator to traceContext (not set by default).
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)

	return tracerProvider, nil
}
