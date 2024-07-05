package routers

import (
	"net/http"

	"github.com/ahmed-aladdiin/gobank/handlers"
	"github.com/ahmed-aladdiin/gobank/types"
	"github.com/ahmed-aladdiin/gobank/utils"
)

func AccountsRouter(w http.ResponseWriter, r *http.Request, server *types.Server) {
	switch r.Method {
	case "GET":
		handlers.HandleGETccounts(w, r, server)
		return
	case "POST":
		handlers.HandlePOSTAccount(w, r, server)
		return
	}

	utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "Method not supported1"})
}

func AccountsRouterByID(w http.ResponseWriter, r *http.Request, server *types.Server) {
	switch r.Method {
	case "GET":
		handlers.HandleGETccount(w, r, server)
		return
	case "DELETE":
		handlers.HandleDELETEAccount(w, r, server)
		return
	}

	utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "Method not supported2"})
}
