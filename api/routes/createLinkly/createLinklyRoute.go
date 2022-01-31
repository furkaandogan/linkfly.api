package routes

import (
	"context"
	"net/http"

	application_createlinkly "api.linkfly.com/application/createLinkly"
	echo_customcontext "api.linkfly.com/pkg/echo/customContext"
	"github.com/labstack/echo/v4"
)

type CreateLinklyRequest struct {
	Link string `json:"link"`
}

type CreateLinklyResponse struct {
}

func CreateLinklyRoute(e *echo.Echo) {
	e.POST("/", func(c echo.Context) error {
		customContext := c.(*echo_customcontext.CustomContext)
		requestBody := CreateLinklyRequest{}

		if err := (&echo.DefaultBinder{}).BindBody(customContext, &requestBody); err != nil {
			return err
		}

		command := application_createlinkly.CreateLinklyCommand{
			IP:        customContext.GetIP(),
			Url:       requestBody.Link,
			Useragent: c.Request().UserAgent(),
		}
		result, _ := command.Execute(context.Background())
		return c.JSON(http.StatusOK, result)
	})
}
