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

	router.HandleFunc(
		"/accounts",
		func(w http.ResponseWriter, r *http.Request) {
			routers.AccountsRouter(w, r, server)
		},
	)

	router.HandleFunc(
		"/accounts/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			routers.AccountsRouterByID(w, r, server)
		},
	)

	router.HandleFunc(
		"/transaction/{type}",
		func(w http.ResponseWriter, r *http.Request) {
			routers.TransactionsRouter(w, r, server)
		},
	)

	log.Println("Server started at port ", server.ListenAddr)

	http.ListenAndServe(server.ListenAddr, router)
}
