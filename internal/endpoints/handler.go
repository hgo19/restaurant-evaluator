package endpoints

import "restaurant-evaluator/internal/domain/user"

type Handler struct {
	UserService user.Service
}
