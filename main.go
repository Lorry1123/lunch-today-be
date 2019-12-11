package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	restaurants "lunch-today-be/restaurants"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	// 样例代码中的 logger 部分，保留


	app.Get("/random", func(ctx iris.Context) {
		ctx.JSON(restaurants.GetRandomRestuarants())
	})

	app.Get("/date", func(ctx iris.Context) {
		ctx.JSON(restaurants.GetRestuarantsByDate())
	})

	app.Run(iris.Addr(":8010"), iris.WithoutServerError(iris.ErrServerClosed))
}