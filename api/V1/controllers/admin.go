package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/pkg/utils"
	"net/http"
	"strings"
	"time"
)

// CreateAdmin method create admin
// @Security ApiKeyAuth
// @Description create admin
// @Description avatar mazil yoziladi
// @Description birthday "2001-02-26" farmatda yoziladi
// @Summary create admin
// @Tags admin
// @Accept json
// @Produce json
// @Param create_admin body models.CreateUserRequest true "create_admin"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/create-admin/ [POST]
func (a *Api) CreateAdmin(c *fiber.Ctx) error {
	req := models.CreateUserRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !strings.Contains(req.PhoneNumber, "+") ||
		len(req.PhoneNumber) != 13 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "phone number is not correctly filled",
		})
	}

	err = a.AdminService.CreateUser(req, "admin")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// EditAdmin method edit admin
// @Security ApiKeyAuth
// @Description edit admin
// @Description avatar mazil yoziladi
// @Description birthday "2001-02-26" farmatda yoziladi
// @Summary edit admin
// @Tags admin
// @Accept json
// @Produce json
// @Param edit_admin body models.EditUserRequest true "edit_admin"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/edit-admin/ [PUT]
func (a *Api) EditAdmin(c *fiber.Ctx) error {
	req := models.EditUserRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !strings.Contains(req.PhoneNumber, "+") ||
		len(req.PhoneNumber) != 13 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "phone number is not correctly filled",
		})
	}

	err = a.AdminService.EditUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// DeleteAdmin method delete admin
// @Security ApiKeyAuth
// @Description delete admin
// @Summary delete admin
// @Tags admin
// @Accept json
// @Produce json
// @Param admin_id path string true "admin_id"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/delete-admin/{admin_id}/ [DELETE]
func (a *Api) DeleteAdmin(c *fiber.Ctx) error {
	req := models.DeleteUserRequest{}
	req.UserID = c.Params("admin_id")

	if _, err := uuid.Parse(req.UserID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	err := a.AdminService.DeleteUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// CreateProgrammer method create programmer
// @Security ApiKeyAuth
// @Description create programmer
// @Description avatar mazil yoziladi
// @Description birthday "2001-02-26" farmatda yoziladi
// @Summary create programmer
// @Tags admin
// @Accept json
// @Produce json
// @Param create_programmer body models.CreateUserRequest true "create_programmer"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/create-programmer/ [POST]
func (a *Api) CreateProgrammer(c *fiber.Ctx) error {
	req := models.CreateUserRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !strings.Contains(req.PhoneNumber, "+") ||
		len(req.PhoneNumber) != 13 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "phone number is not correctly filled",
		})
	}

	err = a.AdminService.CreateUser(req, "user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// EditProgrammer method edit programmer
// @Security ApiKeyAuth
// @Description edit programmer
// @Description avatar mazil yoziladi
// @Description birthday "2001-02-26" farmatda yoziladi
// @Summary edit programmer
// @Tags admin
// @Accept json
// @Produce json
// @Param edit_programmer body models.EditUserRequest true "edit_programmer"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/edit-programmer/ [PUT]
func (a *Api) EditProgrammer(c *fiber.Ctx) error {
	req := models.EditUserRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !strings.Contains(req.PhoneNumber, "+") ||
		len(req.PhoneNumber) != 13 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "phone number is not correctly filled",
		})
	}

	err = a.AdminService.EditUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// DeleteProgrammer method delete programmer
// @Security ApiKeyAuth
// @Description delete programmer
// @Summary delete programmer
// @Tags admin
// @Accept json
// @Produce json
// @Param programmer_id path string true "programmer_id"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/delete-programmer/{programmer_id}/ [DELETE]
func (a *Api) DeleteProgrammer(c *fiber.Ctx) error {
	req := models.DeleteUserRequest{}
	req.UserID = c.Params("programmer_id")

	if _, err := uuid.Parse(req.UserID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	err := a.AdminService.DeleteUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// CreateProject method create project
// @Security ApiKeyAuth
// @Description create project
// @Description attachments mazil yoziladi
// @Description start va end date "2001-02-26" farmatda yoziladi
// @Summary create project
// @Tags admin
// @Accept json
// @Produce json
// @Param create_project body models.CreateProjectRequest true "create_project"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/create-project/ [POST]
func (a *Api) CreateProject(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.CreateProjectRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err1 := time.Parse("2006-01-02", req.StartDate)
	_, err2 := time.Parse("2006-01-02", req.EndDate)

	if err1 != nil || err2 != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "start date yoki end date yozishda hatolik",
		})
	}

	err = a.AdminService.CreateProject(req, user.UserID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// EditProject method edit project
// @Security ApiKeyAuth
// @Description edit project
// @Description attachments mazil yoziladi
// @Description start va end date "2001-02-26" farmatda yoziladi
// @Summary edit project
// @Tags admin
// @Accept json
// @Produce json
// @Param edit_project body models.EditProjectRequest true "edit_project"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/edit-project/ [PUT]
func (a *Api) EditProject(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.EditProjectRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err1 := time.Parse("2006-01-02", req.StartDate)
	_, err2 := time.Parse("2006-01-02", req.EndDate)

	if err1 != nil || err2 != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "start date yoki end date yozishda hatolik",
		})
	}

	err = a.AdminService.EditProject(req, user.UserID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// DeleteProject method delete project
// @Security ApiKeyAuth
// @Description delete project
// @Summary delete project
// @Tags admin
// @Accept json
// @Produce json
// @Param project_id path string true "project_id"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/delete-project/{project_id}/ [DELETE]
func (a *Api) DeleteProject(c *fiber.Ctx) error {

	req := models.DeleteProjectRequest{}
	req.ID = c.Params("project_id")

	if _, err := uuid.Parse(req.ID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	err := a.AdminService.DeleteProject(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// UpdateProjectStatus method update project status
// @Security ApiKeyAuth
// @Description update project status
// @Description status ga "in_process","new","finished" larni yozish mumkin
// @Summary update project status
// @Tags admin
// @Accept json
// @Produce json
// @Param update_project_status body models.UpdateProjectStatusRequest true "update_project_status"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/update-project-status/ [PUT]
func (a *Api) UpdateProjectStatus(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := models.UpdateProjectStatusRequest{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if req.Status != "in_process" && req.Status != "new" && req.Status != "finished" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "status da hatolik",
		})
	}

	check, err := a.AdminService.CheckTeamLead(models.CheckTeamLeadRequest{
		UserID:    user.UserID.String(),
		ProjectID: req.ProjectID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if !check {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "bu adminni vazifasi emas",
		})
	}

	err = a.AdminService.UpdateProjectStatus(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result": models.SuccessMessage{
			Success: true,
		},
	})
}

