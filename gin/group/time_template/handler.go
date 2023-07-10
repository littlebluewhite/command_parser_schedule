package time_template

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

// GetTimeTemplates swagger
// @Summary     Show all time templates
// @Description Get all time templates
// @Tags        time_template
// @Produce     json
// @Success     200 {array} time_template.TimeTemplate
// @Router      /time_template/ [get]
func (h *Handler) GetTimeTemplates(c *gin.Context) {
	tt, err := h.O.List()
	result := Format(tt)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetheaderTemplates: ", err)
		return
	}
	h.L.Info().Println("GetheaderTemplates: success")
	c.JSON(200, result)
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
	if len(tt) == 0 {
		util.Err(c, errors.New("empty time template"), 0)
		h.L.Error().Println("GetTimeTemplateById: ", "empty time template")
		return
	}
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
	tt := CreateConvert(entry)
	tt, err := h.O.Create(tt)
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
	ids := make([]int32, 0, len(entry))
	uMap := make(map[int32]*TimeTemplateUpdate)
	for _, item := range entry {
		ids = append(ids, item.ID)
		uMap[item.ID] = item
	}
	tt, err := h.O.Find(ids)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateTimeTemplate: ", err)
		return
	}
	tt = UpdateConvert(tt, uMap)
	err = h.O.Update(tt)
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
	tt, err := h.O.Find(entry)
	if len(tt) == 0 {
		util.Err(c, errors.New("empty time template"), 0)
		h.L.Error().Println("DeleteTimeTemplate: ", err)
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteTimeTemplate: ", err)
		return
	}
	err = h.O.Delete(tt)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteTimeTemplate: ", err)
		return
	}
	c.JSON(200, "delete successfully")
}
