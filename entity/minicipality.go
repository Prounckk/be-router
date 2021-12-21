package entity

type Municipality struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	ParkingGroups  []ParkingGroup   `json:"parkingGroups"`
	Administrators []Administrators `json:"administrators"`
	Inspectors     []Inspector      `json:"inspectors"`
}

type City struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Municipalities []Municipality `json:"municipalities"`
}

type Region struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Cities []City `json:"cities"`
}

type Country struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Regions []Region `json:"regions"`
}
