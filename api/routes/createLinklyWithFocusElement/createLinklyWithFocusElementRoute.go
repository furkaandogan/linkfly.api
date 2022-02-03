package routes

import (
	"context"
	"net/http"

	application_createlinkly "api.linkfly.com/application/createLinklyWithFocusElement"
	echo_customcontext "api.linkfly.com/pkg/echo/customContext"
	"github.com/labstack/echo/v4"
)

type CreateLinklyWithFocusElementRequest struct {
	Link  string `json:"link"`
	XPath string `json:xpath`
}

type CreateLinklyWithFocusElementResponse struct {
}

func CreateLinklyWithFocusElementRoute(e *echo.Echo) {
	e.POST("/focus-element", func(c echo.Context) error {
		customContext := c.(*echo_customcontext.CustomContext)
		requestBody := CreateLinklyWithFocusElementRequest{}

		if err := (&echo.DefaultBinder{}).BindBody(customContext, &requestBody); err != nil {
			return err
		}

		command := application_createlinkly.CreateLinklyWithFocusElementCommand{
			IP:        customContext.GetIP(),
			Url:       requestBody.Link,
			XPath:     requestBody.XPath,
			Useragent: c.Request().UserAgent(),
		}
		result, _ := command.Execute(context.Background())
		return c.JSON(http.StatusOK, result)
	})
}
