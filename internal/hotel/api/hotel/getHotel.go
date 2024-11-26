package hotel

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/lib/api/response"
	"github.com/vitaliysev/mts_go_project/internal/lib/logger/sl"
	"log/slog"
	"net/http"
)

type getHotelRequest struct {
	ID int64 `json:"id" validate:"required"`
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
