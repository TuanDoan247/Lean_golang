package service

import (
    "context"
    "todo-list-ver-golang/core/entity/dto"
    "todo-list-ver-golang/core/usecase"
)

type ITodoService interface {
    CreateTodo(ctx context.Context, createTodoReq *dto.CreateTodoDTO)
}

type TodoService struct {
    createTodoUseCase   usecase.ICreateTodoUseCase
    databaseTransaction usecase.IDatabaseTransactionUseCase
}

func NewTodoService(
        createToDoUseCase usecase.ICreateTodoUseCase,
        databaseTraction usecase.IDatabaseTransactionUseCase,
) ITodoService {
    return &TodoService{
        createTodoUseCase:   createToDoUseCase,
        databaseTransaction: databaseTraction,
    }
}

func (os *TodoService) CreateTodo(ctx context.Context, createTodoReq *dto.CreateTodoDTO) {
    dbTraction := os.databaseTransaction.StartTransaction()
    os.createTodoUseCase.MakeTodo(ctx, dbTraction, createTodoReq)
}
