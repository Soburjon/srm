package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/storage"
)

type AdminService struct {
	storage storage.IStorage
}

func NewAdminService(db *sqlx.DB) *AdminService {
	return &AdminService{
		storage: storage.New(db),
	}
}

func (a *AdminService) CreateUser(req models.CreateUserRequest, role string) error {
	return a.storage.Admin().CreateUser(req, role)
}

func (a *AdminService) EditUser(req models.EditUserRequest) error {
	return a.storage.Admin().EditUser(req)
}

func (a *AdminService) DeleteUser(req models.DeleteUserRequest) error {
	return a.storage.Admin().DeleteUser(req)
}

func (a *AdminService) CreateProject(req models.CreateProjectRequest, teamLeadID string) error {
	return a.storage.Admin().CreateProject(req, teamLeadID)
}

func (a *AdminService) EditProject(req models.EditProjectRequest, teamLeadID string) error {
	return a.storage.Admin().EditProject(req, teamLeadID)
}

func (a *AdminService) DeleteProject(req models.DeleteProjectRequest) error {
	return a.storage.Admin().DeleteProject(req)
}

func (a *AdminService) UpdateProjectStatus(req models.UpdateProjectStatusRequest) error {
	return a.storage.Admin().UpdateProjectStatus(req)
}

func (a *AdminService) GetUsersList(role string) (models.GetUserListResponse, error) {
	return a.storage.Admin().GetUsersList(role)
}

func (a *AdminService) GetProjectsList() (models.GetProjectListsResponse, error) {
	return a.storage.Admin().GetProjectsList()
}

func (a *AdminService) GetMyProjects(userID string) (models.GetMyProjectsResponse, error) {
	return a.storage.Admin().GetMyProjects(userID)
}

func (a *AdminService) AddPeopleProject(req models.AddPeopleProjectRequest) error {
	return a.storage.Admin().AddPeopleProject(req)
}

func (a *AdminService) CheckTeamLead(req models.CheckTeamLeadRequest) (bool, error) {
	return a.storage.Admin().CheckTeamLead(req)
}

func (a *AdminService) GetUserRole(userID string) (string, error) {
	return a.storage.Admin().GetUserRole(userID)
}

func (a *AdminService) GetUser(userID string) (models.GetUser, error) {
	return a.storage.Admin().GetUser(userID)
}

func (a *AdminService) GetProject(projectID string) (models.GetProject, error) {
	return a.storage.Admin().GetProject(projectID)
}

func (a *AdminService) GetAttendance(programmerID string) (models.GetAttendanceResponse, error) {
	return a.storage.Admin().GetAttendance(programmerID)
}
