package hotel

import (
	"context"
	"fmt"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	metric "github.com/vitaliysev/mts_go_project/internal/hotel/metrics"
	"github.com/vitaliysev/mts_go_project/internal/lib/api/response"
	"github.com/vitaliysev/mts_go_project/internal/lib/logger"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	descAccess "github.com/vitaliysev/mts_go_project/pkg/access_v1"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

type GetHotelRequest struct {
	ID           int64  `json:"id" validate:"required"`
	Access_token string `json:"access_token"`
}
type GetHotelResponse struct {
	response.Response
	model.Hotel
}

// NewGetHotel Получение одного отеля.
// @Summary Получение одного отеля
// @Description Получение одного отеля используя HTTP API.
// @Tags Hotel
// @Accept json
// @Produce json
// @Param hotelBody body GetHotelRequest true "Hotel Data"
// @Success 200 {object} GetHotelResponse "Hotel got successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /getHotel [post]
func NewGetHotel(ctx context.Context, hotel *Implementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "hotel.GetHotel"
		ctx, span := tracing.Tracer.Tracer("Hotel-service").Start(ctx, op)
		defer span.End()
		traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
		ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)
		timeStart := time.Now()
		log := logger.With(
			zap.String("op", op),
		)
		metric.IncRequestCounter()

		var req GetHotelRequest
		err := render.DecodeJSON(r.Body, &req)
		log.Info("Decoding request body...")
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			diffTime := time.Since(timeStart)
			log.Error("failed to decode request body", zap.Error(err))
			metric.IncResponseCounter("error", op)
			metric.HistogramResponseTimeObserve("error", diffTime.Seconds())
			render.JSON(w, r, response.Error("failed to decode request"))
			return
		}
		log.Info("request body decoded", zap.Any("request", req))
		accessToken := req.Access_token

		ctx_curr := context.Background()
		md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
		ctx_curr = metadata.NewOutgoingContext(ctx_curr, md)
		ctx_curr = metadata.AppendToOutgoingContext(ctx_curr, "x-trace-id", traceId)

		conn, err := grpc.Dial(
			"localhost:50055",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Error("failed to dial GRPC client: %v", zap.Error(err))
		}

		cl := descAccess.NewAccessV1Client(conn)

		_, err = cl.Check(ctx_curr, &descAccess.CheckRequest{
			EndpointAddress: "/getHotel",
		})
		if err != nil {
			log.Error(err.Error())
		}

		fmt.Println("Access granted")

		log.Info("validating request...")

		if err := validator.New().Struct(req); err != nil {
			span.SetStatus(codes.Error, err.Error())
			diffTime := time.Since(timeStart)
			validateErr := err.(validator.ValidationErrors)
			log.Error("invalid request", zap.Error(err))
			metric.IncResponseCounter("error", op)
			metric.HistogramResponseTimeObserve("error", diffTime.Seconds())
			render.JSON(w, r, response.ValidationError(validateErr))
			return
		}
		log.Info("getting hotel...")
		data, err := hotel.GetHotel(ctx, req.ID)

		if err != nil {
			diffTime := time.Since(timeStart)
			log.Error("failed to get hotel", zap.Error(err))
			metric.IncResponseCounter("error", op)
			metric.HistogramResponseTimeObserve("error", diffTime.Seconds())
			render.JSON(w, r, response.Error("failed to get hotel"))
			return
		}

		diffTime := time.Since(timeStart)
		metric.HistogramResponseTimeObserve("success", diffTime.Seconds())
		metric.IncResponseCounter("success", op)
		render.JSON(w, r, GetHotelResponse{
			Response: response.OK(),
			Hotel:    *data,
		})
	}
}
