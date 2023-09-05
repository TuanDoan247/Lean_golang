package request

type UpdateTodoRequest struct {
	ID 			int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
