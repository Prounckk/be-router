package entity

// User is parking spot consumer.
// User can check parking spot pricing, book the parking spot and pay for it
// User can have zero or more cars
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}

// Car is a car with type and licence plate number.
// Car can have a nickname, model and colour. It can belong to a multiple users
type Car struct {
	ID       int    `json:"id"`
	Nickname int    `json:"nickname"`
	Type     string `json:"type"`
	Colour   string `json:"colour"`
	Number   string `json:"number"`
}

// UserParking is many to manu relation between Users and their parking's
type UserParking struct {
	User    User    `json:"user"`
	Parking Parking `json:"takenParking"`
}

// UserCars is many to manu relation between Users and their cars
type UserCars struct {
	User User `json:"user"`
	Car  Car  `json:"cars"`
}
