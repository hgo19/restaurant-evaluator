package user

type Repository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
}
