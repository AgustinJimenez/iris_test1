package main

import (
	"github.com/kataras/rizla/rizla"
)

func main() {
	// Build, run & start monitoring the projects
	rizla.Run("D:/repositories/go/frameworks/iris/iris_test1/main.go")
	// rizla.RunWith(rizla.WatcherFromFlag("-walk"), "./main.go")
}
