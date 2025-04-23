// filepath: /Volumes/Storage/Document/learning/coding-test/ecommerce-api/cmd/main.go
package main

import (
	"ecommerce-api/configs"
	"ecommerce-api/internal/infrastructure/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    // Init config
    db := configs.InitDB()
    configs.MigrateDB(db)
    e := echo.New()

    // Global middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routing
    router.Init(e, db)

    // Start server
    port := configs.GetEnv("APP_PORT", "8080")
    e.Logger.Fatal(e.Start(":" + port))
}