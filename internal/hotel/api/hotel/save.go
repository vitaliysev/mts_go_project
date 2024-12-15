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

type SaveHotelRequest struct {
	Info         model.HotelInfo `json:"info"`
	Access_token string          `json:"access_token"`
}

type SaveResponse struct {
	response.Response
}

func NewSave(ctx context.Context, log *slog.Logger, hotel *Implementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "hotel.handlers.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req SaveHotelRequest

		err := render.DecodeJSON(r.Body, &req)
		fmt.Println(req.Info)
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

		username, errReq := cl.Check(ctx_curr, &descAccess.CheckRequest{
			EndpointAddress: "/saveHotel",
		})
		if errReq != nil {
			log.Error(errReq.Error())
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
		log.Info(username.GetUsername())
		err = hotel.SaveHotel(ctx, &req.Info, username.GetUsername())
		if err != nil {
			log.Error("failed to save hotel", sl.Err(err))
			render.JSON(w, r, response.Error("failed to save hotel"))

			return
		}
		log.Info("hotel successfully saved", slog.String("hotel", req.Info.Name))

		render.JSON(w, r, SaveResponse{
			Response: response.OK(),
		})
	}
}
