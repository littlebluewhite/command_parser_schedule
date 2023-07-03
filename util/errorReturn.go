package util

import "github.com/gin-gonic/gin"

func Err(c *gin.Context, err error, inner int) {
	c.AbortWithStatusJSON(484, gin.H{"message": err.Error(), "inner_code": inner})
}
