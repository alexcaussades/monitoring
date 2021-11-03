package main

import (
	"fmt"
	"net/http"
	"test/urlliste"
	"test/webhookperso"
	"time"

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
