package main

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"
	_ "pet/docs" // Замените "your-package" на путь к вашему пакету с API
	"pet/handler/petHandler"
	"pet/repository/petRepo"
	"pet/service/petService"
)

func main() {
	r := mux.NewRouter()

	petHandler := petHandler.NewPetHandler(petService.NewPetService(petRepo.NewPetRepo()))

	r.HandleFunc("/pets/{id}", petHandler.GetPetByIDHandler).Methods("GET")
	r.HandleFunc("/pets", petHandler.PostPetHandler).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
