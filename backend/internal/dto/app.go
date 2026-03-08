package dto

type CreateAppRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
}

type AppResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	OwnerID   uint   `json:"owner_id"`
	CreatedAt string `json:"created_at"`
}
