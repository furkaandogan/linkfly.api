package main

import (
	"context"

	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	routes "linkfly.api/api/routes"
	domain_linkly "linkfly.api/domain/linkly"
	"linkfly.api/infrastructure/repository"
	echo_customcontext "linkfly.api/pkg/echo/customContext"
	"linkfly.api/pkg/mongo"
)

func main() {
	e := echo.New()
	containerRegister()

	echo_customcontext.InjectCustomContext(e)
	routes.RegisterRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func containerRegister() {

	container.Singleton(func() mongo.Database {
		return mongo.Connect(context.Background(), mongo.DatabaseConfig{
			ConnectionString: "mongodb+srv://sa:kC8Yf*HEX3v82VwH@cluster0.gghum.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
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
