package controllers

import (
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	mvc.C
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "templates/admin/index.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}
