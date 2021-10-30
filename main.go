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

func Caseurl() {
	resp, _ := urlinport("https://www.alexcaussades.com")
	discord.WebhookURL = webhookperso.TokenPerso()
	
	if resp != "No access the adresse web" {
		switch resp {
		case "200 OK":
			fmt.Println("ok ", discord.Say("Hello, world!"))
		case "404":
			fmt.Println("404")
		default:
			fmt.Println("default", resp)
		}
	} else {
		fmt.Println("No access the adresse web")
	}

}

func main() {

	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Println(now.Format("15:04:05"))
		Caseurl()
	}
}
