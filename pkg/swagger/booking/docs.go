// Package booking Code generated by swaggo/swag. DO NOT EDIT
package booking

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/booking/v1/create": {
            "post": {
                "description": "Создание нового бронирования используя HTTP API.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "Создание нового бронирования",
                "parameters": [
                    {
                        "description": "Booking Data",
                        "name": "bookingBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.CreateBookingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Booking created successfully",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.CreateBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/booking/v1/getrefr": {
            "patch": {
                "description": "Обновление RefreshToken используя HTTP API.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "Обновление RefreshToken",
                "parameters": [
                    {
                        "description": "Auth Data",
                        "name": "bookingBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.GetRefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Refresh token got succesfully",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.GetRefreshTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/booking/v1/listCl": {
            "post": {
                "description": "Получение бронирований отелей клиента используя HTTP API.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "Получение бронирований отелей клиента",
                "parameters": [
                    {
                        "description": "Booking Data",
                        "name": "bookingBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.GetBookingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bookings listed successfully",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.GetBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/booking/v1/listHo": {
            "post": {
                "description": "Получение бронирований отелей владельца используя HTTP API.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "Получение бронирований отелей владельца",
                "parameters": [
                    {
                        "description": "Booking Data",
                        "name": "bookingBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.GetBookingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bookings listed successfully",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.GetBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/booking/v1/login": {
            "post": {
                "description": "Вход используя HTTP API.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "Обновление RefreshToken",
                "parameters": [
                    {
                        "description": "Access Data",
                        "name": "bookingBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.LoginClientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Access if succesful",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.LoginClientResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/booking/v1/signin": {
            "post": {
                "description": "Регистрация используя HTTP API.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "Регистрация",
                "parameters": [
                    {
                        "description": "Auth Data",
                        "name": "bookingBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.SigninClientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Authentication is succesful",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.SigninClientResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_booking_app.ErrorResponse"
                        }
                    }
                }
            }
        },
    "definitions": {
        "github_com_vitaliysev_mts_go_project_internal_booking_model.Book": {
            "description": "Book represents a book in the booking system",
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "$ref": "#/definitions/github_com_vitaliysev_mts_go_project_internal_booking_model.BookInfo"
                },
                "update_time": {
                    "$ref": "#/definitions/sql.NullTime"
                }
            }
        },
        "github_com_vitaliysev_mts_go_project_internal_booking_model.BookInfo": {
            "type": "object",
            "properties": {
                "hotel_id": {
                    "type": "integer"
                },
                "period_use": {
                    "type": "string"
                }
            }
        },
        "hotel.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Invalid request"
                }
            }
        },
        "hotel.GetHotelRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "hotel.GetHotelResponse": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "error": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "$ref": "#/definitions/model.HotelInfo"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "hotel.GetHotelsRequest": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "hotel.GetHotelsResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "hotels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Hotel"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "hotel.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "hotel.SaveHotelRequest": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "info": {
                    "$ref": "#/definitions/model.HotelInfo"
                }
            }
        },
        "hotel.UpdateRequest": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "info": {
                    "$ref": "#/definitions/model.Hotel"
                }
            }
        },
        "hotel.updateResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "internal_booking_app.CreateBookingRequest": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "info": {
                    "$ref": "#/definitions/github_com_vitaliysev_mts_go_project_internal_booking_model.BookInfo"
                }
            }
        },
        "internal_booking_app.CreateBookingResponse": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "period": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "internal_booking_app.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Invalid request"
                }
            }
        },
        "internal_booking_app.GetBookingRequest": {
            "description": "GetBookingRequest contains a info for get",
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "internal_booking_app.GetBookingResponse": {
            "description": "GetBookingResponse contains a list of booking information.",
            "type": "object",
            "properties": {
                "info": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_vitaliysev_mts_go_project_internal_booking_model.Book"
                    }
                }
            }
        },
        "internal_booking_app.GetRefreshTokenRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "internal_booking_app.GetRefreshTokenResponse": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "internal_booking_app.LoginClientRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "internal_booking_app.LoginClientResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "internal_booking_app.SigninClientRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "internal_booking_app.SigninClientResponse": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "model.Hotel": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "info": {
                    "$ref": "#/definitions/model.HotelInfo"
                }
            }
        },
        "model.HotelInfo": {
            "type": "object",
            "required": [
                "location",
                "name",
                "price"
            ],
            "properties": {
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "sql.NullTime": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Booking API",
	Description:      "GetBookingResponse contains a list of booking information.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
