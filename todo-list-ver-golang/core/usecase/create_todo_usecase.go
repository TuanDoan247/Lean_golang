package usecase

import (
    "context"
    "gorm.io/gorm"
    "todo-list-ver-golang/core/entity"
    "todo-list-ver-golang/core/entity/dto"
    "todo-list-ver-golang/core/port"
)

type ICreateTodoUseCase interface {
    MakeTodo(ctx context.Context, database *gorm.DB, request *dto.CreateTodoDTO)
}

type CreateTodoUseCase struct {
    todoRepository port.TodoRepository
}

func NewCreateTodoUseCase(todoRepository port.TodoRepository) ICreateTodoUseCase {
    return &CreateTodoUseCase{
        todoRepository: todoRepository,
    }
}

func (cous *CreateTodoUseCase) MakeTodo(ctx context.Context, database *gorm.DB, request *dto.CreateTodoDTO) {
    todo := &entity.Todo{
        Title:   &request.Title,
        Content: &request.Content,
    }

    cous.todoRepository.CreateTodo(ctx, database, todo)
}
