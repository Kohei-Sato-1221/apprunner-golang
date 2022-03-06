package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloAppRunner)

	e.Logger.Fatal(e.Start(":8080"))
}

func helloAppRunner(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, This is AWS App Runner!!")
}
