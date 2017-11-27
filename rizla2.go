package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/kataras/rizla/rizla"
)

func main() {
	// Create a new project by the main source file
	project := rizla.NewProject("D:/repositories/go/frameworks/iris/iris_test1/main.go")

	// The below are optional

	// Optional, set the messages/logs output destination of our app project,
	// let's set them to their defaults
	project.Out = rizla.NewPrinter(os.Stdout)
	project.Err = rizla.NewPrinter(os.Stderr)

	project.Name = "My super project"
	// Allow reload every 3 seconds or more no less
	project.AllowReloadAfter = time.Duration(3) * time.Second
	// Custom subdirectory matcher, for the watcher, return true to include this folder to the watcher
	// the default is:
	project.Watcher = func(absolutePath string) bool {
		base := filepath.Base("D:/repositories/go/frameworks/iris/iris_test1/")
		return !(base == ".git" || base == "node_modules" || base == "vendor")
	}
	// Custom file matcher on runtime (file change), return true to reload when a file with this file name changed
	// the default is:
	project.Matcher = func(filename string) bool {
		isWindows := runtime.GOOS == "windows"
		goExt := ".go"
		return (filepath.Ext(fullname) == goExt) ||
			(!isWindows && strings.Contains(fullname, goExt))
	}
	// Add arguments, these will be used from the executable file
	project.Args = []string{"-myargument", "the value", "-otherargument", "a value"}
	// Custom callback before reload, the default is:
	project.OnReload = func(string) {
		fromproject := ""
		if p.Name != "" {
			fromproject = "From project '" + project.Name + "': "
		}
		project.Out.Infof("\n%sA change has been detected, reloading now...", fromproject)
	}
	// Custom callback after reload, the default is:
	project.OnReloaded = func(string) {
		project.Out.Successf("ready!\n")
	}

	// End of optional

	// Add the project to the rizla container
	rizla.Add(project)
	//  Build, run & start monitoring the project(s)
	rizla.Run()
}
