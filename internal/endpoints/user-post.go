package endpoints

import (
	"errors"
	"net/http"
	"restaurant-evaluator/internal/dto"
	internalerrors "restaurant-evaluator/internal/internal-errors"

	"github.com/go-chi/render"
)

func (h *Handler) UserPost(w http.ResponseWriter, r *http.Request) {
	var request dto.NewUser
	render.DecodeJSON(r.Body, &request)
	id, err := h.UserService.Create(request)

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}
