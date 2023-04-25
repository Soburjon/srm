package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/pkg/sqls"
	"github.com/work/SRM/storage/repo"
)

type programmerRepo struct {
	db *sqlx.DB
}

func NewProgrammerRepo(db *sqlx.DB) repo.Programmer {
	return &programmerRepo{db: db}
}

func (p *programmerRepo) CreateTask(req models.CreateTaskRequest) error {
	id := uuid.NewString()
	_, err := p.db.Exec(sqls.CreateTask, id, req.ProjectID, req.Title, req.Description, req.StartAt, req.FinishAt, req.ProgrammerID, req.Attachments)
	if err != nil {
		return err
	}
	return nil
}

func (p *programmerRepo) EditTask(req models.EditTaskRequest) error {
	_, err := p.db.Exec(sqls.EditTask, req.ID, req.Title, req.Description, req.StartAt, req.FinishAt, req.ProgrammerID, req.Attachments)
	if err != nil {
		return err
	}
	return nil
}

func (p *programmerRepo) DeleteTask(taskID string) error {
	_, err := p.db.Exec(sqls.DeleteTask, taskID)
	if err != nil {
		return err
	}
	return nil
}

func (p *programmerRepo) UpdateTaskStatus(req models.UpdateTaskStatusRequest) error {
	var query string

	if req.Status == "in_process" {
		query = sqls.StatusInProcess
	} else if req.Status == "finished" {
		query = sqls.StatusFinished
	} else {
		query = sqls.StatusDone
	}

	_, err := p.db.Exec(query, req.TaskID)
	if err != nil {
		return err
	}

	return nil
}

func (p *programmerRepo) GetProjectTasks(projectID string) (models.GetProjectTasksResponse, error) {
	var (
		startAt, startedAt, finishAt, finishedAt, attachments sql.NullString
	)
	rows, err := p.db.Query(sqls.GetProjectTasks, projectID)
	if err != nil {
		return models.GetProjectTasksResponse{}, err
	}
	res := models.GetProjectTasksResponse{}
	res.Tasks = make([]models.GetTaskResponse, 0)
	for rows.Next() {
		task := models.GetTaskResponse{}
		err := rows.Scan(&task.ID, &task.ProjectID, &task.Title, &task.Description, &startAt, &finishAt, &task.Status, &startedAt, &finishedAt, &task.ProgrammerID, &attachments, &task.CreatedAt)
		if err != nil {
			return models.GetProjectTasksResponse{}, err
		}

		task.StartAt = startAt.String
		task.StartedAt = startedAt.String
		task.FinishAt = finishAt.String
		task.FinishedAt = finishedAt.String
		task.Attachments = attachments.String

		res.Tasks = append(res.Tasks, task)
	}
	res.Count = uint32(len(res.Tasks))
	return res, nil
}

func (p *programmerRepo) GetTask(taskID string) (models.GetTaskResponse, error) {
	var (
		startAt, startedAt, finishAt, finishedAt, attachments sql.NullString
	)
	task := models.GetTaskResponse{}
	err := p.db.QueryRow(sqls.GetTask, taskID).Scan(&task.ID, &task.ProjectID, &task.Title, &task.Description, &startAt, &finishAt, &task.Status, &startedAt, &finishedAt, &task.ProgrammerID, &attachments, &task.CreatedAt)
	if err != nil {
		return models.GetTaskResponse{}, err
	}

	task.StartAt = startAt.String
	task.StartedAt = startedAt.String
	task.FinishAt = finishAt.String
	task.FinishedAt = finishedAt.String
	task.Attachments = attachments.String

	return task, nil
}

