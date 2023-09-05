package controller

import (
    "github.com/gin-gonic/gin"
    "todo-list-ver-golang/adapter/service"
    "todo-list-ver-golang/internal/mapper"
    "todo-list-ver-golang/internal/resource/request"
)

type TodoController struct {
    todoService service.ITodoService
}

func NewTodoController(todoService service.ITodoService) *TodoController {
    return &TodoController{todoService: todoService}
}

func (t TodoController) CreateTodo(c *gin.Context) {
    var body request.CreateTodoRequest
    if err := c.BindJSON(&body); err != nil {
        c.JSON(400, gin.H{
            "message": "Bad Request",
        })
        return
    }
    createTodoDTO := mapper.CreateTodoRequestToDto(body)
    t.todoService.CreateTodo(c, &createTodoDTO)
}

func (t TodoController) GetAllTodo(c *gin.Context) {

}

func (t TodoController) DeleteTodo(c *gin.Context) {

}

func (t TodoController) UpdateTodo(c *gin.Context) {

}
