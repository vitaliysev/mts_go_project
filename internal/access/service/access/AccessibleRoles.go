package access

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
)

var accessibleRoles map[string]string

func (s *serv) AccessibleRoles(ctx context.Context) (map[string]string, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]string)

		accessibleRoles[model.CreateBookingPath] = "Client"
		accessibleRoles[model.GetHotelsPath] = "Client"
		accessibleRoles[model.ListBookingPathClient] = "Client"
		accessibleRoles[model.GetHotelPath] = "Hotelier"
		accessibleRoles[model.ListBookingPathHotel] = "Hotelier"
		accessibleRoles[model.SaveHotelPath] = "Hotelier"
		accessibleRoles[model.UpdateHotelPath] = "Hotelier"
	}

	return accessibleRoles, nil
}
