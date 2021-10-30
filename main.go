package main

import (
	"fmt"
	"net/http"
	"time"

	"./webhookperso"
	"github.com/ecnepsnai/discord"
)

func urlinport(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "No access the adresse web", err
	}
	defer resp.Body.Close()
	return resp.Status, nil
}

func getmap() map[int]string {

	m := make(map[int]string)
	m[0] = "https://www.alexcaussades.com"
	m[1] = "https://ouioweb.com/"
	m[2] = "https://myalternos.fr/"
	return m

}

func Caseurl(value string) {

	/**
	* changement de formule sur la function
	 */

	resp, error := urlinport(value)
	discord.WebhookURL = webhookperso.TokenPerso()
	now := time.Now()
	if resp != "No access the adresse web" {
		switch resp {
		case "200 OK":
			
		case "404":
			fmt.Println(discord.Say(now.Format("15:04:05") + " " + resp + " => erreur: 404"))
		case "500":
			fmt.Println(discord.Say((now.Format("15:04:05") + " " + resp + " => erreur: 500")))
		case "503":
			fmt.Println(discord.Say((now.Format("15:04:05") + " " + resp + " => erreur: 500")))
		case "504":
			fmt.Println(discord.Say((now.Format("15:04:05") + " " + resp + " => erreur: 504")))
		default:
			fmt.Println("default", resp)
		}
	} else {
		fmt.Println(discord.Say((now.Format("15:04:05") + " " + error.Error() )))
	}

}

func main() {
	m := getmap()
	c := time.Tick(60 * time.Second)
	for now := range c {
		for _, i := range m {
			fmt.Println(now.Format("15:04:05"))
			Caseurl(i)
		}

	}

}
