package routes

import (
	PagesControllers "../../modules/pages/http/controllers"
	UserControllers "../../modules/users/http/controllers"
	middleware "../middlewares"
	"./../../bootstrap"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	backendRoutes := b.Party("/admin", middleware.BasicAuth)
	{
		backendRoutes.Layout("/views/layouts/admin.html")
		backendRoutes.Controller("/", new(UserControllers.UserController))
		backendRoutes.Controller("/logout", new(controllers.AdminController))

		backendRoutes.Controller("/users", new(UserControllers.UsersController))
	}

	frontendRoutes := b.Party("/")
	{
		frontendRoutes.Controller("/", new(PagesControllers.HomeController))
		frontendRoutes.Controller("/", new(UserControllers.UserController), b.Sessions)
		frontendRoutes.Controller("/register", new(UserControllers.UserController))

		//frontendRoutes.Get("/follower/{id:long}", GetFollowerHandler)
		//frontendRoutes.Get("/following/{id:long}", GetFollowingHandler)
		//frontendRoutes.Get("/like/{id:long}", GetLikeHandler)

	}

}
