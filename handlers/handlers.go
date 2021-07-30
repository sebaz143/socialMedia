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

	router := mux.NewRouter() //Es la ruta donde se captura el responseWriter y al Request.

	//rutas
	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT") //se revisa si esta variable de entorno existe.

	if PORT == "" {
		PORT = "8081"
	}

	handler := cors.AllowAll().Handler(router) //es una especie de middleware para revisar los permisos del cliente que se esta conectando

	//ponemos el http a escuchar en el puerto PORT y le pasamos de parametro de entrada
	//todo el handler que seteamos con cors que a su vez viene con un parametro de entrada router creado con mux
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
