package main

import (
	app "portofolio/go-restful/app"
	controller "portofolio/go-restful/controller"
)

func main() {
	restful := app.CreateRestfulApp()

	controller.RegisterController(restful)

	restful.Run()
}