package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Check() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok"})
	}
}
