package routes

import (
	create_route "api.linkfly.com/api/routes/createLinkly"
	createWithFocusElement_route "api.linkfly.com/api/routes/createLinklyWithFocusElement"
	get_route "api.linkfly.com/api/routes/getLinkly"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	create_route.CreateLinklyRoute(e)
	get_route.GetLinklyRoute(e)
	createWithFocusElement_route.CreateLinklyWithFocusElementRoute(e)
}
