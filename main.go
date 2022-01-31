package main

import (
	"context"

	routes "api.linkfly.com/api/routes"
	"api.linkfly.com/configs"
	domain_linkly "api.linkfly.com/domain/linkly"
	"api.linkfly.com/infrastructure/repository"
	echo_customcontext "api.linkfly.com/pkg/echo/customContext"
	"api.linkfly.com/pkg/mongo"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs := configs.Load()
	containerRegister()

	echo_customcontext.InjectCustomContext(e)
	routes.RegisterRoute(e)

	e.Logger.Fatal(e.Start(":" + configs.Server.Port))
}

func containerRegister() {
	container.Singleton(func() configs.Configs {
		return configs.Load()
	})

	container.Singleton(func() mongo.Database {
		configs := configs.Load()
		return mongo.Connect(context.Background(), mongo.DatabaseConfig{
			ConnectionString: configs.Database.ConnectionString,
		})
	})

	container.Singleton(func() domain_linkly.ILinklyRepository {
		var database mongo.Database
		if err := container.Resolve(&database); err != nil {
			panic(err)
		}
		return &repository.LinklyRepository{Database: database}
	})

}
