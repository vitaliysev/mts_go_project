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

type SaveHotelRequest struct {
	Info         model.HotelInfo `json:"info"`
	Access_token string          `json:"access_token"`
}

type SaveResponse struct {
	response.Response
}

// ErrorResponse represents a standard error response.
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Invalid request"`
}

// Save Создание нового отеля.
// @Summary Создание нового отеля
// @SecurityApiKeyAuth
// @Description Создание нового отеля используя HTTP API.
// @Tags Hotel
// @Accept json
// @Produce json
// @Param hotelBody body SaveHotelRequest true "Hotel Data"
// @Success 200 {object} Response "Hotel saved successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /saveHotel [post]
func NewSave(ctx context.Context, hotel *Implementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "hotel.Save"
		timeStart := time.Now()
		log := logger.With(
			zap.String("op", op),
		)

		ctx, span := tracing.Tracer.Tracer("Hotel-service").Start(ctx, op)
		defer span.End()
		traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
		fmt.Println(traceId)
		ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

		metric.IncRequestCounter()

		var req SaveHotelRequest

		err := render.DecodeJSON(r.Body, &req)
		log.Info("decoding request body...")
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

		username, errReq := cl.Check(ctx_curr, &descAccess.CheckRequest{
			EndpointAddress: "/saveHotel",
		})
		if errReq != nil {
			log.Error(errReq.Error())
		}

		fmt.Println("Access granted")

		log.Info("validating request body...")

		if err := validator.New().Struct(req); err != nil {
			span.SetStatus(codes.Error, err.Error())
			diffTime := time.Since(timeStart)
			validateErr := err.(validator.ValidationErrors)
			metric.IncResponseCounter("error", op)
			metric.HistogramResponseTimeObserve("error", diffTime.Seconds())
			log.Error("invalid request", zap.Error(err))
			render.JSON(w, r, response.ValidationError(validateErr))
			return
		}
		log.Info(username.GetUsername())
		err = hotel.SaveHotel(ctx, &req.Info, username.GetUsername())
		if err != nil {
			diffTime := time.Since(timeStart)
			log.Error("failed to save hotel", zap.Error(err))
			metric.IncResponseCounter("error", op)
			metric.HistogramResponseTimeObserve("error", diffTime.Seconds())
			render.JSON(w, r, response.Error("failed to save hotel"))
			return
		}
		log.Info("hotel successfully saved", zap.String("hotel", req.Info.Name))
		diffTime := time.Since(timeStart)

		metric.IncResponseCounter("success", op)
		metric.HistogramResponseTimeObserve("success", diffTime.Seconds())
		render.JSON(w, r, SaveResponse{
			Response: response.OK(),
		})
	}
}
