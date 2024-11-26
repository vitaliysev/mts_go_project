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

type UpdateResponse struct {
	response.Response
}

func NewUpdate(ctx context.Context, log *slog.Logger, hotel *Implementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "hotel.handlers.update.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req model.Hotel
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

		err = hotel.UpdateHotel(ctx, req)
		if err != nil {
			log.Error("failed to update hotel", sl.Err(err))
			render.JSON(w, r, response.Error("failed to update hotel"))
			return
		}
		render.JSON(w, r, UpdateResponse{
			Response: response.OK(),
		})
	}
}
