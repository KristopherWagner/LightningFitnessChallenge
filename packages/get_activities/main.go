package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/KristopherWagner/LightningFitnessChallenge/packages/types"
)

func makeGetRequest(accessToken string) (response []byte, err error) {
	startTime := time.Date(2022, 8, 8, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2022, 8, 14, 23, 59, 59, 0, time.Local)
	url, err := url.Parse("https://www.strava.com/api/v3/activities")
	if err != nil {
		err = fmt.Errorf("unable to create url: " + err.Error())
		return
	}
	q := url.Query()
	q.Set("after", fmt.Sprint(startTime.Unix()))
	q.Set("before", fmt.Sprint(endTime.Unix()))
	q.Set("page", "1")
	q.Set("per_page", "200")
	url.RawQuery = q.Encode()

	request, err := http.NewRequest("GET", url.String(), nil)
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

func main() {
	file, err := os.ReadFile("../register_athlete/tokens.json")
	if err != nil {
		log.Fatalf(err.Error())
	}

	userToken := types.TokenInfo{}
	err = json.Unmarshal([]byte(file), &userToken)
	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := makeGetRequest(userToken.AccessToken)
	if err != nil {
		log.Fatalf(err.Error())
	}

	activities := make([]types.SummaryActivity, 0)
	err = json.Unmarshal(response, &activities)
	if err != nil {
		log.Fatalf(err.Error())
	}

	totalDistance := float32(0)
	for _, activity := range activities {
		fmt.Printf(activity.StartDateLocal.String()+" "+activity.Name+": %.2f meters\n", activity.Distance)
		totalDistance += activity.Distance
	}

	fmt.Println("----------")
	fmt.Printf("Total distance: %.2f miles\n", totalDistance/1609.34)
}
