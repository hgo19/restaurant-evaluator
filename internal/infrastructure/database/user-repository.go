package database

import "restaurant-evaluator/internal/domain/user"

type UserRepository struct {
	users []user.User
}

func (r *UserRepository) Save(user *user.User) error {
	r.users = append(r.users, *user)
	return nil
}
