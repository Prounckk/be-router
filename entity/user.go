package entity

// User is parking spot consumer.
// User can check parking spot pricing, book the parking spot and pay for it
type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Parking  Parking `json:"takenParking"`
}
