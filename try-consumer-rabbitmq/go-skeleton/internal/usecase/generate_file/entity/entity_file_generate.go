package entity

type GenerateReq struct {
	Filename string `json:"filename,omitempty" validate:"required" name:"Filename"`
	Message  string `json:"message" validate:"required" name:"Container"`
}

type GenerateResponse struct {
	Filename   string `json:"filename,omitempty"`
	Message    string `json:"message"`
	Status     string `json:"status"`
	GenerateAt string `json:"generate_at"`
}
