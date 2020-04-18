package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello World!")
}
