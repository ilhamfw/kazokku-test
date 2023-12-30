package app

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
    // Inisialisasi Echo
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routing
    // ...
    // Tambahkan routing sesuai dengan kebutuhan aplikasi Anda

    return e
}
