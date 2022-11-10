package main

import (
	"log"

	"github.com/cr0dr1gu3z/socialtwitt/bd"
	"github.com/cr0dr1gu3z/socialtwitt/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}

	handlers.Manejadores()
}
