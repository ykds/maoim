package main

import (
	"log"
	"maoim/internal/logic"
)

func main() {
	server := logic.New()
	if err := server.Start(); err != nil {
		log.Println(err)
	}
}
