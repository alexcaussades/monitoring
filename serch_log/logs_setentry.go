package srech_log

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func logs_setEntry(code string, url string) {
	db, err := sql.Open("sqlite3", "test1.db")
	if err != nil {
		fmt.Printf("Cannot open database. err=%v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	var chemain = `
	CREATE TABLE IF NOT EXISTS statusurls
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		url TEXT,
		time INTEGER
	)
	`
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
	stmt, err := db.Prepare(chemain)
	if err != nil {
		fmt.Printf("Cannot prepare statement. err=%v\n", err)
		os.Exit(1)
	}
	stmt.Exec()

	stmt, err = db.Prepare(logs)
	if err != nil {
		fmt.Printf("Cannot prepare statement. err=%v\n", err)
		os.Exit(1)
	}
	stmt.Exec()

	now := time.Now()
	var time_in int64 = now.Unix()
	var date = now.Format("2006-01-02")

	stmt, err = db.Prepare("INSERT INTO logs (code, url, date, time_in) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("Cannot prepare statement. err=%v\n", err)
		os.Exit(1)
	}
	stmt.Exec(code, url, date, time_in)
}
