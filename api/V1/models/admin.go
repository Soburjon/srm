package models

type CreateUserRequest struct {
	FullName    string `json:"full_name"`
	Avatar      string `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Password    string `json:"password"`
	Position    string `json:"position"`
}

type EditUserRequest struct {
	UserID      string `json:"user_id"`
	FullName    string `json:"full_name"`
	Avatar      string `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Password    string `json:"password"`
	Position    string `json:"position"`
}

type DeleteUserRequest struct {
	UserID string `json:"user_id"`
}

type CreateProjectRequest struct {
	Name        string `json:"name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Attachments string `json:"attachments"`
}

type EditProjectRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Attachments string `json:"attachments"`
}

type DeleteProjectRequest struct {
	ID string `json:"id"`
}

type UpdateProjectStatusRequest struct {
	ProjectID string `json:"project_id"`
	Status    string `json:"status"`
}

type GetUserListResponse struct {
	Count uint32 `json:"count"`
	Users []User `json:"users"`
}

type User struct {
	UserID      string `json:"user_id"`
	FullName    string `json:"full_name"`
	Avatar      string `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Position    string `json:"position"`
	CreatedAt   string `json:"created_at"`
}

type GetProjectListsResponse struct {
	Count    uint32    `json:"count"`
	Projects []Project `json:"projects"`
}

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	TeamleadID  string `json:"teamlead_id"`
	Attachments string `json:"attachments"`
	CreatedAt   string `json:"created_at"`
}

type GetMyProjectsResponse struct {
	Count    uint32    `json:"count"`
	Projects []Project `json:"projects"`
}

type AddPeopleProjectRequest struct {
	Position     string `json:"position"`
	ProjectID    string `json:"project_id"`
	ProgrammerID string `json:"programmer_id"`
}

type CheckTeamLeadRequest struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
}

type GetUser struct {
	UserID      string `json:"user_id"`
	FullName    string `json:"full_name"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Position    string `json:"position"`
	CreatedAt   string `json:"created_at"`
}

type GetProject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	TeamleadID  string `json:"teamlead_id"`
	Attachments string `json:"attachments"`
	CreatedAt   string `json:"created_at"`
}
