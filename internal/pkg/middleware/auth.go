package middleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"maoim/internal/user"
)

func Auth(f func(string) (*user.User, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Request.Header.Get("Cookie")
		if cookie == "" {
			c.JSON(401, gin.H{"code": 401, "message": "no cookie"})
			c.Abort()
			return
		}
		decodeString, err := base64.StdEncoding.DecodeString(cookie)
		if err != nil {
			c.JSON(500, gin.H{"code": 500, "message": "Internal Error"})
			c.Abort()
			return
		}
		data := map[string]string{}
		err = json.Unmarshal(decodeString, &data)
		if err != nil {
			c.JSON(500, gin.H{"code": 500, "message": "Internal Error"})
			c.Abort()
			return
		}

		u, err := f(data["username"])
		if err != nil {
			c.JSON(500, gin.H{"code": 500, "message": "Internal Error"})
			c.Abort()
			return
		}
		c.Set("user", u)
		c.Next()
	}
}