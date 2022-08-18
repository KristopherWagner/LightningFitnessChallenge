// https://developers.strava.com/docs/reference/#api-models-SummaryActivity
package types

import (
	"time"
)

type SummaryActivity struct {
	ID                   int         `json:"id"`
	ExternalID           string      `json:"external_id"`
	UploadID             int         `json:"upload_id"`
	Athlete              Athlete     `json:"athlete"`
	Name                 string      `json:"name"`
	Distance             float32     `json:"distance"`
	MovingTime           int         `json:"moving_time"`
	ElapsedTime          int         `json:"elapsed_time"`
	TotalElevationGain   float32     `json:"total_elevation_gain"`
	ElevHigh             float32     `json:"elev_high"`
	ElevLow              float32     `json:"elev_low"`
	SportType            string      `json:"sport_type"` // https://developers.strava.com/docs/reference/#api-models-SportType
	StartDate            string      `json:"start_date"`
	StartDateLocal       time.Time   `json:"start_date_local"`
	Timezone             string      `json:"timezone"`
	StartLatLng          []float64   `json:"start_latlng"`
	EndLatLng            []float64   `json:"end_latlng"`
	AchievementCount     int         `json:"achievement_count"`
	KudosCount           int         `json:"kudos_count"`
	CommentCount         int         `json:"comment_count"`
	AthleteCount         int         `json:"athlete_count"`
	PhotoCount           int         `json:"photo_count"`
	TotalPhotoCount      int         `json:"total_photo_count"`
	Map                  PolylineMap `json:"map"`
	Trainer              bool        `json:"trainer"`
	Commute              bool        `json:"commute"`
	Manual               bool        `json:"manual"`
	Private              bool        `json:"private"`
	Flagged              bool        `json:"flagged"`
	WorkoutType          int         `json:"workout_type"`
	UploadIDStr          string      `json:"upload_id_str"`
	AverageSpeed         float32     `json:"average_speed"`
	MaxSpeed             float32     `json:"max_speed"`
	HasKudoed            bool        `json:"has_kudoed"`
	HideFromHome         bool        `json:"hide_from_home"`
	GearID               string      `json:"gear_id"`
	Kilojoules           float32     `json:"kilojoules"`
	AverageWatts         float32     `json:"average_watts"`
	DeviceWatts          bool        `json:"device_watts"`
	MaxWatts             int         `json:"max_watts"`
	WeightedAverageWatts int         `json:"weighted_average_watts"`
}
