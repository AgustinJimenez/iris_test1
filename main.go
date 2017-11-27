// file: main.go

package main

import (
	"./bootstrap"
	"./web/middlewares/identity"
	"./web/routes"
)

func main() {
	app := bootstrap.New("Iris Framework Test 1", "agus.jmenez.caba@gmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8000")
}
