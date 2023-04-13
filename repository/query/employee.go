package query

const (
	AddEmployee = `
		INSERT INTO
			db_go_sql.employees
		(
			full_name,
			email,
			age,
			division
		)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`
)
