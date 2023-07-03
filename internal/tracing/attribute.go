// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package tracing

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
)

const telemetrySDKName = "opentelemetry"

// SetAttributes .set tracing attributes
func SetAttributes(r *ghttp.Request, span *gtrace.Span) {
	span.SetAttributes(semconv.HTTPURL(r.URL.Path))
	span.SetAttributes(semconv.HTTPMethod(r.Method))
	span.SetAttributes(semconv.NetHostName(r.GetHost()))
	span.SetAttributes(semconv.HTTPScheme(r.Proto))
	span.SetAttributes(semconv.HTTPStatusCode(r.Response.Status))
	span.SetAttributes(semconv.UserAgentOriginal(r.UserAgent()))
}

// CommonEventOption .
func CommonEventOption(_ context.Context, namespace string) trace.SpanStartEventOption {
	return trace.WithAttributes(
		semconv.ServiceNamespace(namespace),
		semconv.TelemetrySDKName(telemetrySDKName),
		semconv.TelemetrySDKVersion("1.0.0"),
		semconv.TelemetryAutoVersion("1.0.0"),
		semconv.TelemetrySDKLanguageGo,
	)
}
