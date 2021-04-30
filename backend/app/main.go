package main

import (
	"log"
	"todo-list/app/interfaces"
)

func main() {
	if err := interfaces.Dispatch(); err != nil {
		log.Fatal(err)
	}
}
