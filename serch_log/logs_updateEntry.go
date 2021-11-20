package srech_log

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func logs_update(id string, time_out int64) {
	db, err := sql.Open("sqlite3", "test1.db")
	if err != nil {
		fmt.Printf("Cannot open database. err=%v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	var logs = `
	CREATE TABLE IF NOT EXISTS logs
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		url TEXT,
		date TEXT,
		time_in INTEGER NOT NULL,
		time_out INTEGER DEFAULT 0
	)
	`
	stmt, err := db.Prepare(logs)
	if err != nil {
		fmt.Printf("Cannot prepare statement. err=%v\n", err)
		os.Exit(1)
	}
	stmt.Exec()

	stmt, err = db.Prepare("UPDATE logs SET time_out = ? WHERE id = ?")
	if err != nil {
		fmt.Printf("Cannot prepare statement. err=%v\n", err)
		os.Exit(1)
	}
	stmt.Exec(time_out, id)
}
