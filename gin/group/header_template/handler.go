package header_template

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

// GetheaderTemplates swagger
// @Summary     Show all header templates
// @Description Get all header templates
// @Tags        header_template
// @Produce     json
// @Success     200 {array} header_template.HeaderTemplate
// @Router      /header_template/ [get]
func (h *Handler) GetheaderTemplates(c *gin.Context) {
	ht, err := h.O.List()
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetheaderTemplates: ", err)
		return
	}
	h.L.Info().Println("GetheaderTemplates: success")
	c.JSON(200, ht)
	return
}

// GetHeaderTemplateById swagger
// @Summary     Show header templates
// @Description Get header templates by id
// @Tags        header_template
// @Produce     json
// @Param       id  path     int true "header template id"
// @Success     200 {object} header_template.HeaderTemplate
// @Router      /header_template/{id} [get]
func (h *Handler) GetHeaderTemplateById(c *gin.Context) {
	id := c.Param("id")
	IdInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetHeaderTemplateById: ", err)
		return
	}
	ht, err := h.O.Find([]int32{int32(IdInt)})
	if len(ht) == 0 {
		util.Err(c, errors.New("empty header template"), 0)
		h.L.Error().Println("GetHeaderTemplateById: ", "empty header template")
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetHeaderTemplateById: ", err)
		return
	}
	h.L.Info().Println("GetHeaderTemplateById: success")
	c.JSON(200, ht[0])
	return
}

// AddHeaderTemplate swagger
// @Summary Create header templates
// @Tags    header_template
// @Accept  json
// @Produce json
// @Param   header_template body     []header_template.HeaderTemplateCreate true "header template body"
// @Success 200           {array} header_template.HeaderTemplate
// @Router  /header_template/ [post]
func (h *Handler) AddHeaderTemplate(c *gin.Context) {
	entry := []*HeaderTemplateCreate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddHeaderTemplate: ", err)
		return
	}
	ht := CreateConvert(entry)
	ht, err := h.O.Create(ht)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddHeaderTemplate: ", err)
		return
	}
	c.JSON(200, ht)
}

// UpdateHeaderTemplate swagger
// @Summary Update header templates
// @Tags    header_template
// @Accept  json
// @Produce json
// @Param   header_template body     []header_template.HeaderTemplateUpdate true "modify header template body"
// @Success 200           {string} string "update successfully"
// @Router  /header_template/ [patch]
func (h *Handler) UpdateHeaderTemplate(c *gin.Context) {
	entry := []*HeaderTemplateUpdate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateHeaderTemplate: ", err)
		return
	}
	ids := make([]int32, 0, len(entry))
	uMap := make(map[int32]*HeaderTemplateUpdate)
	for _, item := range entry {
		ids = append(ids, item.ID)
		uMap[item.ID] = item
	}
	ht, err := h.O.Find(ids)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateHeaderTemplate: ", err)
		return
	}
	ht = UpdateConvert(ht, uMap)
	err = h.O.Update(ht)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateHeaderTemplate: ", err)
		return
	}
	c.JSON(200, "update successfully")
}

// DeleteHeaderTemplate swagger
// @Summary Delete header templates
// @Tags    header_template
// @Produce json
// @Param ids body []int true "header_template id"
// @Success 200 {string} string "delete successfully"
// @Router  /header_template/ [delete]
func (h *Handler) DeleteHeaderTemplate(c *gin.Context) {
	entry := make([]int32, 0, 10)
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteHeaderTemplate: ", err)
		return
	}
	ht, err := h.O.Find(entry)
	if len(ht) == 0 {
		util.Err(c, errors.New("empty header template"), 0)
		h.L.Error().Println("DeleteHeaderTemplate: ", err)
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteHeaderTemplate: ", err)
		return
	}
	err = h.O.Delete(ht)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteHeaderTemplate: ", err)
		return
	}
	c.JSON(200, "delete successfully")
}
