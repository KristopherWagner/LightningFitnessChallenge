package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/KristopherWagner/LightningFitnessChallenge/packages/types"
)

func makeGetRequest(accessToken string, url string) (response []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("unable to generate request: " + err.Error())
		return
	}

	bearer := "Bearer " + accessToken
	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(request)
	if err != nil {
		err = fmt.Errorf("unable to complete request: " + err.Error())
		return
	}

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("unable to read request body: " + err.Error())
		return
	}

	err = resp.Body.Close()
	if err != nil {
		fmt.Println("WARNING: Unable to close response body.")
		err = nil
	}
	return
}

func GetActivities(accessToken string) (activities []types.SummaryActivity, err error) {
	url := "https://www.strava.com/api/v3/clubs/502620/activities"
	response, err := makeGetRequest(accessToken, url)
	if err != nil {
		err = fmt.Errorf("failed to make GET request: " + err.Error())
		return
	}

	activities = make([]types.SummaryActivity, 0)
	err = json.Unmarshal(response, &activities)
	if err != nil {
		err = fmt.Errorf("unable to unmarshal response: " + err.Error())
	}
	return
}

func GetAthletes(accessToken string) (athletes []types.SummaryAthlete, err error) {
	url := "https://www.strava.com/api/v3/clubs/502620/members"
	response, err := makeGetRequest(accessToken, url)
	if err != nil {
		err = fmt.Errorf("failed to make GET request: " + err.Error())
		return
	}

	athletes = make([]types.SummaryAthlete, 0)
	err = json.Unmarshal(response, &athletes)
	if err != nil {
		err = fmt.Errorf("unable to unmarshal response: " + err.Error())
	}
	return
}

func getAthlete(accessToken string) (athlete types.SummaryAthlete, err error) {
	url := "https://www.strava.com/api/v3/athlete"
	response, err := makeGetRequest(accessToken, url)
	if err != nil {
		err = fmt.Errorf("failed to make GET request: " + err.Error())
		return
	}

	err = json.Unmarshal(response, &athlete)
	if err != nil {
		err = fmt.Errorf("unable to unmarshal response: " + err.Error())
	}
	return
}

func main() {
	accessToken := flag.String("token", "", "Access Token from Strava")
	flag.Parse()
	athlete, err := getAthlete(*accessToken)
	if err != nil {
		err = fmt.Errorf("Unable to get athlete: " + err.Error())
		log.Fatalf(err.Error())
	}
	fmt.Println(athlete.FirstName + " " + athlete.LastName)
}
