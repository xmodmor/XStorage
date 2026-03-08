package dto

type CreateAPIKeyRequest struct {
	Permissions string `json:"permissions" binding:"omitempty"`
}

type APIKeyResponse struct {
	ID          uint   `json:"id"`
	AppID       uint   `json:"app_id"`
	AccessKey   string `json:"access_key"`
	SecretKey   string `json:"secret_key,omitempty"`
	Permissions string `json:"permissions"`
	CreatedAt   string `json:"created_at"`
}
