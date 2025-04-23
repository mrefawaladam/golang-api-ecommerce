package middleware

import (
	"ecommerce-api/internal/util"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
    return echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte(util.GetJWTSecret()),
        ErrorHandler: func(c echo.Context, err error) error {
            return c.JSON(http.StatusUnauthorized, map[string]string{
                "message": "Token tidak valid atau sudah kadaluarsa",
            })
        },
    })
}
