package controllers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/pkg/utils"
	"net/http"
)

// CreateTask method create task
// @Security ApiKeyAuth
// @Description create task
// @Description Attachments mazil yoziladi
// @Description start va finish atlar "2001-02-26" farmatda yoziladi
// @Summary create task
// @Tags programmer
// @Accept json
// @Produce json
// @Param create_task body models.CreateTaskRequest true "create_task"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/create-task/ [POST]
func (a *Api) CreateTask(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.CreateTaskRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	role, err := a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: req.ProjectID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if role != "team_lead" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz team lead emassiz",
		})
	}

	err = a.ProgrammerService.CreateTask(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// EditTask method edit task
// @Security ApiKeyAuth
// @Description edit task
// @Description Attachments mazil yoziladi
// @Description start va finish atlar "2001-02-26" farmatda yoziladi
// @Summary edit task
// @Tags programmer
// @Accept json
// @Produce json
// @Param edit_task body models.EditTaskRequest true "edit_task"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/edit-task/ [PUT]
func (a *Api) EditTask(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.EditTaskRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	role, err := a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: req.ID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if role != "team_lead" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz team lead emassiz",
		})
	}

	err = a.ProgrammerService.EditTask(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// DeleteTask method delete task
// @Security ApiKeyAuth
// @Description delete task
// @Summary delete task
// @Tags programmer
// @Accept json
// @Produce json
// @Param task_id path string true "task_id"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/delete-task/{task_id}/ [DELETE]
func (a *Api) DeleteTask(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	taskID := c.Params("task_id")
	if _, err := uuid.Parse(taskID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	task, err := a.ProgrammerService.GetTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	role, err := a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: task.ProjectID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if role != "team_lead" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz team lead emassiz",
		})
	}

	err = a.ProgrammerService.DeleteTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// UpdateTaskStatus method update task status
