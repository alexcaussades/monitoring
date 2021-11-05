package urlliste

import "net/http"

func Urlinport(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "No access the adresse web", err
	}
	defer resp.Body.Close()
	return resp.Status, nil
}
