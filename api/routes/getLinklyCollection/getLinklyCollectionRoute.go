package routes

import (
	"context"
	"net/http"

	application_getlinklyCollection "api.linkfly.com/application/getLinklyCollection"
	echo_customcontext "api.linkfly.com/pkg/echo/customContext"
	"github.com/labstack/echo/v4"
)

func GetLinklyCollectionRoute(e *echo.Echo) {
	e.GET("/linkly/collection", func(c echo.Context) error {
		customContext := c.(*echo_customcontext.CustomContext)
		hash := c.Param("hash")

		command := application_getlinklyCollection.GetLinklyCollectionCommand{
			IP:        customContext.GetIP(),
			Hash:      hash,
			Useragent: c.Request().UserAgent(),
		}
		result, _ := command.Execute(context.Background())
		return c.JSON(http.StatusOK, result)
	})
}
