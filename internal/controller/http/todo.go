package http

import (
	"gotest/internal/core/common/router"
	"gotest/internal/core/model/request"
	"gotest/internal/core/port/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	gin         *gin.Engine
	todoService service.TodoService
}

func NewTodoController(gin *gin.Engine, todoService service.TodoService) TodoController {
	return TodoController{
		gin:         gin,
		todoService: todoService,
	}
}

func (t TodoController) InitRouter() {
	api := t.gin.Group("/api/v1")
	router.Post(api, "add-todo", t.addTodo)
	router.Get(api, "get-todos", t.getListTodo)
}

func (t TodoController) addTodo(c *gin.Context) {
	var req request.AddTodoRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &invalidRequestResponse)
		return
	}

	respon := t.todoService.AddTodo(&req)
	c.JSON(http.StatusOK, respon)
}

func (t TodoController) getListTodo(c *gin.Context) {
	res := t.todoService.GetListTodo()

	c.JSON(http.StatusOK, res)
}
