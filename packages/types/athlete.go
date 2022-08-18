package types

import "time"

type Athlete struct {
	ID                       int       `json:"id"`
	ResourceState            int       `json:"resource_state"`
	FirstName                string    `json:"firstname"`
	LastName                 string    `json:"lastname"`
	ProfileMedium            string    `json:"profile_medium"`
	Profile                  string    `json:"profile"`
	City                     string    `json:"city"`
	State                    string    `json:"state"`
	Country                  string    `json:"country"`
	Sex                      string    `json:"sex"`
	Summit                   bool      `json:"summit"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	FollowerCount            int       `json:"follower_count"`
	FriendCount              int       `json:"friend_count"`
	MeasurementPreference    string    `json:"measurement_preference"`
	FunctionalThresholdPower int       `json:"ftp"`
	Weight                   float64   `json:"weight"`
	Clubs                    []Club    `json:"clubs"`
	Bikes                    []Gear    `json:"bikes"`
	Shoes                    []Gear    `json:"shoes"`
}
