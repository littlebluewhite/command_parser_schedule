package command_template

import (
	"command_parser_schedule/entry/e_command_template"
	"command_parser_schedule/util"
	"command_parser_schedule/util/logFile"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)

type Handler struct {
	O Operate
	L logFile.LogFile
}

// GetCommandTemplates swagger
// @Summary     Show all command templates
// @Description Get all command templates
// @Tags        command_template
// @Produce     json
// @Success     200 {array} e_command_template.CommandTemplate
// @Router      /command_template/ [get]
func (h *Handler) GetCommandTemplates(c *gin.Context) {
	ct, err := h.O.List()
	result := e_command_template.Format(ct)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetCommandTemplates: ", err)
		return
	}
	h.L.Info().Println("GetCommandTemplates: success")
	c.JSON(200, result)
	return
}

// GetCommandTemplateById swagger
// @Summary     Show command templates
// @Description Get command templates by id
// @Tags        command_template
// @Produce     json
// @Param       id  path     int true "command template id"
// @Success     200 {object} e_command_template.CommandTemplate
// @Router      /command_template/{id} [get]
func (h *Handler) GetCommandTemplateById(c *gin.Context) {
	id := c.Param("id")
	IdInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetCommandTemplateById: ", err)
		return
	}
	ht, err := h.O.Find([]int32{int32(IdInt)})
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetCommandTemplateById: ", err)
		return
	}
	h.L.Info().Println("GetCommandTemplateById: success")
	c.JSON(200, e_command_template.Format(ht)[0])
	return
}

// AddCommandTemplate swagger
// @Summary Create command templates
// @Tags    command_template
// @Accept  json
// @Produce json
// @Param   command_template body     []e_command_template.CommandTemplateCreate true "command template body"
// @Success 200           {array} e_command_template.CommandTemplate
// @Router  /command_template/ [post]
func (h *Handler) AddCommandTemplate(c *gin.Context) {
	entry := []*e_command_template.CommandTemplateCreate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddCommandTemplate: ", err)
		return
	}
	result, err := h.O.Create(entry)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddCommandTemplate: ", err)
		return
	}
	c.JSON(200, e_command_template.Format(result))
}

// DeleteCommandTemplate swagger
// @Summary Delete command templates
// @Tags    command_template
// @Produce json
// @Param ids body []int true "command_template id"
// @Success 200 {string} string "delete successfully"
// @Router  /command_template/ [delete]
func (h *Handler) DeleteCommandTemplate(c *gin.Context) {
	entry := make([]int32, 0, 10)
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteCommandTemplate: ", err)
		return
	}
	err := h.O.Delete(entry)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteCommandTemplate: ", err)
		return
	}
	c.JSON(200, "delete successfully")
}
