package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ahmed-aladdiin/gobank/routers"
	"github.com/ahmed-aladdiin/gobank/types"
)

func Run(server *types.Server) {
	router := mux.NewRouter()

	router.HandleFunc("/accounts", routers.AccountsRouter)
	router.HandleFunc("/accounts/{id}", routers.AccountsRouterByID)

	log.Println("Server started at port ", server.ListenAddr)

	http.ListenAndServe(server.ListenAddr, router)
}
