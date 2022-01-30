package routes

import (
	"github.com/labstack/echo/v4"
	create_route "linkfly.api/api/routes/createLinkly"
	get_route "linkfly.api/api/routes/getLinkly"
)

func RegisterRoute(e *echo.Echo) {
	create_route.CreateLinklyRoute(e)
	get_route.GetLinklyRoute(e)
}
