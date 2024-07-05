package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmed-aladdiin/gobank/types"
	"github.com/ahmed-aladdiin/gobank/utils"
)

func HandleGETccounts(w http.ResponseWriter, r *http.Request, server *types.Server) {
	accounts, err := server.DB.GetAccounts()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, accounts)
}

func HandleGETccount(w http.ResponseWriter, r *http.Request, server *types.Server) {
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	account, err := server.DB.GetAccountByID(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, account)
}

func HandlePOSTAccount(w http.ResponseWriter, r *http.Request, server *types.Server) {
	account := new(types.Account)
	json.NewDecoder(r.Body).Decode(account)

	newAcc, err := server.DB.CreateAccount(account)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, newAcc)
}

func HandleDELETEAccount(w http.ResponseWriter, r *http.Request, server *types.Server) {
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	err2 := server.DB.DeleteAccount(id)
	if err2 != nil {
		utils.WriteJSON(w, http.StatusNotFound, utils.ApiError{Error: err2.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
