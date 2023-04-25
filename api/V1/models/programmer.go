package models

type CreateTaskRequest struct {
	ProjectID    string `json:"project_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	StartAt      string `json:"start_at"`
	FinishAt     string `json:"finish_at"`
	ProgrammerID string `json:"programmer_id"`
	Attachments  string `json:"attachments"`
}

type EditTaskRequest struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	StartAt      string `json:"start_at"`
	FinishAt     string `json:"finish_at"`
	ProgrammerID string `json:"programmer_id"`
	Attachments  string `json:"attachments"`
}

type UpdateTaskStatusRequest struct {
	TaskID string `json:"task_id"`
	Status string `json:"status"`
}

type GetTaskResponse struct {
	ID           string `json:"id"`
	ProjectID    string `json:"project_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	StartedAt    string `json:"started_at"`
	FinishedAt   string `json:"finished_at"`
	StartAt      string `json:"start_at"`
	FinishAt     string `json:"finish_at"`
	ProgrammerID string `json:"programmer_id"`
	Attachments  string `json:"attachments"`
	CreatedAt    string `json:"created_at"`
}

type GetProjectTasksResponse struct {
	Count uint32 `json:"count"`
	Tasks []GetTaskResponse
}

type GetMyTaskRequest struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
}

type CreateCommitRequest struct {
	TaskID string `json:"task_id"`
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

type EditCommitRequest struct {
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type GetCommitsResponse struct {
	Count   uint32   `json:"count"`
	Commits []Commit `json:"commits"`
}

type Commit struct {
	ProgrammerID string `json:"programmer_id"`
	Text         string `json:"text"`
	CreatedAt    string `json:"created_at"`
}

type CreateAttendanceRequest struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

type GetAttendanceResponse struct {
	Count       uint32       `json:"count"`
	Attendances []Attendance `json:"attendances"`
}

type Attendance struct {
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}

type UserRoleInProjectRequest struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
}

type UsersInProjectResponse struct {
	Count uint32   `json:"count"`
	Users []Worker `json:"users"`
}

type Worker struct {
	Position string `json:"position"`
	UserID   string `json:"user_id"`
}

type MyProjectsResponse struct {
	Count      uint32   `json:"count"`
	ProjectIds []string `json:"project_ids"`
}
