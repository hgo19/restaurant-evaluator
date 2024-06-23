package dto

type NewUser struct {
	Username     string
	Email        string
	PasswordHash string
	UserType     string
	Token        string
}
