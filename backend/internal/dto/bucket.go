package dto

type CreateBucketRequest struct {
	Name       string `json:"name" binding:"required,min=3,max=63"`
	Visibility string `json:"visibility" binding:"required,oneof=public private"`
}

type BucketResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Visibility string `json:"visibility"`
	CreatedAt  string `json:"created_at"`
}
