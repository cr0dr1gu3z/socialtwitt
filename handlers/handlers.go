package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/cr0dr1gu3z/socialtwitt/middlew"
	"github.com/cr0dr1gu3z/socialtwitt/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores seteo mi puerto, el handler y pongo a escuchar el server
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//cors son los permisos que se le dan a la api
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
