package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"test/urlliste"
	"test/webhookperso"
	"time"

	"github.com/ecnepsnai/discord"
	_ "github.com/mattn/go-sqlite3"
)

func Caseurl(value string) {
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
		url TEXT
	)
	`
	stmt, err := db.Prepare(chemain)
	if err != nil {
		fmt.Printf("Cannot prepare statement. err=%v\n", err)
		os.Exit(1)
	}
	stmt.Exec()
	resp, error := urlliste.Urlinport(value)
	discord.WebhookURL = webhookperso.TokenPerso()
	now := time.Now()
	req, err := db.Query("SELECT * FROM statusurls WHERE url = ?", value)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(req , value)

	if req.Next() {
		var id int
		var code string
		var url string
		err := req.Scan(&id, &code, &url)
		if err != nil {
			log.Fatal(err)
		}
		responsecode := strings.Split(resp, " ")
		if url == value {
			log.Println(code, url, id, resp)
			if code != responsecode[0] {
				db.Query("UPDATE statusurls SET code = ? WHERE id = ? AND url = ?", resp, id, url)
			} 
		}
	} else {
		req, _ := db.Prepare("INSERT INTO statusurls(code, url) VALUES(?, ?)")
		req.Exec("200", value)
		log.Println("insert is good")
	}

	if resp != "No access the adresse web" {
		switch resp {
		case "200 OK":
			fmt.Println("OK" + " " + value)
		case "404":
			fmt.Println(discord.Say(now.Format("15:04:05") + " " + value + " => erreur: 404"))
		case "500":
			fmt.Println(discord.Say((now.Format("15:04:05") + " " + value + " => erreur: 500")))
		case "503":
			fmt.Println(discord.Say((now.Format("15:04:05") + " " + value + " => erreur: 500")))
		case "504":
			fmt.Println(discord.Say((now.Format("15:04:05") + " " + value + " => erreur: 504")))
		default:
			fmt.Println(discord.Say((now.Format("15:04:05") + "default: " + value + error.Error())))
		}
	} else {
		fmt.Println(discord.Say((now.Format("15:04:05") + " " + value + " => erreur: NO DNS")))
	}

}

func main() {
	m := urlliste.Setmap()
	for {
		for i := 0; i < len(m); i++ {
			Caseurl(m[i])
			time.Sleep(6 * time.Second)
		}
	}
}
