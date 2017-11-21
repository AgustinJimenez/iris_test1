package routes

import (
	"../controllers"
	"../middlewares"
	"./../../bootstrap"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	frontendRoutes := b.Party("/")
	{
		frontendRoutes.Controller("/", new(controllers.HomeController))
		frontendRoutes.Controller("/login", new(controllers.LoginController))
		frontendRoutes.Get("/follower/{id:long}", GetFollowerHandler)
		frontendRoutes.Get("/following/{id:long}", GetFollowingHandler)
		frontendRoutes.Get("/like/{id:long}", GetLikeHandler)

	}

	backendRoutes := b.Party("/admin", middleware.BasicAuth)
	{
		//backendRoutes.Layout("/views/layouts/admin.html")
		backendRoutes.Controller("/", new(controllers.AdminController))
		backendRoutes.Controller("/users", new(controllers.UsersController))
	}

	// backendRoutes.Layout("/views/layouts/admin.html") // set a view layout for these routes, see more at view examples.

}
