package endpoints

import (
	"net/http"
	"restaurant-evaluator/internal/dto"

	"github.com/go-chi/render"
)

func (h *Handler) UserPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request dto.NewUser
	render.DecodeJSON(r.Body, &request)
	id, err := h.UserService.Create(request)
	return map[string]string{"id": id}, 201, err
}
