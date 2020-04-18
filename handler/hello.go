package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// Hello: function of Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello World!")
}
