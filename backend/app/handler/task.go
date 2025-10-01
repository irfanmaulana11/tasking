package handler

import (
	"be-tasking/app/model"
	"be-tasking/app/service/dto"
	"be-tasking/constanta"
	"be-tasking/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ApiHandler) GetTaskList(c *gin.Context) {
	var (
		limit       = c.Request.FormValue("limit")
		page        = c.Request.FormValue("page")
		search      = c.Request.FormValue("search")
		pageInt, _  = strconv.Atoi(page)
		limitInt, _ = strconv.Atoi(limit)
	)

	code, tasks, paging, err := h.taskService.GetTaskList(c, model.TableFilter{
		Search: search,
		Page:   pageInt,
		Limit:  limitInt,
	})
	if err != nil {
		helper.RespondJSON(c, code, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "get failed!",
			Error:   err.Error(),
		})
		return
	}

	helper.RespondJSON(c, code, helper.ResponseWrapper{
		Data:    tasks,
		Meta:    paging,
		Message: http.StatusText(http.StatusOK),
		Error:   nil,
	})
}

func (h *ApiHandler) CreateTask(c *gin.Context) {
	var (
		reqBody = dto.Task{}
	)

	if err := c.BindJSON(&reqBody); err != nil {
		helper.RespondJSON(c, http.StatusInternalServerError, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "can't bind request to struct",
			Error:   err.Error(),
		})
		return
	}

	// if err := reqBody.Validate(http.MethodPost); err != nil {
	// 	helper.RespondJSON(c, http.StatusBadRequest, helper.ResponseWrapper{
	// 		Data:    nil,
	// 		Meta:    nil,
	// 		Message: "error validation",
	// 		Error:   err.Error(),
	// 	})
	// 	return
	// }

	code, res, err := h.taskService.CreateTask(c, reqBody)
	if err != nil {
		helper.RespondJSON(c, code, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "create failed!",
			Error:   err.Error(),
		})
		return
	}

	helper.RespondJSON(c, code, helper.ResponseWrapper{
		Data:    res,
		Meta:    nil,
		Message: "create success!",
		Error:   nil,
	})
}

func (h *ApiHandler) UpdateTask(c *gin.Context) {
	var (
		reqBody = dto.Task{}
		id      = c.Param("id")
	)

	if err := c.BindJSON(&reqBody); err != nil {
		helper.RespondJSON(c, http.StatusInternalServerError, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "can't bind request to struct",
			Error:   err.Error(),
		})
		return
	}

	// if err := reqBody.Validate(http.MethodPost); err != nil {
	// 	helper.RespondJSON(c, http.StatusBadRequest, helper.ResponseWrapper{
	// 		Data:    nil,
	// 		Meta:    nil,
	// 		Message: "error validation",
	// 		Error:   err.Error(),
	// 	})
	// 	return
	// }

	reqBody.ID = id
	code, res, err := h.taskService.UpdateTask(c, reqBody)
	if err != nil {
		helper.RespondJSON(c, code, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "create failed!",
			Error:   err.Error(),
		})
		return
	}

	helper.RespondJSON(c, code, helper.ResponseWrapper{
		Data:    res,
		Meta:    nil,
		Message: "create success!",
		Error:   nil,
	})
}

func (h *ApiHandler) TaskInprogress(c *gin.Context) {
	h.TaskStatusUpdate(c, constanta.TaskStatusInProgress)
}
func (h *ApiHandler) TaskRevise(c *gin.Context) {
	h.TaskStatusUpdate(c, constanta.TaskStatusRevision)
}

func (h *ApiHandler) TaskApproved(c *gin.Context) {
	h.TaskStatusUpdate(c, constanta.TaskStatusApproved)
}
func (h *ApiHandler) TaskOveride(c *gin.Context) {
	h.TaskStatusUpdate(c, constanta.TaskStatusInProgress)
}

func (h *ApiHandler) TaskCompleted(c *gin.Context) {
	h.TaskStatusUpdate(c, constanta.TaskStatusCompleted)
}

func (h *ApiHandler) TaskStatusUpdate(c *gin.Context, status string) {
	var (
		reqBody = dto.TaskProgress{}
		id      = c.Param("id")
	)

	if err := c.BindJSON(&reqBody); err != nil {
		helper.RespondJSON(c, http.StatusInternalServerError, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "can't bind request to struct",
			Error:   err.Error(),
		})
		return
	}

	reqBody.ID = id
	reqBody.Status = status
	code, res, err := h.taskService.UpdateTaskStatus(c, reqBody)
	if err != nil {
		helper.RespondJSON(c, code, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "update failed!",
			Error:   err.Error(),
		})
		return
	}

	helper.RespondJSON(c, code, helper.ResponseWrapper{
		Data:    res,
		Meta:    nil,
		Message: "update success!",
		Error:   nil,
	})
}