// @Security ApiKeyAuth
// @Description update task status
// @Description status ga "in_process","done","finished" larni yozish mumkin
// @Summary update task status
// @Tags programmer
// @Accept json
// @Produce json
// @Param update_task_status body models.UpdateTaskStatusRequest true "update_task_status"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/update-task-status/ [PUT]
func (a *Api) UpdateTaskStatus(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.UpdateTaskStatusRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if req.Status != "in_process" && req.Status != "finished" && req.Status != "done" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "status da hatolik",
		})
	}

	task, err := a.ProgrammerService.GetTask(req.TaskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	role, err := a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: task.ProjectID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if role != "team_lead" && user.UserID.String() != task.ProgrammerID {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "sizga ruxsat yo'q",
		})
	} else if role != "team_lead" && user.UserID.String() == task.ProgrammerID && req.Status == "done" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "sizga ruxsat yo'q",
		})
	}

	err = a.ProgrammerService.UpdateTaskStatus(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// GetProjectTasks method get project tasks
// @Security ApiKeyAuth
// @Description get project tasks
// @Summary get project tasks
// @Tags programmer
// @Accept json
// @Produce json
// @Param project_id path string true "project_id"
// @Success 200 {object} models.GetProjectTasksResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/get-project-tasks/{project_id}/ [GET]
func (a *Api) GetProjectTasks(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	projectID := c.Params("project_id")
	if _, err := uuid.Parse(projectID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	_, err = a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: projectID,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "sizga ruxsat yo'q",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := a.ProgrammerService.GetProjectTasks(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// GetTask method get task
// @Security ApiKeyAuth
// @Description get task
// @Summary get task
// @Tags programmer
// @Accept json
// @Produce json
// @Param task_id path string true "task_id"
// @Success 200 {object} models.GetTaskResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/get-task/{task_id}/ [GET]
func (a *Api) GetTask(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	taskID := c.Params("task_id")
	if _, err := uuid.Parse(taskID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	task, err := a.ProgrammerService.GetTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: task.ProjectID,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "sizga ruxsat yo'q",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := a.ProgrammerService.GetTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// GetMyTasks method get my tasks
// @Security ApiKeyAuth
// @Description get my tasks
// @Summary get my tasks
// @Tags programmer
// @Accept json
// @Produce json
// @Param project_id path string true "project_id"
// @Success 200 {object} models.GetProjectTasksResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/get-my-tasks/{project_id}/ [GET]
func (a *Api) GetMyTasks(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	projectID := c.Params("project_id")

	_, err = a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: projectID,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz bu projectda yo'qsiz",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := a.ProgrammerService.GetMyTasks(models.GetMyTaskRequest{
		UserID:    user.UserID.String(),
		ProjectID: projectID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// CreateCommit method create commit
// @Security ApiKeyAuth
// @Description create commit
// @Summary create commit
// @Tags programmer
// @Accept json
// @Produce json
// @Param create_commit body models.CreateCommitRequest true "create_commit"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/create-commit/ [POST]
func (a *Api) CreateCommit(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.CreateCommitRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task, err := a.ProgrammerService.GetTask(req.TaskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: task.ProjectID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz bu projectda yo'qsiz",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = a.ProgrammerService.CreateCommit(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// EditCommit method edit commit
// @Security ApiKeyAuth
// @Description edit commit
// @Summary edit commit
// @Tags programmer
// @Accept json
// @Produce json
// @Param edit_commit body models.EditCommitRequest true "edit_commit"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/edit-commit/ [PUT]
func (a *Api) EditCommit(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.EditCommitRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = a.ProgrammerService.EditCommit(req, user.UserID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// DeleteCommit method delete commit
// @Security ApiKeyAuth
// @Description delete commit
// @Summary delete commit
// @Tags programmer
// @Accept json
// @Produce json
// @Param created_at path string true "created_at"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/delete-commit/{created_at}/ [DELETE]
func (a *Api) DeleteCommit(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdAt := c.Params("created_at")

	err = a.ProgrammerService.DeleteCommit(createdAt, user.UserID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// GetCommits method get commits
// @Security ApiKeyAuth
// @Description get commits
// @Summary get commits
// @Tags programmer
// @Accept json
// @Produce json
// @Param task_id path string true "task_id"
// @Success 200 {object} models.GetCommitsResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/get-commits/{task_id}/ [GET]
func (a *Api) GetCommits(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	taskID := c.Params("task_id")
	if _, err := uuid.Parse(taskID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	task, err := a.ProgrammerService.GetTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: task.ProjectID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz bu projectda yo'qsiz",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := a.ProgrammerService.GetCommits(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// CreateAttendance method create attendance
// @Security ApiKeyAuth
// @Description create attendance
// @Summary create attendance
// @Tags programmer
// @Accept json
// @Produce json
// @Param type path string true "type"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/create-attendance/{type}/ [POST]
func (a *Api) CreateAttendance(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	Type := c.Params("type")

	if Type != "came" && Type != "gone" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "type da hatolik",
		})
	}

	err = a.ProgrammerService.CreateAttendance(models.CreateAttendanceRequest{
		UserID: user.UserID.String(),
		Type:   Type,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// UsersInProject method users in the project
// @Security ApiKeyAuth
// @Description projectdagi userlar
// @Summary users in the project
// @Tags programmer
// @Accept json
// @Produce json
// @Param project_id path string true "project_id"
// @Success 200 {object} models.UsersInProjectResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/users-in-project/{project_id}/ [GET]
func (a *Api) UsersInProject(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	projectID := c.Params("project_id")

	if _, err := uuid.Parse(projectID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	_, err = a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: projectID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz bu projectda yo'qsiz",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := a.ProgrammerService.UsersInProject(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// MyProjects method get my projects
// @Security ApiKeyAuth
// @Description qatnashgan projectlari
// @Summary get my projects
// @Tags programmer
// @Accept json
// @Produce json
// @Success 200 {object} models.MyProjectsResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/get-my-projects/ [GET]
func (a *Api) MyProjects(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.ProgrammerService.MyProjects(user.UserID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// GetProjects method get project
// @Security ApiKeyAuth
// @Description get project
// @Summary get project
// @Tags programmer
// @Accept json
// @Produce json
// @Param project_id path string true "project_id"
// @Success 200 {object} models.GetProject
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/get-project/{project_id}/ [GET]
func (a *Api) GetProjects(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	projectID := c.Params("project_id")

	if _, err := uuid.Parse(projectID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	_, err = a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: projectID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz bu projectda yo'qsiz",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := a.AdminService.GetProject(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// AddPeoplesProject method add people project
// @Security ApiKeyAuth
// @Description team lead tamonidan projectga odam qo'shish
// @Description position ga "team_lead","programmer","intern" larni yozish mumkin
// @Summary add people project
// @Tags programmer
// @Accept json
// @Produce json
// @Param add_people_project body models.AddPeopleProjectRequest true "add_people_project"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /programmer/add-people-project/ [POST]
func (a *Api) AddPeoplesProject(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.AddPeopleProjectRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	role, err := a.ProgrammerService.UserRoleInProject(models.UserRoleInProjectRequest{
		UserID:    user.UserID.String(),
		ProjectID: req.ProjectID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if role != "team_lead" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "siz team lead emassiz",
		})
	}

	err = a.AdminService.AddPeopleProject(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}
