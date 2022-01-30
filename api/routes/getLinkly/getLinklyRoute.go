package routes

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	application_getlinkly "linkfly.api/application/getLinkly"
	echo_customcontext "linkfly.api/pkg/echo/customContext"
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
