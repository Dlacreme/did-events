package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterMiddlewares(e *echo.Echo) {

}

func RegisterEndpoints(e *echo.Echo) {
	e.GET("/", health)
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "<Events>: Im alive.")
}
