package route

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"onlineStore/internal/handlers"
)

func Router(h handlers.Handler) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthCheck", healthCheckHandler)
	})

	return router
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
