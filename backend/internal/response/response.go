package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type body struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *errorBody  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page    int   `json:"page"`
	PerPage int   `json:"per_page"`
	Total   int64 `json:"total"`
}

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, body{Success: true, Data: data})
}

func Error(c *gin.Context, status int, code, message string) {
	c.JSON(status, body{
		Success: false,
		Error:   &errorBody{Code: code, Message: message},
	})
}

func Paginated(c *gin.Context, data interface{}, meta Meta) {
	c.JSON(http.StatusOK, body{Success: true, Data: data, Meta: &meta})
}
