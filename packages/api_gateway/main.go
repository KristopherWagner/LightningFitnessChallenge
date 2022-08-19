package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ExpectedFormat struct {
	Code string `json:"code"`
}

func getTokenFromStrava(clientSecret, code string) (err error) {
	fmt.Println(code)
	response, err := http.PostForm("https://www.strava.com/oauth/token", url.Values{
		"client_id":     {"45959"},
		"client_secret": {clientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	})
	if err != nil {
		err = fmt.Errorf("unable to make request: %v", err)
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
	return
}

func handlePOSTToken(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token, access-control-allow-origin")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if req.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("unable to read body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	infoFromUser := ExpectedFormat{}
	err = json.Unmarshal(body, &infoFromUser)
	if err != nil {
		err = fmt.Errorf("unable to read json: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	code := infoFromUser.Code
	if len(code) == 0 {
		fmt.Printf("%+v\n", req)
		err = fmt.Errorf("missing code")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Get secret from environment
	err = getTokenFromStrava("", code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/token", handlePOSTToken)
	http.ListenAndServe(":8080", nil)
}
