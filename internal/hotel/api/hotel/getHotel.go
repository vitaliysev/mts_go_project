package hotel

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/lib/api/response"
	"github.com/vitaliysev/mts_go_project/internal/lib/logger/sl"
	descAccess "github.com/vitaliysev/mts_go_project/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log/slog"
	"net/http"
)

type getHotelRequest struct {
	ID           int64  `json:"id" validate:"required"`
	Access_token string `json:"access_token"`
}
type getHotelResponse struct {
	response.Response
	model.Hotel
}

func NewGetHotel(ctx context.Context, log *slog.Logger, hotel *Implementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "hotel.handlers.get.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req getHotelRequest
		err := render.DecodeJSON(r.Body, &req)
		accessToken := req.Access_token

		ctx_curr := context.Background()
		md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
		ctx_curr = metadata.NewOutgoingContext(ctx_curr, md)

		conn, err := grpc.Dial(
			"localhost:50055",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Error("failed to dial GRPC client: %v", err)
		}

		cl := descAccess.NewAccessV1Client(conn)

		_, err = cl.Check(ctx_curr, &descAccess.CheckRequest{
			EndpointAddress: "/getHotel",
		})
		if err != nil {
			log.Error(err.Error())
		}

		fmt.Println("Access granted")

		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, response.Error("failed to decode request"))
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("invalid request", sl.Err(err))
			render.JSON(w, r, response.ValidationError(validateErr))
			return
		}
		data, err := hotel.GetHotel(ctx, req.ID)
		if err != nil {
			log.Error("failed to get hotel", sl.Err(err))
			render.JSON(w, r, response.Error("failed to get hotel"))
			return
		}
		render.JSON(w, r, getHotelResponse{
			Response: response.OK(),
			Hotel:    *data,
		})
	}
}
