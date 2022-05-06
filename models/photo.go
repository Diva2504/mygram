package models

type Photo struct {
	ID       int
	Title    string
	Caption  string
	PhotoUrl string
	User     *User
}
