package time_template

import (
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

// GetTimeTemplates swagger
// @Summary     Show all time templates
// @Description Get all time templates
// @Tags        time_template
// @Produce     json
// @Success     200 {array} time_template.TimeTemplate
// @Router      /time_template/ [get]
func (h *Handler) GetTimeTemplates(c *gin.Context) {
	tt, err := h.O.List()
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetTimeTemplates: ", err)
		return
	}
	h.L.Info().Println("GetTimeTemplates: success")
	c.JSON(200, Format(tt))
	return
}

// GetTimeTemplateById swagger
// @Summary     Show time templates
// @Description Get time templates by id
// @Tags        time_template
// @Produce     json
// @Param       id  path     int true "time template id"
// @Success     200 {object} time_template.TimeTemplate
// @Router      /time_template/{id} [get]
func (h *Handler) GetTimeTemplateById(c *gin.Context) {
	id := c.Param("id")
	IdInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetTimeTemplateById: ", err)
		return
	}
	tt, err := h.O.Find([]int32{int32(IdInt)})
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetTimeTemplateById: ", err)
		return
	}
	result := Format(tt)
	h.L.Info().Println("GetTimeTemplateById: success")
	c.JSON(200, result[0])
	return
}

// AddTimeTemplate swagger
// @Summary Create time templates
// @Tags    time_template
// @Accept  json
// @Produce json
// @Param   time_template body     []time_template.TimeTemplateCreate true "time template body"
// @Success 200           {array} time_template.TimeTemplate
// @Router  /time_template/ [post]
func (h *Handler) AddTimeTemplate(c *gin.Context) {
	entry := []*TimeTemplateCreate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddTimeTemplate: ", err)
		return
	}
	tt, err := h.O.Create(entry)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddTimeTemplate: ", err)
		return
	}
	c.JSON(200, Format(tt))
}

// UpdateTimeTemplate swagger
// @Summary Update time templates
// @Tags    time_template
// @Accept  json
// @Produce json
// @Param   time_template body     []time_template.TimeTemplateUpdate true "modify time template body"
// @Success 200           {string} string "update successfully"
// @Router  /time_template/ [patch]
func (h *Handler) UpdateTimeTemplate(c *gin.Context) {
	entry := []*TimeTemplateUpdate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateTimeTemplate: ", err)
		return
	}
	err := h.O.Update(entry)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateTimeTemplate: ", err)
		return
	}
	c.JSON(200, "update successfully")
}

// DeleteTimeTemplate swagger
// @Summary Delete time templates
// @Tags    time_template
// @Produce json
// @Param ids body []int true "time_template id"
// @Success 200 {string} string "delete successfully"
// @Router  /time_template/ [delete]
func (h *Handler) DeleteTimeTemplate(c *gin.Context) {
	entry := make([]int32, 0, 10)
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteTimeTemplate: ", err)
		return
	}
	err := h.O.Delete(entry)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteTimeTemplate: ", err)
		return
	}
	c.JSON(200, "delete successfully")
}
