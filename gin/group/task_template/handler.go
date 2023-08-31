package task_template

import (
	"command_parser_schedule/entry/e_task_template"
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

// GetTaskTemplates swagger
// @Summary     Show all task templates
// @Description Get all task templates
// @Tags        task_template
// @Produce     json
// @Success     200 {array} e_task_template.TaskTemplate
// @Router      /task_template/ [get]
func (h *Handler) GetTaskTemplates(c *gin.Context) {
	ht, err := h.O.List()
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetTaskTemplates: ", err)
		return
	}
	h.L.Info().Println("GetTaskTemplates: success")
	c.JSON(200, e_task_template.Format(ht))
	return
}

// GetTaskTemplateById swagger
// @Summary     Show task templates
// @Description Get task templates by id
// @Tags        task_template
// @Produce     json
// @Param       id  path     int true "task template id"
// @Success     200 {object} e_task_template.TaskTemplate
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
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetTaskTemplateById: ", err)
		return
	}
	h.L.Info().Println("GetTaskTemplateById: success")
	c.JSON(200, e_task_template.Format(ht)[0])
	return
}

// AddTaskTemplate swagger
// @Summary Create task templates
// @Tags    task_template
// @Accept  json
// @Produce json
// @Param   task_template body     []e_task_template.TaskTemplateCreate true "task template body"
// @Success 200           {array} e_task_template.TaskTemplate
// @Router  /task_template/ [post]
func (h *Handler) AddTaskTemplate(c *gin.Context) {
	entry := []*e_task_template.TaskTemplateCreate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddTaskTemplate: ", err)
		return
	}
	tt, err := h.O.Create(entry)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddTaskTemplate: ", err)
		return
	}
	c.JSON(200, e_task_template.Format(tt))
}

// UpdateTaskTemplate swagger
// @Summary Update task templates
// @Tags    task_template
// @Accept  json
// @Produce json
// @Param   task_template body     []e_task_template.TaskTemplateUpdate true "modify task template body"
// @Success 200           {string} string "update successfully"
// @Router  /task_template/ [patch]
func (h *Handler) UpdateTaskTemplate(c *gin.Context) {
	entry := []*e_task_template.TaskTemplateUpdate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateTaskTemplate: ", err)
		return
	}
	err := h.O.Update(entry)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateTaskTemplate: ", err)
		return
	}
	c.JSON(200, "update successfully")
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
	err := h.O.Delete(entry)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteTaskTemplate: ", err)
		return
	}
	c.JSON(200, "delete successfully")
}
