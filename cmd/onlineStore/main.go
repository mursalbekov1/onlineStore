package main

import (
	"fmt"
	"net/http"
	"onlineStore/internal/config"
	"onlineStore/internal/route"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg.Http.Host)
	fmt.Println(cfg.Http.Port)

	r := route.Router()
	err := http.ListenAndServe(cfg.Http.Host+":"+cfg.Http.Port, r)
	if err != nil {
		return
	}

}
