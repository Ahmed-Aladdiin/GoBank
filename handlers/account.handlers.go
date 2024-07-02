package handlers

import (
	"log"
	"net/http"

	"github.com/ahmed-aladdiin/gobank/types"
	"github.com/ahmed-aladdiin/gobank/utils"
)

func HandleGETccounts(w http.ResponseWriter, r *http.Request) {
}

func HandleGETccount(w http.ResponseWriter, r *http.Request) {
	log.Println("once")
	utils.WriteJSON(w, http.StatusOK, types.NewAccount("Ahmed", "Aladdin"))
}

func HandlePOSTAccount(w http.ResponseWriter, r *http.Request) {
}

func HandleDELETEAccount(w http.ResponseWriter, r *http.Request) {
}
