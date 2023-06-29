package time_template

import (
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
	result := Format(tt)
	if err != nil {
		c.AbortWithStatusJSON(484, err)
		h.L.Error().Println("GetTimeTemplates: ", err)
		return
	}
	h.L.Info().Println("GetTimeTemplates: success")
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
		c.AbortWithStatusJSON(484, err)
		h.L.Error().Println("GetTimeTemplateById: ", err)
		return
	}
	tt, err := h.O.Find([]int32{int32(IdInt)})
	if len(tt) == 0 {
		c.AbortWithStatusJSON(484, "empty time template")
		h.L.Error().Println("GetTimeTemplateById: ", "empty time template")
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(484, err)
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
		c.AbortWithStatusJSON(484, err.Error())
		return
	}
	tt := CreateConvert(entry)
	tt, err := h.O.Create(tt)
	if err != nil {
		c.AbortWithStatusJSON(484, err.Error())
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
// @Success 200           {string} string "Update successfully
// @Router  /time_template/ [patch]
func (h *Handler) UpdateTimeTemplate(c *gin.Context) {
	entry := []*TimeTemplateUpdate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		c.AbortWithStatusJSON(484, err.Error())
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
		c.AbortWithStatusJSON(484, err.Error())
		return
	}
	tt = UpdateConvert(tt, uMap)
	err = h.O.Update(tt)
	if err != nil {
		c.AbortWithStatusJSON(484, err.Error())
		return
	}
	c.JSON(200, "Update successfully")
}
