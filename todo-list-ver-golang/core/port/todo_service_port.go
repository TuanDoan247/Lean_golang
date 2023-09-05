package port

import "todo-list-ver-golang/core/entity/dto"

type ITodoService interface {
    CreateTodo(createTodoReq *dto.CreateTodoDTO)
}
