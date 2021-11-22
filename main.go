package main

import (
	"fmt"
	"net/http"
	"strconv"

	_ "finalproject1/docs"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID       int    `example: "1"`
	Name     string `example: "makan"`
	Complete bool   `example: "false"`
}

var (
	data = []Todo{}
)

var todoN int = 0

// @Tags Todos
// @Summary Get details of all todos
// @Description Get details of all todos
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	var result = gin.H{
		"result": data,
	}

	c.JSON(http.StatusOK, result)
}

// @Tags Todos
// @Summary Create a new todo
// @Description Create a new todo with the input payload
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} Todo
// @Param id path int true "Create todo"
// @Router /todo [post]
func createTodo(c *gin.Context) {
	var todo Todo

	todo.ID = todoN
	todo.Name = c.PostForm("name")
	todo.Complete = false

	data = append(data, todo)

	todoN++

	result := gin.H{
		"result": todo,
	}

	c.JSON(http.StatusOK, result)
}

// @Tags Todos
// @Summary Get specific todo
// @Description Get specific todo
// @Produce json
// @Success 200 {object} Todo
// @Param id path int true "Get todos by Id"
// @Router /todos/{id} [get]
func getTodoById(c *gin.Context) {
	var result gin.H
	id, _ := strconv.Atoi(c.Param("id"))

	if id < len(data) {
		result = gin.H{
			"result": data[id],
		}
		c.JSON(http.StatusOK, result)
		return
	} else {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	}
}

// @Tags Todos
// @Summary Update specific todo
// @Description Update specific todo
// @accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} Todo
// @Param id path int true "Update todos by Id"
// @Router /todos/{id} [put]
func updateTodo(c *gin.Context) {
	var result gin.H

	i, _ := strconv.Atoi(c.Param("id"))

	if i < len(data) {
		// data[i].ID = i
		data[i].Name = c.PostForm("name")
		data[i].Complete, _ = strconv.ParseBool(c.PostForm("complete"))

		result = gin.H{
			"result": data[i],
		}

		c.JSON(http.StatusOK, result)
	} else {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	}
}

// @Tags Todos
// @Summary Delete specific todo
// @Description Delete specific todo
// @Produce json
// @Success 200 {object} Todo
// @Param id path int true "Delete todos by Id"
// @Router /todos/{id} [delete]
func deleteTodo(c *gin.Context) {
	// fmt.Println(c.Query("id"), "id")
	var result gin.H

	i, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(i, len(data))

	if i < len(data) {
		copy(data[i:], data[i+1:]) // Shift a[i+1:] left one index.
		data[len(data)-1] = Todo{} // Erase last element (write zero value).
		data = data[:len(data)-1]  // Truncate slice.

		result = gin.H{
			"result": "delete success",
		}
		c.JSON(http.StatusOK, result)
		return
	} else {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	}

}

// @title Todo API
// @version 1.0
// @description This is a sample API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /todos
func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.GET(":id", getTodoById)
			todos.GET("", GetTodos)
			todos.POST("", createTodo)
			todos.DELETE(":id", deleteTodo)
			todos.PUT(":id", updateTodo)
		}
		//...
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")

}
