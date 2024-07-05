package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ahmed-aladdiin/gobank/handlers"
	"github.com/ahmed-aladdiin/gobank/types"
	"github.com/ahmed-aladdiin/gobank/utils"
)

func TransactionsRouter(w http.ResponseWriter, r *http.Request, server *types.Server) {
	if r.Method != "PUT" {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "Method not supported"})
		return
	}
	transaction_type := mux.Vars(r)["type"]
	switch transaction_type {
	case "deposit":
		handlers.HandleDeposit(w, r, server)
		return
	case "withdraw":
		handlers.HandleWithdraw(w, r, server)
		return
	case "transfare":
		handlers.HandleTransfare(w, r, server)
		return
	}

	utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "Action not supported"})
}
