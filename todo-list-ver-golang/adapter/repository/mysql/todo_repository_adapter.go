package mysql

import (
    "context"
    "gorm.io/gorm"
    "todo-list-ver-golang/adapter/repository/mysql/mapper"
    "todo-list-ver-golang/core/entity"
    "todo-list-ver-golang/core/port"
)

type TodoRepositoryAdapter struct {
    base
}

func NewTodoRepositoryAdapter(db *gorm.DB) port.TodoRepository {
    return &TodoRepositoryAdapter{base{db: db}}
}

func (t *TodoRepositoryAdapter) CreateTodo(ctx context.Context, txDB *gorm.DB, todo *entity.Todo) {
    todoModel := mapper.TodoEntityToModel(todo)
    err := txDB.WithContext(ctx).Create(todoModel).Error
    if err != nil {
        txDB.Rollback()
        panic(err)
    }
    txDB.Commit()
}
