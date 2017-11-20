//file: main.go

package main

import (
	"github.com/kataras/iris"
	/*========================================*/
	"./app/datasources/movies"
	"./app/repositories/movies"
	"./app/services/movies"
	"./app/web/controllers/movies"
	"./app/web/middlewares"
)

func main() {
	app := iris.New()

	// Load the template files.
	app.RegisterView(iris.HTML("./app/modules/movies/web", ".html"))

	// Register our controllers.
	app.Controller("/hello", new(controllers.HelloController))

	// Create our movie repository with some (memory) data from the datasource.
	repo := repositories.NewMovieRepository(datasource.Movies)
	// Create our movie service, we will bind it to the movie controller.
	movieService := services.NewMovieService(repo)

	app.Controller("/movies", new(controllers.MovieController),
		// Bind the "movieService" to the MovieController's Service (interface) field.
		movieService,
		// Add the basic authentication(admin:password) middleware
		// for the /movies based requests.
		middleware.BasicAuth)

	// Start the web server at localhost:8080
	// http://localhost:8080/hello
	// http://localhost:8080/hello/iris
	// http://localhost:8080/movies
	// http://localhost:8080/movies/1
	app.Run(
		iris.Addr("localhost:8000"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, // enables faster json serialization and more
	)
}
