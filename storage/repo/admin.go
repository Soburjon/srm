package repo

import "github.com/work/SRM/api/V1/models"

type Admin interface {
	CreateUser(req models.CreateUserRequest, role string) error
	EditUser(req models.EditUserRequest) error
	DeleteUser(req models.DeleteUserRequest) error
	CreateProject(req models.CreateProjectRequest, teamLeadID string) error
	EditProject(req models.EditProjectRequest, teamLeadID string) error
	DeleteProject(req models.DeleteProjectRequest) error
	UpdateProjectStatus(req models.UpdateProjectStatusRequest) error
	GetUsersList(role string) (models.GetUserListResponse, error)
	GetProjectsList() (models.GetProjectListsResponse, error)
	GetMyProjects(userID string) (models.GetMyProjectsResponse, error)
	AddPeopleProject(req models.AddPeopleProjectRequest) error
	CheckTeamLead(req models.CheckTeamLeadRequest) (bool, error)
	GetUserRole(userID string) (string, error)
	GetUser(userID string) (models.GetUser, error)
	GetProject(projectID string) (models.GetProject, error)
	GetAttendance(programmerID string) (models.GetAttendanceResponse, error)
}
