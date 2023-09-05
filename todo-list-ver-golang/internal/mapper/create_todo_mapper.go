package mapper

import (
	"todo-list-ver-golang/core/entity/dto"
	"todo-list-ver-golang/internal/resource/request"
)

func CreateTodoRequestToDto(createTodoRequest request.CreateTodoRequest) dto.CreateTodoDTO {
	return dto.CreateTodoDTO{
		Title:   createTodoRequest.Title,
		Content: createTodoRequest.Description,
	}
}
