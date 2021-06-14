package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sebaz143/socialMedia/middleware"
	"github.com/sebaz143/socialMedia/routers"
)

/*Manejadores configura el puerto y pone a escuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter()

	//rutas
	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	handler := cors.AllowAll().Handler(router) //es una especie de middleware para revisar los permisos del cliente que se esta conectando
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
