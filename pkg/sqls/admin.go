package sqls

const (
	CreateUsers = `
		INSERT INTO users (
		                   id, 
		                   full_name,
		                   avatar,
		                   role, 
		                   phone_number, 
		                   birthday, 
		                   password, 
		                   position)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8);
`

	EditUser = `
		UPDATE users SET 
				full_name = $1, 
				avatar = $2, 
				phone_number = $3, 
				birthday = $4, 
				password = $5, 
				position= $6,
				updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
`

	DeleteUser = `
		UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1
`

	CreateProject = `
		INSERT INTO project (
		                     id, 
		                     name,
		                     start_date, 
		                     end_date, 
		                     teamlead_id, 
		                     attachments )
		values ($1,$2,$3,$4,$5,$6)
`

	EditProject = `
		UPDATE project SET
				name = $1,
				status = $2,
				start_date = $3,
				end_date = $4,
				teamlead_id = $5,
				attachments = $6,
				updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
`

	DeleteProject = `
		UPDATE project SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1
`

	UpdateProjectStatus = `
		UPDATE project SET status = $2 WHERE id = $1
`

	GetUsersList = `
		SELECT  
		    id,
		    full_name,
		    avatar,
		    phone_number,
		    birthday,
		    position,
		    created_at
		FROM users where role = $1 and deleted_at IS NULL
`

	GetProjectsList = `
		SELECT
		    id,
		    name,
		    status,
		    start_date,
		    end_date,
		    teamlead_id,
		    attachments,
		    created_at
		FROM project WHERE deleted_at IS NULL
`

	GetMyProjects = `
		SELECT
		    id,
		    name,
		    status,
		    start_date,
		    end_date,
		    teamlead_id,
		    attachments,
		    created_at
		FROM project WHERE teamlead_id = $1 and deleted_at IS NULL
`

	AddPeopleProject = `
		INSERT INTO projects_people (user_id, project_id, position) VALUES ($1,$2,$3)
`

	CheckTeamLead = `
    	SELECT id FROM project WHERE teamlead_id = $1 AND id = $2 and deleted_at IS NULL
`

	GetUserRole = `
		SELECT role FROM users where id = $1 and deleted_at IS NULL
`

	GetUser = `
		SELECT  
		    id,
		    full_name,
		    avatar,
			role,
		    phone_number,
		    birthday,
		    position,
		    created_at
		FROM users where id = $1 and deleted_at IS NULL
`

	GetProject = `
		SELECT
		    id,
		    name,
		    status,
		    start_date,
		    end_date,
		    teamlead_id,
		    attachments,
		    created_at
		FROM project WHERE id = $1 and deleted_at IS NULL
`
	GetAttendance = `
		SELECT type,created_at FROM attendince WHERE user_id = $1
`
)
