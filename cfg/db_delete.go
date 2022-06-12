package cfg

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func DeleteAll(dbName string) {
	db, _ := sql.Open("sqlite3", "./"+dbName)
	stmt := Prepare(db, "DELETE FROM diru")
	stmt.Exec()
	stmt.Close()
	defer db.Close()
}

func DeleteOne(dbName string, gid int) {
	db, _ := sql.Open("sqlite3", "./"+dbName)
	stmt := Prepare(db, "DELETE FROM diru WHERE id = ?")
	stmt.Exec(gid)
	stmt.Close()
	defer db.Close()
}
