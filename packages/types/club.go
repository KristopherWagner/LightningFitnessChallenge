package types

type Club struct {
	ID              int    `json:"id"`
	ResourceState   int    `json:"resource_state"`
	Name            string `json:"name"`
	ProfileMedium   string `json:"profile_medium"`
	CoverPhoto      string `json:"cover_photo"`
	CoverPhotoSmall string `json:"cover_photo_small"`
	SportType       string `json:"sport_type"`
	ActivityType    string `json:"activity_types"`
	City            string `json:"city"`
	State           string `json:"state"`
	Country         string `json:"country"`
	Private         bool   `json:"private"`
	MemberCount     int    `json:"member_count"`
	Featured        bool   `json:"featured"`
	Verified        bool   `json:"verified"`
	URL             string `json:"url"`
	Membership      string `json:"membership"`
	Admin           bool   `json:"admin"`
	Owner           string `json:"owner"`
	FollowingCount  int    `json:"following_count"`
}
