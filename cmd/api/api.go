package api

import (
	"log"

	"github.com/B-danik/GameShop_Go/internal/user"
	"github.com/julienschmidt/httprouter"
)

type api struct {
}

func Header() (*httprouter.Router, error) {

	log.Println("Create router")
	router := httprouter.New()

	log.Println("Create handler")
	handler := user.NewHandler()
	handler.Register(router)
	return router, nil
}
