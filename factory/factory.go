package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	udata "loundryapp/features/users/data"
	udeliv "loundryapp/features/users/delivery"
	ucase "loundryapp/features/users/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	validator := validator.New()

	userData := udata.New(db)
	userCase := ucase.New(userData, validator)
	userHandler := udeliv.New(userCase, userData)
	udeliv.RouteUser(e, userHandler)

}
