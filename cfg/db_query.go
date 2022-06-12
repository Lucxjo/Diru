package cfg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func QueryOne(dbName string, gid int) DiruData {
	db, err := sql.Open("sqlite3", "./"+dbName)

	if err != nil {
		panic(err)
	}

	stmt := Prepare(db, "CREATE TABLE IF NOT EXISTS diru (id INTEGER PRIMARY KEY AUTOINCREMENT, isGtrEnabled INTEGER, isDplEnabled INTEGER)")
	stmt.Exec()
	stmt.Close()

	rows, err := db.Query("SELECT id, task, complete FROM tasks WHERE id = ?", gid)

	if err != nil {
		panic(err)
	}

	var data DiruData
	for rows.Next() {
		rows.Scan(&data.ID, &data.GtrEnabled, &data.DplEnabled)
	}

	defer db.Close()
	return data
}
