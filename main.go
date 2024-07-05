package main

import (
	"log"

	"github.com/ahmed-aladdiin/gobank/types"
)

func main() {
	db, err := types.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
		return
	}

	err2 := db.Init()
	if err2 != nil {
		log.Fatal(err)
		return
	}

	server := types.NewServer(":8000", db)
	Run(server)
}
