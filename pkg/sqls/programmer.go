package sqls

const (
	CreateTask = `
		INSERT INTO task (
		                  id,
						  project_id,
		                  title,
		                  description,
		                  start_at,
		                  finish_at,
		                  programmer_id,
		                  attachments)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
`

	EditTask = `
		UPDATE task SET 
		                  title = $2,
		                  description = $3,
		                  start_at = $4,
		                  finish_at = $5,
		                  programmer_id = $6,
		                  attachments = $7,
						  updated_at = CURRENT_TIMESTAMP 
		WHERE id = $1
`

	DeleteTask = `
		UPDATE task SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1
`

	StatusInProcess = `
		UPDATE task SET status = 'in_process', started_at = CURRENT_TIMESTAMP WHERE id = $1
`

	StatusFinished = `
		UPDATE task SET status = 'finished', finished_at = CURRENT_TIMESTAMP WHERE id = $1
`

	StatusDone = `
		UPDATE task SET status = 'done' WHERE id = $1
`

	GetProjectTasks = `
		SELECT
		    ID,
			PROJECT_ID,
		    TITLE,
		    DESCRIPTION,
		    START_AT,
		    FINISH_AT,
		    STATUS,
		    STARTED_AT,
		    FINISHED_AT,
		    PROGRAMMER_ID,
		    ATTACHMENTS,
		    CREATED_AT
		FROM task WHERE project_id = $1 and deleted_at IS NULL
`

	GetTask = `
		SELECT
		    ID,
			PROJECT_ID,
		    TITLE,
		    DESCRIPTION,
		    START_AT,
		    FINISH_AT,
		    STATUS,
		    STARTED_AT,
		    FINISHED_AT,
		    PROGRAMMER_ID,
		    ATTACHMENTS,
		    CREATED_AT
		FROM task WHERE id = $1 and deleted_at IS NULL
`

	GetMyTasks = `
		SELECT
		    ID,
			PROJECT_ID,
		    TITLE,
		    DESCRIPTION,
		    START_AT,
		    FINISH_AT,
		    STATUS,
		    STARTED_AT,
		    FINISHED_AT,
		    PROGRAMMER_ID,
		    ATTACHMENTS,
		    CREATED_AT
		FROM task WHERE project_id = $1 and programmer_id = $2 and deleted_at IS NULL
`

	CreateComment = `
		INSERT INTO comment (task_id, programmer_id, text) VALUES ($1,$2,$3)
`

	EditComment = `
		UPDATE comment SET text = $1, updated_at = CURRENT_TIMESTAMP  
		WHERE created_at = $2 AND programmer_id = $3
`

	DeleteComment = `
		UPDATE comment SET deleted_at = CURRENT_TIMESTAMP WHERE created_at = $1 AND programmer_id = $2
`

	GetComments = `
		SELECT programmer_id,text,created_at FROM comment WHERE task_id = $1 and deleted_at IS NULL
`
	CreateAttendince = `
		INSERT INTO attendince (user_id, type) VALUES ($1,$2)
`
	GetAttendinces = `
		SELECT type,created_at FROM attendince WHERE user_id = $1
`
	UserRoleInProject = `
		SELECT position FROM projects_people WHERE user_id = $1 AND project_id = $2
`

	UsersInProject = `
		SELECT position,user_id FROM projects_people WHERE project_id = $1
`

	MyProjects = `
		SELECT project_id FROM projects_people WHERE user_id = $1
`
)
