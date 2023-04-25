package postgres

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/work/SRM/api/V1/models"
	"github.com/work/SRM/pkg/sqls"
	"github.com/work/SRM/storage/repo"
)

type adminRepo struct {
	db *sqlx.DB
}

func NewAdminRepo(db *sqlx.DB) repo.Admin {
	return &adminRepo{
		db: db,
	}
}

func (a *adminRepo) CreateUser(req models.CreateUserRequest, role string) error {
	id := uuid.NewString()
	if _, err := a.db.Exec(sqls.CreateUsers, id, req.FullName, req.Avatar, role, req.PhoneNumber, req.Birthday, req.Password, req.Position); err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) EditUser(req models.EditUserRequest) error {
	_, err := a.db.Exec(sqls.EditUser, req.FullName, req.Avatar, req.PhoneNumber, req.Birthday, req.Password, req.Position, req.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) DeleteUser(req models.DeleteUserRequest) error {
	_, err := a.db.Exec(sqls.DeleteUser, req.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) CreateProject(req models.CreateProjectRequest, teamLeadID string) error {
	id := uuid.NewString()
	_, err := a.db.Exec(sqls.CreateProject, id, req.Name, req.StartDate, req.EndDate, teamLeadID, req.Attachments)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) EditProject(req models.EditProjectRequest, teamLeadID string) error {
	_, err := a.db.Exec(sqls.EditProject, req.Name, req.Status, req.StartDate, req.EndDate, teamLeadID, req.Attachments)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) DeleteProject(req models.DeleteProjectRequest) error {
	_, err := a.db.Exec(sqls.DeleteProject, req.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) UpdateProjectStatus(req models.UpdateProjectStatusRequest) error {
	_, err := a.db.Exec(sqls.UpdateProjectStatus, req.ProjectID, req.Status)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) GetUsersList(role string) (models.GetUserListResponse, error) {
	var avatar sql.NullString
	rows, err := a.db.Query(sqls.GetUsersList, role)
	if err != nil {
		return models.GetUserListResponse{}, err
	}

	res := models.GetUserListResponse{}

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.UserID, &user.FullName, &avatar, &user.PhoneNumber, &user.Birthday, &user.Position, &user.CreatedAt)
		if err != nil {
			return models.GetUserListResponse{}, err
		}

		user.Avatar = avatar.String

		res.Users = append(res.Users, user)
	}
	res.Count = uint32(len(res.Users))
	return res, nil
}

func (a *adminRepo) GetProjectsList() (models.GetProjectListsResponse, error) {
	var (
		startDate, endDate, attachments sql.NullString
	)
	rows, err := a.db.Query(sqls.GetProjectsList)
	if err != nil {
		return models.GetProjectListsResponse{}, err
	}

	res := models.GetProjectListsResponse{}

	for rows.Next() {
		project := models.Project{}
		err := rows.Scan(&project.ID, &project.Name, &project.Status, &startDate, &endDate, &project.TeamleadID, &attachments, &project.CreatedAt)
		if err != nil {
			return models.GetProjectListsResponse{}, err
		}

		project.StartDate = startDate.String
		project.EndDate = endDate.String
		project.Attachments = attachments.String

		res.Projects = append(res.Projects, project)
	}
	res.Count = uint32(len(res.Projects))
	return res, nil
}

func (a *adminRepo) GetMyProjects(userID string) (models.GetMyProjectsResponse, error) {
	var (
		startDate, endDate, attachments sql.NullString
	)
	rows, err := a.db.Query(sqls.GetMyProjects, userID)
	if err != nil {
		return models.GetMyProjectsResponse{}, err
	}

	res := models.GetMyProjectsResponse{}

	for rows.Next() {
		project := models.Project{}
		err := rows.Scan(&project.ID, &project.Name, &project.Status, &startDate, &endDate, &project.TeamleadID, &attachments, &project.CreatedAt)
		if err != nil {
			return models.GetMyProjectsResponse{}, err
		}

		project.StartDate = startDate.String
		project.EndDate = endDate.String
		project.Attachments = attachments.String

		res.Projects = append(res.Projects, project)
	}
	res.Count = uint32(len(res.Projects))
	return res, nil
}

func (a *adminRepo) AddPeopleProject(req models.AddPeopleProjectRequest) error {
	_, err := a.db.Exec(sqls.AddPeopleProject, req.ProgrammerID, req.ProjectID, req.Position)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepo) CheckTeamLead(req models.CheckTeamLeadRequest) (bool, error) {
	err := a.db.QueryRow(sqls.CheckTeamLead, req.UserID, req.ProjectID).Scan(&req.UserID)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, err
}

func (a *adminRepo) GetUserRole(userID string) (string, error) {
	var role string
	err := a.db.QueryRow(sqls.GetUserRole, userID).Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}

func (a *adminRepo) GetUser(userID string) (models.GetUser, error) {
	var avatar sql.NullString
	user := models.GetUser{}
	err := a.db.QueryRow(sqls.GetUser, userID).Scan(&user.UserID, &user.FullName, &avatar, &user.Role, &user.PhoneNumber, &user.Birthday, &user.Position, &user.CreatedAt)
	if err != nil {
		return models.GetUser{}, err
	}

	user.Avatar = avatar.String

	return user, nil
}

func (a *adminRepo) GetProject(projectID string) (models.GetProject, error) {
	var (
		startDate, endDate, attachments sql.NullString
	)
	project := models.GetProject{}
	err := a.db.QueryRow(sqls.GetProject, projectID).Scan(&project.ID, &project.Name, &project.Status, &startDate, &endDate, &project.TeamleadID, &attachments, &project.CreatedAt)
	if err != nil {
		return models.GetProject{}, err
	}

	project.StartDate = startDate.String
	project.EndDate = endDate.String
	project.Attachments = attachments.String

	return project, nil
}

func (a *adminRepo) GetAttendance(programmerID string) (models.GetAttendanceResponse, error) {
	rows, err := a.db.Query(sqls.GetAttendance, programmerID)
	if err != nil {
		return models.GetAttendanceResponse{}, err
	}

	res := models.GetAttendanceResponse{}

	for rows.Next() {
		attendance := models.Attendance{}
		err := rows.Scan(&attendance.Type, &attendance.CreatedAt)
		if err != nil {
			return models.GetAttendanceResponse{}, err
		}
		res.Attendances = append(res.Attendances, attendance)
	}
	res.Count = uint32(len(res.Attendances))
	return res, nil
}
