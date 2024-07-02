package routers

import (
	"net/http"

	"github.com/ahmed-aladdiin/gobank/handlers"
	"github.com/ahmed-aladdiin/gobank/utils"
)

func AccountsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.HandleGETccounts(w, r)
		return
	case "POST":
		handlers.HandlePOSTAccount(w, r)
		return
	}

	utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "Method not supported1"})
}

func AccountsRouterByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.HandleGETccount(w, r)
		return
	case "DELETE":
		handlers.HandleDELETEAccount(w, r)
		return
	}

	utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "Method not supported2"})
}
