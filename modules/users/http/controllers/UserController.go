// file: controllers/user_controller.go

package UserControllers

import (
	"fmt"

	"../../datamodels"
	"../../services"
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
	Service services.UserService

	// Session-relative things.
	Manager *sessions.Sessions
	Session *sessions.Session
}

func (this *UserController) BeginRequest(ctx context.Context) {
	this.C.BeginRequest(ctx)

	if this.Manager == nil {
		ctx.Application().Logger().Errorf(`UserController: sessions manager is nil, you should bind it`)
		ctx.StopExecution() // dont run the main method handler and any "done" handlers.
		return
	}
	this.Session = this.Manager.Start(ctx)
}

const userIDKey = "UserID"

func (this *UserController) getCurrentUserID() int64 {
	userID, _ := this.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (this *UserController) isLoggedIn() bool {
	return this.getCurrentUserID() > 0
}

func (this *UserController) logout() {
	this.Manager.DestroyByID(this.Session.ID())
}

// GetRegister handles GET: http://localhost:8080/user/register.
func (this *UserController) GetRegister() mvc.Result {
	if this.isLoggedIn() {
		this.logout()
	}

	return mvc.View{
		Name:   "users/resources/views/frontend/register/index.html",
		Data:   context.Map{"Title": "User Registration"},
		Layout: "core/resources/templates/simple-form/index.html",
	}
}

// PostRegister handles POST: http://localhost:8080/user/register.
func (this *UserController) PostRegister() mvc.Result {

	new_service := services.UserService{}
	new_user, err := new_service.Create(this.Ctx.FormValue("password"), datamodels.User{

		Firstname: this.Ctx.FormValue("first_name"),
		Lastname:  this.Ctx.FormValue("last_name"),
		Username:  this.Ctx.FormValue("username"),
		Email:     this.Ctx.FormValue("email"),
	})
	this.Session.Set(userIDKey, new_user.Id)
	//c.Status = 303
	//c.Path = pathMyProfile

	return mvc.Response{
		// if not nil then this error will be shown instead.
		Err: err,
		// redirect to /user/me.
		Path: "/admin",
		// When redirecting from POST to GET request you -should- use this HTTP status code,
		// however there're some (complicated) alternatives if you
		// search online or even the HTTP RFC.
		// Status "See Other" RFC 7231, however iris can automatically fix that
		// but it's good to know you can set a custom code;
		// Code: 303,
	}
}

// GetRegister handles GET: http://localhost:8080/user/register.
func (this *UserController) GetAdmin() mvc.Result {

	fmt.Println("USER SESSION user_id AND isLoggedIn =====>", this.getCurrentUserID(), this.isLoggedIn())
	fmt.Println("CONTEXT=====>", this)
	return mvc.View{
		Name:   "users/resources/views/backend/dashboard/index.html",
		Data:   context.Map{"Title": "User Registration"},
		Layout: "core/resources/templates/admin/index.html",
	}
}

// GetLogin handles GET: http://localhost:8080/user/login.
func (this *UserController) GetLogin() mvc.Result {
	if this.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
		this.logout()
	}

	return mvc.View{
		Name:   "users/resources/views/frontend/login/index.html",
		Data:   context.Map{"Title": "User Login"},
		Layout: "core/resources/templates/simple-form/index.html",
	}
}

// PostLogin handles POST: http://localhost:8080/user/register.
func (this *UserController) PostLogin() mvc.Result {

	u, err := this.Service.GetByUsernameAndPassword(this.Ctx.FormValue("username"), this.Ctx.FormValue("password"))

	this.Session.Set(userIDKey, u.Id)
	fmt.Println("ON POST LOGIN USER IS =====>", this.getCurrentUserID(), this.isLoggedIn())

	return mvc.Response{
		Err:  err,
		Path: "/admin",
	}

}

func (this *UserController) AnyLogout() {
	if this.isLoggedIn() {
		this.logout()
	}

	this.Ctx.Redirect("/login")
}
