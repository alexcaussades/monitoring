package main

import (
	"fmt"
	"net/http"
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


func Caseurl() {
	resp, _ := urlinport("https://www.alexcaussades.com")
	discord.WebhookURL = "https://discord.com/api/webhooks/787426472268398592/iH7huKfDkUu_fA_0cSh5l1SFbAAQEsKibWBWpuQ_hFKLxRLcLEeAi5R_QdSnsvSLHQ7m"

	if resp != "No access the adresse web" {
		switch resp {
		case "200 OK":
			fmt.Println("ok ", discord.Say("Hello, world!"))
			//send("ok")
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
	// time.AfterFunc(1*time.Second,func() {
	// 	Caseurl()
	// })
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Println(now.Format("15:04:05"))
		Caseurl()
	}
}
