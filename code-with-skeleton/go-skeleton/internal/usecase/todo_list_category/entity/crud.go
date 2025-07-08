package entity

type TodoListCatReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	Name        string `json:"name,omitempty" validate:"required" name:"name"`
	Description string `json:"description" validate:"required" name:"deskripsi"`
}

type TodoListCatResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func (r *TodoListCatReq) SetID(ID int64) {
	r.ID = ID
}
