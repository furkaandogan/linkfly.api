package main

import (
	"context"

	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	routes "linkfly.api/api/routes"
	"linkfly.api/configs"
	domain_linkly "linkfly.api/domain/linkly"
	"linkfly.api/infrastructure/repository"
	echo_customcontext "linkfly.api/pkg/echo/customContext"
	"linkfly.api/pkg/mongo"
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
		var configs configs.Configs
		if err := container.Resolve(&configs); err != nil {
			panic(err)
		}
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
