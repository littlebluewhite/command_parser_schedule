package command_template

import (
	"command_parser_schedule/util"
	"command_parser_schedule/util/logFile"
	"errors"
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
// @Success     200 {array} command_template.CommandTemplate
// @Router      /command_template/ [get]
func (h *Handler) GetCommandTemplates(c *gin.Context) {
	ht, err := h.O.List()
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetCommandTemplates: ", err)
		return
	}
	h.L.Info().Println("GetCommandTemplates: success")
	c.JSON(200, Format(ht))
	return
}

// GetCommandTemplateById swagger
// @Summary     Show command templates
// @Description Get command templates by id
// @Tags        command_template
// @Produce     json
// @Param       id  path     int true "command template id"
// @Success     200 {object} command_template.CommandTemplate
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
	if len(ht) == 0 {
		util.Err(c, errors.New("empty command template"), 0)
		h.L.Error().Println("GetCommandTemplateById: ", "empty command template")
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetCommandTemplateById: ", err)
		return
	}
	h.L.Info().Println("GetCommandTemplateById: success")
	c.JSON(200, Format(ht)[0])
	return
}

// AddCommandTemplate swagger
// @Summary Create command templates
// @Tags    command_template
// @Accept  json
// @Produce json
// @Param   command_template body     []command_template.CommandTemplateCreate true "command template body"
// @Success 200           {array} command_template.CommandTemplate
// @Router  /command_template/ [post]
func (h *Handler) AddCommandTemplate(c *gin.Context) {
	entry := []*CommandTemplateCreate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddCommandTemplate: ", err)
		return
	}
	ht := CreateConvert(entry)
	ht, err := h.O.Create(ht)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddCommandTemplate: ", err)
		return
	}
	c.JSON(200, Format(ht))
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
	ht, err := h.O.Find(entry)
	if len(ht) == 0 {
		util.Err(c, errors.New("empty command template"), 0)
		h.L.Error().Println("DeleteCommandTemplate: ", err)
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteCommandTemplate: ", err)
		return
	}
	err = h.O.Delete(ht)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteCommandTemplate: ", err)
		return
	}
	c.JSON(200, "delete successfully")
}
