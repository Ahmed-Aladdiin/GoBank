package main

import "github.com/ahmed-aladdiin/gobank/types"

func main() {
	server := types.NewServer(":8000")
	Run(server)
}
