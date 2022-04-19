package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, status int, msg string, data interface{}) {
	body := gin.H{"code": status, "message": msg}
	if data != nil {
		body["data"] = data
	}
	c.JSON(status, body)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	body := gin.H{"code": http.StatusOK, "message": "Success"}
	if data != nil {
		body["data"] = data
	}
	c.JSON(http.StatusOK, body)
}
