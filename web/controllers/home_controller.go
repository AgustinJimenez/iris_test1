package controllers

import (
	"github.com/kataras/iris/mvc"
)

type HomeController struct {
	mvc.C
}

func (c *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "templates/simple/index.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}
