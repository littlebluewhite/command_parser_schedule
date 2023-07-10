package schedule

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

// GetSchedules swagger
// @Summary     Show all schedules
// @Description Get all schedules
// @Tags        schedule
// @Produce     json
// @Success     200 {array} schedule.Schedule
// @Router      /schedule/ [get]
func (h *Handler) GetSchedules(c *gin.Context) {
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

// GetScheduleById swagger
// @Summary     Show schedules
// @Description Get schedules by id
// @Tags        schedule
// @Produce     json
// @Param       id  path     int true "schedule id"
// @Success     200 {object} schedule.Schedule
// @Router      /schedule/{id} [get]
func (h *Handler) GetScheduleById(c *gin.Context) {
	id := c.Param("id")
	IdInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetScheduleById: ", err)
		return
	}
	tt, err := h.O.Find([]int32{int32(IdInt)})
	if len(tt) == 0 {
		util.Err(c, errors.New("empty schedule"), 0)
		h.L.Error().Println("GetScheduleById: ", "empty schedule")
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("GetScheduleById: ", err)
		return
	}
	result := Format(tt)
	h.L.Info().Println("GetScheduleById: success")
	c.JSON(200, result[0])
	return
}

// AddSchedule swagger
// @Summary Create schedules
// @Tags    schedule
// @Accept  json
// @Produce json
// @Param   schedule body     []schedule.ScheduleCreate true "schedule body"
// @Success 200           {array} schedule.Schedule
// @Router  /schedule/ [post]
func (h *Handler) AddSchedule(c *gin.Context) {
	entry := []*ScheduleCreate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddSchedule: ", err)
		return
	}
	tt := CreateConvert(entry)
	tt, err := h.O.Create(tt)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("AddSchedule: ", err)
		return
	}
	c.JSON(200, Format(tt))
}

// UpdateSchedule swagger
// @Summary Update schedules
// @Tags    schedule
// @Accept  json
// @Produce json
// @Param   schedule body     []schedule.ScheduleUpdate true "modify schedule body"
// @Success 200           {string} string "update successfully"
// @Router  /schedule/ [patch]
func (h *Handler) UpdateSchedule(c *gin.Context) {
	entry := []*ScheduleUpdate{nil}
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateSchedule: ", err)
		return
	}
	ids := make([]int32, 0, len(entry))
	uMap := make(map[int32]*ScheduleUpdate)
	for _, item := range entry {
		ids = append(ids, item.ID)
		uMap[item.ID] = item
	}
	tt, err := h.O.Find(ids)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateSchedule: ", err)
		return
	}
	tt = UpdateConvert(tt, uMap)
	err = h.O.Update(tt)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("UpdateSchedule: ", err)
		return
	}
	c.JSON(200, "update successfully")
}

// DeleteSchedule swagger
// @Summary Delete schedules
// @Tags    schedule
// @Produce json
// @Param ids body []int true "schedule id"
// @Success 200 {string} string "delete successfully"
// @Router  /schedule/ [delete]
func (h *Handler) DeleteSchedule(c *gin.Context) {
	entry := make([]int32, 0, 10)
	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteSchedule: ", err)
		return
	}
	tt, err := h.O.Find(entry)
	if len(tt) == 0 {
		util.Err(c, errors.New("empty schedule"), 0)
		h.L.Error().Println("DeleteSchedule: ", err)
		return
	}
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteSchedule: ", err)
		return
	}
	err = h.O.Delete(tt)
	if err != nil {
		util.Err(c, err, 0)
		h.L.Error().Println("DeleteSchedule: ", err)
		return
	}
	c.JSON(200, "delete successfully")
}
