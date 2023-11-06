// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package tracing

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"google.golang.org/grpc"

	"github.com/houseme/url-shortenter/utility/env"
)

const tracerHostnameTagKey = "hostname"

// InitTracer initializes and registers jaeger to global TracerProvider.
func InitTracer(ctx context.Context) (shutdown func()) {
	if appEnv, err := env.New(ctx); err != nil {
		g.Log().Fatal(ctx, "InitTracer env new failed, error:", err)
	} else {
		if shutdown, err = Init(appEnv.ApplicationService(), appEnv.Endpoint(ctx), appEnv.TraceToken(ctx), appEnv.Version(ctx), appEnv.Environment(ctx)); err != nil {
			g.Log().Fatal(ctx, "InitTracer init otlp grpc failed err:", err)
		}
	}
	return
}

// Init initializes and registers `otlpgrpc` to global TracerProvider.
//
// The output parameter `Shutdown` is used for waiting exported tracing spans to be uploaded,
// which is useful if your program is ending, and you do not want to lose recent spans.
func Init(serviceName, endpoint, traceToken, version, environment string) (func(), error) {
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

	return func() {
		// Shutdown flushes any remaining spans and shuts down the exporter.
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err = traceExp.Shutdown(ctx); err != nil {
			g.Log().Errorf(ctx, "Shutdown traceExp failed err:%+v", err)
			otel.Handle(err)
		}
		g.Log().Debug(ctx, "Shutdown traceExp success")
	}, nil
}

// InitOtlpHTTP initializes and registers `otlphttp` to global TracerProvider.
//
// The output parameter `Shutdown` is used for waiting exported trace spans to be uploaded,
// which is useful if your program is ending, and you do not want to lose recent spans.
func InitOtlpHTTP(serviceName, endpoint, path, version, environment string) (func(), error) {
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

	traceClientHTTP := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithURLPath(path),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithCompression(1),
	)

	ctx := context.Background()
	traceExp, err := otlptrace.New(ctx, traceClientHTTP)
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
			semconv.ServiceName(serviceName+"-"+environment),
			semconv.ServiceVersion(version),
			semconv.DeploymentEnvironment(environment),
			semconv.HostName(hostIP),
			attribute.String(tracerHostnameTagKey, hostIP),
		),
	)

	bsp := trace.NewBatchSpanProcessor(traceExp)
	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithSpanProcessor(bsp),
	)

	// Set the global propagator to traceContext (not set by default).
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err = traceExp.Shutdown(ctx); err != nil {
			g.Log().Errorf(ctx, "Shutdown traceExp failed err:%+v", err)
			otel.Handle(err)
		}
		g.Log().Debug(ctx, "Shutdown traceExp success")
	}, nil
}
