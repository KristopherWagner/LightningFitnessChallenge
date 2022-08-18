package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/KristopherWagner/LightningFitnessChallenge/packages/types"
)

func createTokenInfo() (tokenInfo types.TokenInfo) {
	athleteID := flag.String("athlete", "", "ID of Athlete")
	accessToken := flag.String("token", "", "Access Token from Strava")
	refreshToken := flag.String("refresh", "", "Refresh Token from Strava")
	expires := flag.Int("expires", 0, "Epoch timestamp of when the token expires")
	flag.Parse()

	tokenInfo = types.TokenInfo{
		ID:           *athleteID,
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
		Expires:      *expires,
	}
	return
}

func writeToFile(tokenInfo types.TokenInfo) (err error) {
	file, err := json.MarshalIndent(tokenInfo, "", "  ")
	if err != nil {
		err = fmt.Errorf("unable to Marshal JSON: " + err.Error())
		return
	}

	err = ioutil.WriteFile("tokens.json", file, 0644)
	if err != nil {
		err = fmt.Errorf("unable to write to file: " + err.Error())
	}
	return
}

func main() {
	tokenInfo := createTokenInfo()
	err := writeToFile(tokenInfo)
	if err != nil {
		log.Fatal("Failed to write user to file: " + err.Error())
	}
}
