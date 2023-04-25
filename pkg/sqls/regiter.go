package sqls

const (
	Login = `
		SELECT id,
		       role 
		FROM users 
		WHERE phone_number = $1 AND password = $2 and deleted_at IS NULL;
`
)
