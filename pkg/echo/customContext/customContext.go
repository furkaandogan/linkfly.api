package echo_customcontext

import "github.com/labstack/echo/v4"

type CustomContext struct {
	echo.Context
}

func (context *CustomContext) GetIP() string {
	return context.Context.RealIP()
}

func InjectCustomContext(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&CustomContext{c})
		}
	})
}
