package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmed-aladdiin/gobank/types"
	"github.com/ahmed-aladdiin/gobank/utils"
)

func HandleDeposit(w http.ResponseWriter, r *http.Request, server *types.Server) {
	transaction := new(types.Transaction)
	json.NewDecoder(r.Body).Decode(transaction)

	err := server.DB.Deposit(transaction.AccountTo, transaction.Amount)
	if err != nil {
		status := http.StatusInternalServerError
		if msg := err.Error(); msg == "Account number was not found" {
			status = http.StatusNotFound
		}
		utils.WriteJSON(w, status, utils.ApiError{Error: err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func HandleWithdraw(w http.ResponseWriter, r *http.Request, server *types.Server) {
	transaction := new(types.Transaction)
	json.NewDecoder(r.Body).Decode(transaction)

	err := server.DB.Withdraw(transaction.AccountFrom, transaction.Amount)
	if err != nil {
		status := http.StatusInternalServerError
		if msg := err.Error(); msg == "Not sufficient balance" {
			status = http.StatusNotAcceptable
		}
		utils.WriteJSON(w, status, utils.ApiError{Error: err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func HandleTransfare(w http.ResponseWriter, r *http.Request, server *types.Server) {
	transaction := new(types.Transaction)
	json.NewDecoder(r.Body).Decode(transaction)

	err := server.DB.Transfare(transaction.AccountFrom, transaction.AccountTo, transaction.Amount)
	if err != nil {
		status := http.StatusInternalServerError
		if msg := err.Error(); msg == "Not sufficient balance" {
			status = http.StatusNotAcceptable
		}
		utils.WriteJSON(w, status, utils.ApiError{Error: err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
