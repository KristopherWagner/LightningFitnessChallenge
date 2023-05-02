package types

type OauthToken struct {
	TokenType    string  `json:"token_type"`
	ExpiresAt    int     `json:"expires_at"`
	ExpiresIn    int     `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
	AccessToken  string  `json:"access_token"`
	Athlete      Athlete `json:"athlete"`
	Message      string  `json:"message"`
}
