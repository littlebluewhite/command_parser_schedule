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

// GetPing swagger
// @Summary    test ping
// @Description test ping
// @Tags        ping
// @Produce     json
// @Success     200 {object} ping.SwaggerPing
// @Router      /ping/test [get]
func (h *Handler) GetPing(c *gin.Context) {
	example := c.MustGet("example").(string)
	c.JSON(200, gin.H{
		"message": example,
	})
	h.L.Info().Println("get ping: example: ", example)
}

// GetListPing swagger
// @Summary     return list ping
// @Description test list ping
// @Tags        ping
// @Produce     json
// @Success     200 {array} ping.SwaggerListPing
// @Router      /ping/list [get]
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
