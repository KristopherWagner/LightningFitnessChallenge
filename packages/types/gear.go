package types

type Gear struct {
	ID            string  `json:"id"`
	ResourceState int     `json:"resource_state"`
	Primary       bool    `json:"primary"`
	Name          string  `json:"name"`
	Distance      float64 `json:"distance"`
	Brand         string  `json:"brand_name"`
	Model         string  `json:"model_name"`
	Frame         string  `json:"frame_type"` // bike only
	Description   string  `json:"description"`
}
