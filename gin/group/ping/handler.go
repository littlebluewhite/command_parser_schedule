package ping

import (
	"command_parser_schedule/util/logFile"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	O Operate
	L logFile.LogFile
}

func (h *Handler) GetPing(c *gin.Context) {
	example := c.MustGet("example").(string)
	c.JSON(200, gin.H{
		"message": example,
	})
	h.L.Info().Println("get ping: example: ", example)
}

func (h *Handler) GetListPing(c *gin.Context) {
	data := []map[string]interface{}{
		{
			"name": "wilson",
			"age":  5,
			"time": time.Now(),
		},
		{
			"name": "phoebe",
			"age":  4,
		},
	}
	c.JSON(200, data)
	h.L.Info().Println("get ping list: data: ", data)
}
