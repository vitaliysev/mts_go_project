package interceptor

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc"
)

const traceIDKey = "x-trace-id"

func ServerTracingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// Embedding span config into the context

	ctx, span := tracing.Tracer.Tracer("Auth-service").Start(ctx, info.FullMethod)
	defer span.End()
	res, err := handler(ctx, req)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	} else {
		// Ответ может быть большим, поэтому не стоит добавлять его в теги
		// Здесь это лишь пример, как можно добавить ответ в тег
	}

	return res, err
}
