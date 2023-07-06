package task_template

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

// GetTaskTemplates swagger
// @Summary     Show all task templates
// @Description Get all task templates
// @Tags        task_template
// @Produce     json
// @Success     200 {array} task_template.TaskTemplate
// @Router      /task_template/ [get]
func (h *Handler) GetTaskTemplates(c *gin.Context) {
	ht, err := h.O.List()
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetTaskTemplates: ", err)
		return
	}
	h.L.Info().Println("GetTaskTemplates: success")
	c.JSON(200, Format(ht))
	return
}

// GetTaskTemplateById swagger
// @Summary     Show task templates
// @Description Get task templates by id
// @Tags        task_template
// @Produce     json
// @Param       id  path     int true "task template id"
// @Success     200 {object} task_template.TaskTemplate
// @Router      /task_template/{id} [get]
func (h *Handler) GetTaskTemplateById(c *gin.Context) {
	id := c.Param("id")
	IdInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetTaskTemplateById: ", err)
		return
	}
	ht, err := h.O.Find([]int32{int32(IdInt)})
	if len(ht) == 0 {
		util.Err(c, errors.New("empty task template"), 0)
		h.L.Error().Println("GetTaskTemplateById: ", "empty task template")
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetTaskTemplateById: ", err)
		return
	}
	h.L.Info().Println("GetTaskTemplateById: success")
	c.JSON(200, Format(ht)[0])
	return
}

// AddTaskTemplate swagger
// @Summary Create task templates
// @Tags    task_template
// @Accept  json
// @Produce json
// @Param   task_template body     []task_template.TaskTemplateCreate true "task template body"
// @Success 200           {array} task_template.TaskTemplate
// @Router  /task_template/ [post]
func (h *Handler) AddTaskTemplate(c *gin.Context) {
	entry := []*TaskTemplateCreate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddTaskTemplate: ", err)
		return
	}
	ht := CreateConvert(entry)
	ht, err := h.O.Create(ht)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddTaskTemplate: ", err)
		return
	}
	c.JSON(200, Format(ht))
}

// DeleteTaskTemplate swagger
// @Summary Delete task templates
// @Tags    task_template
// @Produce json
// @Param ids body []int true "task_template id"
// @Success 200 {string} string "delete successfully"
// @Router  /task_template/ [delete]
func (h *Handler) DeleteTaskTemplate(c *gin.Context) {
	entry := make([]int32, 0, 10)
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteTaskTemplate: ", err)
		return
	}
	ht, err := h.O.Find(entry)
	if len(ht) == 0 {
		util.Err(c, errors.New("empty task template"), 0)
		h.L.Error().Println("DeleteTaskTemplate: ", err)
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteTaskTemplate: ", err)
		return
	}
	err = h.O.Delete(ht)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteTaskTemplate: ", err)
		return
	}
	c.JSON(200, "delete successfully")
}
