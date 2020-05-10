package server

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type FleetUnit struct {
	ID string `json:"id" rethinkdb:"id"`

	Position Position `json:"position"`
}
