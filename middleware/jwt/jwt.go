package jwt

import (
	"cloudcost-demo/controllers/base"
	"fmt"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = base.SUCCESS
		token := c.GetHeader("X-TOKEN")
		if token == "" {
			err := fmt.Errorf("code is empty")
			base.ReturnError(c, base.UNAUTHORIZED, err)
			return
		}
	}
}
