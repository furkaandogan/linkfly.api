package routes

import (
	"context"
	"net/http"

	application_getlinkly "api.linkfly.com/application/getLinkly"
	echo_customcontext "api.linkfly.com/pkg/echo/customContext"
	"github.com/labstack/echo/v4"
)

func GetLinklyRoute(e *echo.Echo) {
	e.GET("/:hash", func(c echo.Context) error {
		customContext := c.(*echo_customcontext.CustomContext)
		hash := c.Param("hash")

		command := application_getlinkly.GetLinklyCommand{
			IP:        customContext.GetIP(),
			Hash:      hash,
			Useragent: c.Request().UserAgent(),
		}
		result, _ := command.Execute(context.Background())
		return c.JSON(http.StatusOK, result)
	})
}
