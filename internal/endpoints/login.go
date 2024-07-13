package endpoints

import (
	"net/http"
	"restaurant-evaluator/internal/dto"

	"github.com/go-chi/render"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request dto.Login
	render.DecodeJSON(r.Body, &request)
	token, err := h.UserService.GenerateToken(request.Email)

	return map[string]string{"token": token}, 200, err
}
