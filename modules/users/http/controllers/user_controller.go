// file: controllers/user_controller.go

package UserControllers

import (
	"../../datamodels"
	"../../services"

	"fmt"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

// UserController is our /user controller.
// UserController is responsible to handle the following requests:
// GET  			/user/register
// POST 			/user/register
// GET 				/user/login
// POST 			/user/login
// GET 				/user/me
// All HTTP Methods /user/logout
type UserController struct {
	// mvc.C is just a lightweight lightweight alternative
	// to the "mvc.Controller" controller type,
	// use it when you don't need mvc.Controller's fields
	// (you don't need those fields when you return values from the method functions).
	mvc.C

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service *services.UserService

	// Session-relative things.
	Manager *sessions.Sessions
	Session *sessions.Session
}

// BeginRequest will set the current session to the controller.
//
// Remember: iris.Context and context.Context is exactly the same thing,
// iris.Context is just a type alias for go 1.9 users.
// We use context.Context here because we don't need all iris' root functions,
// when we see the import paths, we make it visible to ourselves that this file is using only the context.
func (c *UserController) BeginRequest(ctx context.Context) {
	c.C.BeginRequest(ctx)
	if c.Manager == nil {
		ctx.Application().Logger().Errorf(`UserController: sessions manager is nil, you should bind it`)
		ctx.StopExecution() // dont run the main method handler and any "done" handlers.
		return
	}

	c.Session = c.Manager.Start(ctx)
}

const userIDKey = "UserID"

func (c *UserController) getCurrentUserID() int64 {
	return 22
}

func (c *UserController) isLoggedIn() bool {
	return false
}

func (c *UserController) logout() {

}

// GetRegister handles GET: http://localhost:8080/user/register.
func (c *UserController) GetRegister() mvc.Result {
	if c.isLoggedIn() {
		c.logout()
	}

	return mvc.View{
		Name:   "users/resources/views/frontend/register/index.html", /*"pages/users/register.html"*/
		Data:   context.Map{"Title": "User Registration"},
		Layout: "core/resources/templates/simple-form/index.html",
	}
}

// PostRegister handles POST: http://localhost:8080/user/register.
func (c *UserController) PostRegister() mvc.Result {
	fmt.Printf("\nHERE IS CONTROLLER=======================>\n")
	// get firstname, username and password from the form.

	// create the new user, the password will be hashed by the service.
	/*u, err := */
	new_user := services.Create(datamodels.User{

		Firstname: c.Ctx.FormValue("first_name"),
		Lastname:  c.Ctx.FormValue("last_name"),
		Username:  c.Ctx.FormValue("username"),
		Email:     c.Ctx.FormValue("email"),
		Password:  c.Ctx.FormValue("password"),
	})
	fmt.Printf("\nHERE IS CONTROLLER END CREATE=======================>\n", new_user)
	/*
		// set the user's id to this session even if err != nil,
		// the zero id doesn't matters because .getCurrentUserID() checks for that.
		// If err != nil then it will be shown, see below on mvc.Response.Err: err.
		c.Session.Set(userIDKey, u.ID)

		return mvc.Response{
			// if not nil then this error will be shown instead.
			Err: err,
			// redirect to /user/me.
			Path: "/user/me",
			// When redirecting from POST to GET request you -should- use this HTTP status code,
			// however there're some (complicated) alternatives if you
			// search online or even the HTTP RFC.
			// Status "See Other" RFC 7231, however iris can automatically fix that
			// but it's good to know you can set a custom code;
			// Code: 303,
		}
	*/
	//?first_name=Agustin&last_name=Jimenez&username=agusjim&email=agus.jimenez.caba@gmail.com&password=123456
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<p> saved, check your db <p>",
	}
}

// GetLogin handles GET: http://localhost:8080/user/login.
func (c *UserController) GetLogin() mvc.Result {
	if c.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
		c.logout()
	}

	return mvc.View{
		Name:   "users/resources/views/frontend/login/index.html",
		Data:   context.Map{"Title": "User Login"},
		Layout: "core/resources/templates/simple-form/index.html",
	}
}

// PostLogin handles POST: http://localhost:8080/user/register.
func (c *UserController) PostLogin() mvc.Result {

	return mvc.Response{
		Path: "/user/me",
	}
}

//c.Ctx.Redirect("/user/login")