func (p *programmerRepo) GetMyTasks(req models.GetMyTaskRequest) (models.GetProjectTasksResponse, error) {
	var (
		startAt, startedAt, finishAt, finishedAt, attachments sql.NullString
	)
	rows, err := p.db.Query(sqls.GetMyTasks, req.ProjectID, req.UserID)
	if err != nil {
		return models.GetProjectTasksResponse{}, err
	}
	res := models.GetProjectTasksResponse{}
	for rows.Next() {
		task := models.GetTaskResponse{}
		err := rows.Scan(&task.ID, &task.ProjectID, &task.Title, &task.Description, &startAt, &finishAt, &task.Status, &startedAt, &finishedAt, &task.ProgrammerID, &attachments, &task.CreatedAt)
		if err != nil {
			return models.GetProjectTasksResponse{}, err
		}

		task.StartAt = startAt.String
		task.StartedAt = startedAt.String
		task.FinishAt = finishAt.String
		task.FinishedAt = finishedAt.String
		task.Attachments = attachments.String

		res.Tasks = append(res.Tasks, task)
	}
	res.Count = uint32(len(res.Tasks))
	return res, nil
}

func (p *programmerRepo) CreateCommit(req models.CreateCommitRequest) error {
	_, err := p.db.Exec(sqls.CreateComment, req.TaskID, req.UserID, req.Text)
	if err != nil {
		return err
	}
	return nil
}

func (p *programmerRepo) EditCommit(req models.EditCommitRequest, userID string) error {
	_, err := p.db.Exec(sqls.EditComment, req.Text, req.CreatedAt, userID)
	if err != nil {
		return err
	}
	return nil
}

func (p *programmerRepo) DeleteCommit(createdAt string, userID string) error {
	_, err := p.db.Exec(sqls.DeleteComment, createdAt, userID)
	if err != nil {
		return err
	}
	return nil
}

func (p *programmerRepo) GetCommits(taskID string) (models.GetCommitsResponse, error) {
	rows, err := p.db.Query(sqls.GetComments, taskID)
	if err != nil {
		return models.GetCommitsResponse{}, err
	}
	res := models.GetCommitsResponse{}
	for rows.Next() {
		comment := models.Commit{}
		err := rows.Scan(&comment.ProgrammerID, &comment.Text, &comment.CreatedAt)
		if err != nil {
			return models.GetCommitsResponse{}, err
		}
		res.Commits = append(res.Commits, comment)
	}
	res.Count = uint32(len(res.Commits))
	return res, nil
}

func (p *programmerRepo) CreateAttendance(req models.CreateAttendanceRequest) error {
	_, err := p.db.Exec(sqls.CreateAttendince, req.UserID, req.Type)
	if err != nil {
		return err
	}
	return nil
}

func (p *programmerRepo) UsersInProject(projectID string) (models.UsersInProjectResponse, error) {
	rows, err := p.db.Query(sqls.UsersInProject, projectID)
	if err != nil {
		return models.UsersInProjectResponse{}, err
	}

	res := models.UsersInProjectResponse{}

	for rows.Next() {
		user := models.Worker{}
		if err := rows.Scan(&user.Position, &user.UserID); err != nil {
			return models.UsersInProjectResponse{}, err
		}
		res.Users = append(res.Users, user)
	}
	res.Count = uint32(len(res.Users))
	return res, nil
}

func (p *programmerRepo) UserRoleInProject(req models.UserRoleInProjectRequest) (string, error) {
	var position string
	err := p.db.QueryRow(sqls.UserRoleInProject, req.UserID, req.ProjectID).Scan(&position)
	if err != nil {
		return "", err
	}
	return position, nil
}

func (p *programmerRepo) MyProjects(userID string) (models.MyProjectsResponse, error) {
	rows, err := p.db.Query(sqls.MyProjects, userID)
	if err != nil {
		return models.MyProjectsResponse{}, err
	}

	res := models.MyProjectsResponse{}

	for rows.Next() {
		projectID := ""
		if err := rows.Scan(&projectID); err != nil {
			return models.MyProjectsResponse{}, err
		}
		res.ProjectIds = append(res.ProjectIds, projectID)
	}
	res.Count = uint32(len(res.ProjectIds))
	return res, nil
}
