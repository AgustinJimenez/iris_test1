package controllers

import (
	"../../datamodels"
	"github.com/kataras/iris/mvc"
)

type UserController struct {
	mvc.C
}

func (c *UserController) Get() mvc.Result {
	return mvc.View{
		Name: "pages/users/index.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}

func (c *UserController) GetBy(id int64) (user datamodels.User, found bool) {
	u, found := c.Service.GetByID(id)
	if !found {
		// this message will be binded to the
		// main.go -> app.OnAnyErrorCode -> NotFound -> shared/error.html -> .Message text.
		c.Ctx.Values().Set("message", "User couldn't be found!")
	}
	return u, found // it will throw/emit 404 if found == false.
}
