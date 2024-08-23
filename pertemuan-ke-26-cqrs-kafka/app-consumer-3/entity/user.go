package entity

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Users = []User{
	{ID: 1, Name: "Emma"},
	{ID: 2, Name: "Bruno"},
	{ID: 3, Name: "Rick"},
	{ID: 4, Name: "Lena"},
}
