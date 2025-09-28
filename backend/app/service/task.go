package service

import (
	"be-tasking/app/model"
	"be-tasking/app/service/dto"
	"be-tasking/constanta"
	"be-tasking/helper"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func (s *taskService) CreateTask(c *gin.Context, req dto.Task) (int, *dto.Task, error) {
	var (
		token    = helper.GetTokenHeader(c)
		claim, _ = helper.GetTokenClaims(token)
	)

	task := model.Task{
		ID:             ulid.Make().String(),
		Title:          req.Title,
		Description:    req.Description,
		Assignee:       req.AssignedLeader,
		AssignedLeader: req.AssignedLeader,
		Status:         constanta.TaskStatusSubmitted,
		Progress:       0,
		ProgressBy:     &claim.UserName,
		//Deadline:       &time.Time{},
		CreatedBy: claim.UserName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.MySQL.CreateTask(c, task); err != nil {
		log.Println("error CreateTask : ", err)
		return http.StatusInternalServerError, nil, err
	}

	if err := s.MySQL.CreateTaskHistory(c, model.TaskHistory{
		TaskID:    task.ID,
		ActionBy:  *task.ProgressBy,
		Action:    task.Status,
		Note:      "",
		CreatedAt: time.Now(),
	}); err != nil {
		log.Println("error CreateTaskHistory : ", err)
		// keep oke when failed save history
	}

	return http.StatusOK, &dto.Task{
		ID:             task.ID,
		Title:          task.Title,
		Description:    task.Description,
		Assignee:       task.Assignee,
		AssignedLeader: task.AssignedLeader,
		Status:         task.Status,
		Progress:       task.Progress,
		ProgressBy:     task.ProgressBy,
		// Deadline:       task.Deadline,
		CreatedBy: task.CreatedBy,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (s *taskService) UpdateTask(c *gin.Context, req dto.Task) (int, *dto.Task, error) {
	var (
		token    = helper.GetTokenHeader(c)
		claim, _ = helper.GetTokenClaims(token)
	)

	task, err := s.MySQL.GetTaskByID(c, req.ID)
	if err != nil {
		log.Println("error GetTaskByID : ", err)
		return http.StatusInternalServerError, nil, err
	}

	if task.Status != constanta.TaskStatusRevision {
		return http.StatusBadRequest, nil, fmt.Errorf("the task no need any revision")
	}

	updates := map[string]interface{}{
		"title":       req.Title,
		"description": req.Description,
		"updated_at":  time.Now(),
		"assignee":    task.AssignedLeader,

		// "deadline" :
	}

	if err := s.MySQL.UpdateTask(c, req.ID, updates); err != nil {
		log.Println("error UpdateTask : ", err)
		return http.StatusInternalServerError, nil, err
	}

	if err := s.MySQL.CreateTaskHistory(c, model.TaskHistory{
		TaskID:    req.ID,
		ActionBy:  claim.UserName,
		Action:    constanta.TaskStatusUpdated,
		Note:      "",
		CreatedAt: time.Now(),
	}); err != nil {
		log.Println("error CreateTaskHistory : ", err)
		// keep oke when failed save history
	}

	return http.StatusInternalServerError, &req, nil
}

func (s *taskService) UpdateTaskStatus(c *gin.Context, req dto.TaskProgress) (int, *dto.TaskProgress, error) {
	var (
		token    = helper.GetTokenHeader(c)
		claim, _ = helper.GetTokenClaims(token)
		assignee string
	)

	task, err := s.MySQL.GetTaskByID(c, req.ID)
	if err != nil {
		log.Println("error GetTaskByID : ", err)
		return http.StatusInternalServerError, nil, err
	}

	progress := task.Progress

	switch req.Status {
	case constanta.TaskStatusApproved, constanta.TaskStatusRevision:
		assignee = task.CreatedBy
	case constanta.TaskStatusInProgress:
		if task.Status != constanta.TaskStatusApproved {
			return http.StatusBadRequest, nil, fmt.Errorf("the task must be approved by leader")
		}
		progress = req.Progress
		assignee = task.CreatedBy
		if progress >= 100 && claim.Role == constanta.RoleTypePelaksana {
			assignee = task.AssignedLeader
		}
	default:
		assignee = task.CreatedBy
	}

	updates := map[string]interface{}{
		"status":     req.Status,
		"assignee":   assignee,
		"updated_at": time.Now(),
		"progress":   progress,
	}

	if err := s.MySQL.UpdateTask(c, req.ID, updates); err != nil {
		log.Println("error UpdateTask : ", err)
		return http.StatusInternalServerError, nil, err
	}

	if err := s.MySQL.CreateTaskHistory(c, model.TaskHistory{
		TaskID:    req.ID,
		ActionBy:  claim.UserName,
		Action:    req.Status,
		Note:      req.Note,
		CreatedAt: time.Now(),
	}); err != nil {
		log.Println("error CreateTaskHistory : ", err)
		// keep oke when failed save history
	}

	return http.StatusOK, &req, nil
}

func (s *taskService) GetTaskList(c *gin.Context, filter model.TableFilter) (int, []dto.Task, interface{}, error) {
	var (
		err      error
		tasks    = []model.Task{}
		res      = []dto.Task{}
		total    int
		token    = helper.GetTokenHeader(c)
		claim, _ = helper.GetTokenClaims(token)
	)

	if filter.Page == 0 {
		filter.Page = 1
	}

	if filter.Limit == 0 {
		filter.Limit = 5
	}

	filter.Role = claim.Role

	tasks, total, err = s.MySQL.GetTaskList(c, filter)
	if err != nil {
		return http.StatusInternalServerError, res, nil, fmt.Errorf("can't find data : %s", err.Error())
	}

	for _, task := range tasks {

		history, err := s.MySQL.GetTaskHistory(c, task.ID)
		if err != nil {
			log.Println("error GetTaskHistory : ", err)
		}

		histories := []dto.TaskHistory{}
		for _, h := range history {
			histories = append(histories, dto.TaskHistory{
				ActionBy:  h.ActionBy,
				Action:    h.Action,
				Note:      h.Note,
				CreatedAt: h.CreatedAt,
			})
		}

		res = append(res, dto.Task{
			ID:             task.ID,
			Title:          task.Title,
			Description:    task.Description,
			Assignee:       task.Assignee,
			AssignedLeader: task.AssignedLeader,
			Status:         task.Status,
			Progress:       task.Progress,
			ProgressBy:     task.ProgressBy,
			// Deadline:       task.Deadline,
			CreatedBy:   task.CreatedBy,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			TaskHistory: histories,
		})
	}

	paging := helper.BuildPagination(filter.Page, len(res), total)

	return http.StatusOK, res, paging, nil
}
