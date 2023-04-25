package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/storage"
)

type ProgrammerService struct {
	storage storage.IStorage
}

func NewProgrammerService(db *sqlx.DB) *ProgrammerService {
	return &ProgrammerService{
		storage: storage.New(db),
	}
}

func (p *ProgrammerService) CreateTask(req models.CreateTaskRequest) error {
	return p.storage.Programmer().CreateTask(req)
}

func (p *ProgrammerService) EditTask(req models.EditTaskRequest) error {
	return p.storage.Programmer().EditTask(req)
}

func (p *ProgrammerService) DeleteTask(taskID string) error {
	return p.storage.Programmer().DeleteTask(taskID)
}

func (p *ProgrammerService) UpdateTaskStatus(req models.UpdateTaskStatusRequest) error {
	return p.storage.Programmer().UpdateTaskStatus(req)
}

func (p *ProgrammerService) GetProjectTasks(projectID string) (models.GetProjectTasksResponse, error) {
	return p.storage.Programmer().GetProjectTasks(projectID)
}

func (p *ProgrammerService) GetTask(taskID string) (models.GetTaskResponse, error) {
	return p.storage.Programmer().GetTask(taskID)
}

func (p *ProgrammerService) GetMyTasks(req models.GetMyTaskRequest) (models.GetProjectTasksResponse, error) {
	return p.storage.Programmer().GetMyTasks(req)
}

func (p *ProgrammerService) CreateCommit(req models.CreateCommitRequest) error {
	return p.storage.Programmer().CreateCommit(req)
}

func (p *ProgrammerService) EditCommit(req models.EditCommitRequest, userID string) error {
	return p.storage.Programmer().EditCommit(req, userID)
}

func (p *ProgrammerService) DeleteCommit(createdAt string, userID string) error {
	return p.storage.Programmer().DeleteCommit(createdAt, userID)
}

func (p *ProgrammerService) GetCommits(taskID string) (models.GetCommitsResponse, error) {
	return p.storage.Programmer().GetCommits(taskID)
}

func (p *ProgrammerService) CreateAttendance(req models.CreateAttendanceRequest) error {
	return p.storage.Programmer().CreateAttendance(req)
}

func (p *ProgrammerService) UsersInProject(projectID string) (models.UsersInProjectResponse, error) {
	return p.storage.Programmer().UsersInProject(projectID)
}

func (p *ProgrammerService) UserRoleInProject(req models.UserRoleInProjectRequest) (string, error) {
	return p.storage.Programmer().UserRoleInProject(req)
}

func (p *ProgrammerService) MyProjects(userID string) (models.MyProjectsResponse, error) {
	return p.storage.Programmer().MyProjects(userID)
}
