package database

import "restaurant-evaluator/internal/domain/user"

type UserRepository struct {
	users []user.User
}

func (r *UserRepository) Save(user *user.User) error {
	r.users = append(r.users, *user)
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	for i := range r.users {
		if r.users[i].Email == email {
			return &r.users[i], nil
		}
	}

	return nil, nil
}
