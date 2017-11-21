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

/*
app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	// Load the template files.

	tmpl := iris.HTML("./web/views", ".html").
		Layout("shared/layout.html").
		Reload(true)

	app.RegisterView(tmpl)

	app.StaticWeb("/public", "./web/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	// Create our repositories and services.
	db, err := datasource.LoadUsers(datasource.Memory)
	if err != nil {
		app.Logger().Fatalf("error while loading the users: %v", err)
		return
	}
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)

	app.Controller("/", new(controllers.HomeController))
	app.Controller("/login", new(controllers.LoginController))
	app.Controller("/admin", new(controllers.AdminController), middleware.BasicAuth)
	// Register our controllers.
	app.Controller("/users", new(controllers.UsersController),
		// Add the basic authentication(admin:password) middleware
		// for the /users based requests.
		middleware.BasicAuth,
		// Bind the "userService" to the UserController's Service (interface) field.
		userService,
	)

	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookiename",
		Expires: 24 * time.Hour,
	})

	app.Controller("/user", new(controllers.UserController), userService, sessManager)

	// Start the web server at localhost:8080
	// http://localhost:8080/hello
	// http://localhost:8080/hello/iris
	// http://localhost:8080/users/1
	app.Run(
		iris.Addr("localhost:8000"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, // enables faster json serialization and more
	)
*/
