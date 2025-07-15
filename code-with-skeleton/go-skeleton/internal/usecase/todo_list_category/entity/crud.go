package entity

type TodoListCatReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	UserID      int64  `json:"user_id,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required" name:"name"`
	Description string `json:"description" validate:"required" name:"deskripsi"`
}

func (r *TodoListCatReq) SetUserID(i int64) {
	r.UserID = i
}

type TodoListCatResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
}

func (r *TodoListCatReq) SetID(ID int64) {
	r.ID = ID
}