// GetAminList method get admin list
// @Security ApiKeyAuth
// @Description get admin list
// @Summary get admin list
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserListResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/get-admin-list/ [GET]
func (a *Api) GetAminList(c *fiber.Ctx) error {

	res, err := a.AdminService.GetUsersList("admin")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// GetProgrammerList method get programmer list
// @Security ApiKeyAuth
// @Description get programmer list
// @Summary get programmer list
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserListResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/get-programmer-list/ [GET]
func (a *Api) GetProgrammerList(c *fiber.Ctx) error {

	res, err := a.AdminService.GetUsersList("user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// GetProjectsList method get project list
// @Security ApiKeyAuth
// @Description get project list
// @Summary get project list
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} models.GetProjectListsResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/get-project-list/ [GET]
func (a *Api) GetProjectsList(c *fiber.Ctx) error {
	res, err := a.AdminService.GetProjectsList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// GetMyProjects method get my projects
// @Security ApiKeyAuth
// @Description Admin team lead bolgan projectlari
// @Summary get my projects
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} models.GetMyProjectsResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/get-my-projects/ [GET]
func (a *Api) GetMyProjects(c *fiber.Ctx) error {
	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.AdminService.GetMyProjects(user.UserID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// AddPeopleProject method add people project
// @Security ApiKeyAuth
// @Description team lead tamonidan projectga odam qo'shish
// @Description position ga "team_lead","programmer","intern" larni yozish mumkin
// @Summary add people project
// @Tags admin
// @Accept json
// @Produce json
// @Param add_people_project body models.AddPeopleProjectRequest true "add_people_project"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/add-people-project/ [POST]
func (a *Api) AddPeopleProject(c *fiber.Ctx) error {
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

	if req.Position != "team_lead" && req.Position != "programmer" && req.ProjectID != "intern" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "position da hatolik",
		})
	}

	role, err := a.AdminService.GetUserRole(req.ProgrammerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if role != "user" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "qo'shayot odam programmer emas",
		})
	}

	check, err := a.AdminService.CheckTeamLead(models.CheckTeamLeadRequest{
		UserID:    user.UserID.String(),
		ProjectID: req.ProjectID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if !check {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "bu adminni vazifasi emas",
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

// GetUser method get user
// @Security ApiKeyAuth
// @Description user haqida ma'lumotlari
// @Summary get user
// @Tags admin
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} models.GetUser
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/get-user/{user_id}/ [GET]
func (a *Api) GetUser(c *fiber.Ctx) error {

	userID := c.Params("user_id")

	if _, err := uuid.Parse(userID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	res, err := a.AdminService.GetUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// GetProject method get project
// @Security ApiKeyAuth
// @Description project haqida ma'lumotlari
// @Summary get project
// @Tags admin
// @Accept json
// @Produce json
// @Param project_id path string true "project_id"
// @Success 200 {object} models.GetProject
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/get-project/{project_id}/ [GET]
func (a *Api) GetProject(c *fiber.Ctx) error {

	projectID := c.Params("project_id")

	if _, err := uuid.Parse(projectID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	res, err := a.AdminService.GetProject(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}

// GetAttendance method get attendance
// @Security ApiKeyAuth
// @Description get attendance
// @Summary get attendance
// @Tags admin
// @Accept json
// @Produce json
// @Param programmer_id path string true "programmer_id"
// @Success 200 {object} models.GetAttendanceResponse
// @Failure 404 {object} models.StandardErrorModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /admin/get-attendance/{programmer_id}/ [GET]
func (a *Api) GetAttendance(c *fiber.Ctx) error {

	programmerID := c.Params("programmer_id")

	if _, err := uuid.Parse(programmerID); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "id da hatolik bor",
		})
	}

	res, err := a.AdminService.GetAttendance(programmerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "Application/json")
	return c.Status(http.StatusOK).JSON(res)
}
