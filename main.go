package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"test/urlliste"
	"test/webhookperso"
	"time"

	"github.com/ecnepsnai/discord"
	_ "github.com/mattn/go-sqlite3"
)

type JsonInfo struct {
	Url  string `json:"url"`
	Code string `json:"code"`
	Time int64  `json:"time"`
}

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
		time_in INTEGER,
		time_out INTEGER
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

	resp, error := urlliste.Urlinport(value)
	discord.WebhookURL = webhookperso.TokenPerso()
	now := time.Now()
	req, err := db.Query("SELECT * FROM statusurls WHERE url = ?", value)
	if err != nil {
		log.Fatal(err)
	}

	if req.Next() {
		var id int
		var code string
		var url string
		var time int
		err := req.Scan(&id, &code, &url, &time)
		if err != nil {
			log.Fatal(err)
		}
		responsecode := strings.Split(resp, " ")
		if url == value {
			log.Println(code, url, id, resp)
			if code != responsecode[0] {
				db.Query("UPDATE statusurls SET code = ? WHERE id = ? AND url = ?", resp, id, url)
				db.Query("UPDATE statusurls SET time = ? WHERE id = ? AND url = ?", now.Unix(), id, url)
				if responsecode[0] != "200" {
					db.Query("INSERT INTO logs (code, url, date, time_in) VALUES (?, ?, ?, ?)", responsecode[0], url, now.Format("2006-01-02"), now.Unix())
				} else {
					req, err := db.Query("SELECT max(id) FROM logs WHERE url = ?", url)
					if err != nil {
						log.Fatal(err)
					}
					if req.Next() {
						var id int
						var code string
						var url string
						var date string
						var time_in int
						var time_out int
						err := req.Scan(&id, &code, &url, &date, &time_in, &time_out)
						if err != nil {
							log.Fatal(err)
						}
						if time_out == 0 {
							db.Query("UPDATE logs SET time_out = ? WHERE id = ? AND url = ?", now.Unix(), id, url)
						}

					}
				}
			} //sans doute un probleme ici avec le time_out
		}
	} else {
		req, _ := db.Prepare("INSERT INTO statusurls(code, url, time) VALUES(?, ?, ?)")
		req.Exec("200", value, now.Unix())
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

func jsonInformation(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "test1.db")

	if err != nil {
		fmt.Printf("Cannot open database. err=%v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	res, _ := db.Query("SELECT count(id) FROM statusurls")

	
	if res.Next() {
		var id int
		err := res.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		
		for i := 0; i < id; i++ {
			res, _ := db.Query("SELECT * FROM statusurls WHERE id = ?", i)
			if res.Next() {
				var id int
				var code string
				var url string
				var time int
				err := res.Scan(&id, &code, &url, &time)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(id, code, url, time)
			}
		}

	}
	// jsonInf := []JsonInfo{
	// 	{
	// 		Code: "200",
	// 		Url: "https://www.google.com",
	// 		Time: "15:04:05",
	// 	},
	// }

}



func main() {

	http.HandleFunc("/monitoring-serveur", jsonInformation)
	http.ListenAndServe(":8080", nil)

	m := urlliste.Setmap()
	for {
		for i := 0; i < len(m); i++ {
			Caseurl(m[i])
			time.Sleep(6 * time.Second)
		}
	}
}
