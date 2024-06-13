package main

import (
	"fmt"
)

type UserRepo struct {
	listUser []User
}

// tanpa iface
// func NewUserRepo(lu []User) *UserRepo {
// dengan iface
func NewUserRepo(lu []User) IUserRepo {
	return &UserRepo{
		listUser: lu,
	}
}

type IUserRepo interface {
	GetAllUser() []User
	SetListUser([]User)
}

func (u *UserRepo) GetAllUser() []User {
	return u.listUser
}

func (u *UserRepo) SetListUser(in []User) {
	u.listUser = in
}

type User struct {
	listKelas []string
	Name      string
}

func NewUser(n string, lk []string) User {
	return User{
		Name:      n,
		listKelas: lk,
	}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GantiName(name string) {
	u.Name = name
}

func (u *User) GetListKelas() []string {
	return u.listKelas
}

func main() {
	// u := User{
	// 	Name:      "ibam",
	// 	listKelas: []string{"1", "2"},
	// }
	uIbam := NewUser("ibam", []string{"1", "2"})
	uBudi := NewUser("budi", []string{"1", "2"})
	// userRepo := UserRepo{listUser: []User{uIbam, uBudi}}
	userRepo := NewUserRepo([]User{uIbam, uBudi})
	listUsers := userRepo.GetAllUser()
	fmt.Println(listUsers)
	/*
		fmt.Println(u.Name)
		u.GantiName("budi")
		fmt.Println(u.GetName())
		fmt.Println(u.GetListKelas())
	*/
}
