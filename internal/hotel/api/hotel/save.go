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

		var req model.HotelInfo

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

		err = hotel.SaveHotel(ctx, &req)
		if err != nil {
			log.Error("failed to save hotel", sl.Err(err))
			render.JSON(w, r, response.Error("failed to save hotel"))

			return
		}
		log.Info("hotel successfully saved", slog.String("hotel", req.Name))

		render.JSON(w, r, SaveResponse{
			Response: response.OK(),
		})
	}
}
