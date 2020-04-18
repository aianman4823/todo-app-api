package main

import (
	"github.com/aianman4823/todo-app-api/handler"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.hello)

	e.Logger.Fatal(e.Start(":1323"))
}
