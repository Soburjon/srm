package repo

import "github.com/work/SRM/api/V1/models"

type Programmer interface {
	CreateTask(req models.CreateTaskRequest) error
	EditTask(req models.EditTaskRequest) error
	DeleteTask(taskID string) error
	UpdateTaskStatus(req models.UpdateTaskStatusRequest) error
	GetProjectTasks(projectID string) (models.GetProjectTasksResponse, error)
	GetTask(taskID string) (models.GetTaskResponse, error)
	GetMyTasks(req models.GetMyTaskRequest) (models.GetProjectTasksResponse, error)
	CreateCommit(req models.CreateCommitRequest) error
	EditCommit(req models.EditCommitRequest, userID string) error
	DeleteCommit(createdAt string, userID string) error
	GetCommits(taskID string) (models.GetCommitsResponse, error)
	CreateAttendance(req models.CreateAttendanceRequest) error
	UsersInProject(projectID string) (models.UsersInProjectResponse, error)
	UserRoleInProject(req models.UserRoleInProjectRequest) (string, error)
	MyProjects(userID string) (models.MyProjectsResponse, error)
}
