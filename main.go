package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
    "github.com/kataras/iris/mvc"
)

var basicAuth = basicauth.New(basicauth.Config{
    Users: map[string]string{
        "admin": "password",
    },
})

func main() {
	app := iris.New()
    app.Use(basicAuth)
    app.Controller("/movies", new(MoviesController))
    app.Run(iris.Addr(":8080"))
}
