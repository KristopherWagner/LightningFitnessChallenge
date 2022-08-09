package types

import "time"

type SummaryAthlete struct {
	ID            int       `json:"id"`
	ResourceState int       `json:"resource_state"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	ProfileMedium string    `json:"profile_medium"`
	Profile       string    `json:"profile"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	Sex           string    `json:"sex"`
	Summit        bool      `json:"summit"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
