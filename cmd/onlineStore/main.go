package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"onlineStore/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg.Http.Host)
	fmt.Println(cfg.Http.Port)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	err := http.ListenAndServe(cfg.Http.Host+":"+cfg.Http.Port, r)
	if err != nil {
		return
	}

}
