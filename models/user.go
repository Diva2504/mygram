package models

type User struct {
	ID       uint
	Username string
	Email    string
	Password string
	Age      uint
}

type SocialMedia struct {
	ID             uint
	name           string
	SocialMediaUrl string
	UserId         int
}
