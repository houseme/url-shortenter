package tracing

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

const telemetrySDKName = "opentelemetry"

// SetAttributes .set tracing attributes
func SetAttributes(r *ghttp.Request, span *gtrace.Span) {
	span.SetAttributes(semconv.HTTPURLKey.String(r.URL.Path))
	span.SetAttributes(semconv.HTTPMethodKey.String(r.Method))
	span.SetAttributes(semconv.HTTPHostKey.String(r.GetHost()))
	span.SetAttributes(semconv.HTTPSchemeKey.String(r.Proto))
	span.SetAttributes(semconv.HTTPStatusCodeKey.String(gconv.String(r.Response.Status)))
	span.SetAttributes(semconv.HTTPUserAgentKey.String(r.UserAgent()))
}

// CommonEventOption .
func CommonEventOption(_ context.Context, namespace string) trace.SpanStartEventOption {
	return trace.WithAttributes(
		semconv.ServiceNamespaceKey.String(namespace),
		semconv.TelemetrySDKNameKey.String(telemetrySDKName),
		semconv.TelemetrySDKVersionKey.String("1.0.0"),
		semconv.TelemetryAutoVersionKey.String("1.0.0"),
		semconv.TelemetrySDKLanguageGo,
	)
}
