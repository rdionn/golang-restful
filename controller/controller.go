package controller

import (
	app "portofolio/go-restful/app"
)

func RegisterController(restfulapp app.App) {
	todoController := &todoController{}

	restfulapp.AddPost("/todo", todoController.addTodo)
	restfulapp.AddGet("/todos", todoController.listTodo)
	restfulapp.AddPut("/todo/:id", todoController.updateTodo)
	restfulapp.AddDelete("/todo/:id", todoController.deleteTodo)
}