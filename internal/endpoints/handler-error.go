package endpoints

import (
	"errors"
	"net/http"
	internalerrors "restaurant-evaluator/internal/internal-errors"

	"github.com/go-chi/render"
)

type EndpointsFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointsFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)
		if err != nil {
			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, 500)
			} else if errors.Is(err, internalerrors.NotFound) {
				render.Status(r, 404)
			} else {
				render.Status(r, 400)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
