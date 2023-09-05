package router

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/fx"
    "todo-list-ver-golang/internal/controller"
)

type RegisterRouterIn struct {
    fx.In
    Engine         *gin.Engine
    TodoController *controller.TodoController
}

func RegisterGinRouters(p RegisterRouterIn) {
    todo := p.Engine.Group("/v1/todo")
    {
        todo.POST("", p.TodoController.CreateTodo)
        todo.GET("", p.TodoController.GetAllTodo)
        todo.DELETE("", p.TodoController.DeleteTodo)
        todo.PUT("", p.TodoController.UpdateTodo)
    }
}
