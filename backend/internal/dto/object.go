package dto

type ObjectResponse struct {
	ID        uint   `json:"id"`
	Key       string `json:"key"`
	Size      int64  `json:"size"`
	Mime      string `json:"mime"`
	Checksum  string `json:"checksum"`
	CreatedAt string `json:"created_at"`
}

type ListObjectsParams struct {
	Page    int `form:"page,default=1" binding:"min=1"`
	PerPage int `form:"per_page,default=20" binding:"min=1,max=100"`
}
