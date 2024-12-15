package model

import "github.com/dgrijalva/jwt-go"

const (
	CreateBookingPath     = "/booking/v1/create"
	ListBookingPathClient = "/booking/v1/listCl"
	ListBookingPathHotel  = "/booking/v1/listHo"
	SaveHotelPath         = "/saveHotel"
	UpdateHotelPath       = "/updateHotel"
	GetHotelPath          = "/getHotel"
	GetHotelsPath         = "/getHotels"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
