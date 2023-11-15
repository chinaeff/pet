package router

import (
	"github.com/gorilla/mux"
	"pet/handler/petHandler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/pets/{id}", petHandler.PetHandler{}.GetPetByIDHandler).Methods("GET")
	r.HandleFunc("/pets", petHandler.PetHandler{}.PostPetHandler).Methods("POST")

	return r
}
