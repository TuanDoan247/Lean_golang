package port

import (
    "context"
    "gorm.io/gorm"
    "todo-list-ver-golang/core/entity"
)

type TodoRepository interface {
    CreateTodo(ctx context.Context, txDB *gorm.DB, todo *entity.Todo)
}
