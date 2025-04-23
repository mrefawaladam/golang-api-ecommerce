package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
    api := e.Group("/api")
    ProductRouter(api, db)
    UserRouter(api, db)
    CartRouter(api, db)
    OrderRouter(api, db)
    RegisterAccountRoutes(api)
}