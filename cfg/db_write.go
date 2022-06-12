package cfg

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func WriteDB(dbName string, isGtrEnabled interface{}, isDplEnabled interface{}) {
	db, err := sql.Open("sqlite3", "./"+dbName)

	if err != nil {
		panic(err)
	}

	stmt := Prepare(db, "CREATE TABLE IF NOT EXISTS diru (id INTEGER PRIMARY KEY AUTOINCREMENT, isGtrEnabled INTEGER, isDplEnabled INTEGER)")
	stmt.Exec()
	stmt.Close()

	var insrt *sql.Stmt

	if isGtrEnabled != nil && isDplEnabled != nil {
		insrt = Prepare(db, "INSERT INTO diru (isGtrEnabled, isDplEnabled) VALUES (?, ?)")
		insrt.Exec(isGtrEnabled, isDplEnabled)
	} else if isGtrEnabled != nil {
		insrt = Prepare(db, "INSERT INTO diru (isGtrEnabled) VALUES (?)")
		insrt.Exec(isGtrEnabled)
	} else if isDplEnabled != nil {
		insrt = Prepare(db, "INSERT INTO diru (isDplEnabled) VALUES (?)")
		insrt.Exec(isDplEnabled)
	}

	insrt.Close()

	db.Close()
}
