package controllers

import (
	"github.com/kataras/iris/mvc"
)

type LoginController struct {
	mvc.C
}

func (c *LoginController) Get() mvc.Result {
	return mvc.View{
		Name: "templates/login/index.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}
