package controller

import (
	"github.com/gin-gonic/gin"
	"portofolio/go-restful/models"
)

type todoController struct {}

type todoPayload struct {
	Name string `json:"name" binding:"required"`
	Todo string `json:"todo" binding:"required"`
}

func (t *todoController) addTodo(context *gin.Context) {
	var addTodo todoPayload
	if err := context.ShouldBindJSON(&addTodo); err != nil {
		context.JSON(422, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	todoDb := models.Todo{
		Name: addTodo.Name,
		Todo: addTodo.Todo,
	}

	db := models.GetSqliteDbConnection()
	db.Create(&todoDb)

	context.JSON(200, gin.H{
		"success": true,
		"message": "add toto success",
	})
}

func (t *todoController) listTodo(context *gin.Context) {
	var todoList []models.Todo

	db := models.GetSqliteDbConnection()
	db.Model(&models.Todo{}).Find(&todoList)

	context.JSON(200, gin.H{
		"success": true,
		"data": todoList,
	})
}

func (t *todoController) updateTodo(context *gin.Context) {
	var todo models.Todo
	var payload todoPayload

	idTodo := context.Param("id")
	db := models.GetSqliteDbConnection()

	result := db.First(&todo, idTodo)

	if result.Error != nil {
		context.JSON(404, gin.H{
			"success": false,
			"message": "Todo not found",
		})
		return
	}

	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(422, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	todo.Name = payload.Name
	todo.Todo = payload.Todo
	db.Save(&todo)

	context.JSON(200, gin.H{
		"success": true,
		"message": "Update Success",
		"data": todo,
	})
}

func (t *todoController) deleteTodo(context *gin.Context) {
	idTodo := context.Param("id")
	db := models.GetSqliteDbConnection()
	result := db.Delete(&models.Todo{}, idTodo)

	if result.Error != nil {
		context.JSON(400, gin.H{
			"success": false,
			"message": "Error Delete Todo",
		})
		return
	}

	context.JSON(200, gin.H{
		"success": true,
		"message": "Delete Success",
	})
}