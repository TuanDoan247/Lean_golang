package mapper

import (
    "todo-list-ver-golang/adapter/repository/mysql/model"
    "todo-list-ver-golang/core/entity"
)

func TodoEntityToModel(todoEntity *entity.Todo) *model.Todo {
    todoModel := &model.Todo{
        Title:   todoEntity.Title,
        Content: todoEntity.Content,
    }
    return todoModel
}
