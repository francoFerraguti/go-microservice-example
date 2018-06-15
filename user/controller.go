package user

import(
	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	c.JSON(200, "{user: franco}")
}