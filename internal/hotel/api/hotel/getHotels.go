package hotel

import (
	"context"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/api/hotel/model"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/lib/api/response"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type getHotelsRequest struct {
}
type getHotelsResponse struct {
	response.Response
	Hotels []model.Hotel `json:"hotels"`
}

func NewGetHotels(ctx context.Context, log *slog.Logger, hotel *Implementation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "hotel.handlers.get.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req getHotelsRequest
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

		hotels, err := hotel.GetHotels(ctx)
		if err != nil {
			log.Error("failed to get hotels", sl.Err(err))
			render.JSON(w, r, response.Error("failed to get hotels"))
			return
		}
		render.JSON(w, r, getHotelsResponse{
			Response: response.OK(),
			Hotels:   hotels,
		})
	}
}
