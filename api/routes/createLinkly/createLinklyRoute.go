package routes

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	application_createlinkly "linkfly.api/application/createLinkly"
	echo_customcontext "linkfly.api/pkg/echo/customContext"
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
