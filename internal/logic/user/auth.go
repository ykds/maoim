package user

import "github.com/gin-gonic/gin"

func Auth(srv Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(401, gin.H{"code": 401, "message": "缺少token"})
			c.Abort()
			return
		}

		u, err := srv.Auth(token)
		if err != nil {
			c.JSON(401, gin.H{"code": 401, "message": "token认证失败," + err.Error()})
			c.Abort()
			return
		}

		c.Set("user", u)
		c.Next()
	}
}
