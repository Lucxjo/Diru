package cfg

import "database/sql"

func Prepare(db *sql.DB, query string) *sql.Stmt {
	stmt, stmtErr := db.Prepare(query)

	if stmtErr != nil {
		panic(stmtErr)
	}

	return stmt
}
