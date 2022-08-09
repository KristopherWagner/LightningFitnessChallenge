// https://developers.strava.com/docs/reference/#api-models-PolylineMap
package types

type PolylineMap struct {
	ID              string `json:"id"`
	Polyline        string `json:"polyline"`
	SummaryPolyline string `json:"summary_polyline"`
}
