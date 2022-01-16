package entity

// Administrators manage parking groups belongs to a Municipality
// it can change the pricing model and status of ParkingGroups or individual parking spots
type Administrators struct {
	ID            int            `json:"id"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	ParkingGroups []ParkingGroup `json:"parkingGroups"`
}

// Inspector supervise parking groups belongs to a Municipality
type Inspector struct {
	ID            int            `json:"id"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	ParkingGroups []ParkingGroup `json:"parkingGroups"`
}
