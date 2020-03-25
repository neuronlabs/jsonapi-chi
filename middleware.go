package jsonapichi

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/neuronlabs/jsonapi-handler"
)

// GetID is the middleware that stores the 'id' in the request context with the key
// specified in the `jsonapi-handler.IDKey`.
func GetID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, "id")
		ctx := context.WithValue(req.Context(), handler.IDKey, id)
		req = req.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
