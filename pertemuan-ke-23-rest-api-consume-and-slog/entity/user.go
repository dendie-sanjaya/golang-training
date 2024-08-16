package entity

type User struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Address  UserAddress `json:"address"`
	Phone    string      `json:"phone"`
	Website  string      `json:"website"`
	Company  UserCompany `json:"company"`
}

type UserAddress struct {
	Street  string  `json:"street"`
	Suite   string  `json:"suite"`
	City    string  `json:"city"`
	ZipCode string  `json:"zipcode"`
	Geo     UserGeo `json:"geo"`
}

type UserGeo struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
}

type UserCompany struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Business    string `json:"bs"`
}
