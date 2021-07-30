package main

import (
	"log"

	"github.com/sebaz143/socialMedia/bd"
	"github.com/sebaz143/socialMedia/handlers"
)

func main() {

	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
	}
	handlers.Manejadores()

}
