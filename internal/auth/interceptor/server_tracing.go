package interceptor

import (
	"context"

	"github.com/vitaliysev/mts_go_project/internal/tracing"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const traceIDKey = "x-trace-id"

func ServerTracingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	traceIdString := md["x-trace-id"][0]
	traceId, err := trace.TraceIDFromHex(traceIdString)
	if err != nil {
		return nil, err
	}
	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
	})
	ctx = trace.ContextWithSpanContext(ctx, spanContext)
	ctx, span := tracing.Tracer.Tracer("Auth-service").Start(ctx, info.FullMethod)
	defer span.End()
	res, err := handler(ctx, req)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
