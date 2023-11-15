package petHandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"pet/mw"
	"pet/repository"
	"pet/repository/petRepo"
	"pet/service/petService"
	"strconv"
)

type PetHandler struct {
	service *petService.PetService
}

func NewPetHandler(service *petService.PetService) *PetHandler {
	return &PetHandler{service: service}
}

func RegisterPetHandlers(r *mux.Router, service *petService.PetService) {
	handler := NewPetHandler(service)

	r.Handle("/pets/{id}", mw.TokenAuthMiddleware(http.HandlerFunc(handler.GetPetByIDHandler))).Methods("GET")
	r.Handle("/pets", mw.TokenAuthMiddleware(http.HandlerFunc(handler.GetPetByStatusHandler))).Methods("GET")
	r.Handle("/pets/{id}", mw.TokenAuthMiddleware(http.HandlerFunc(handler.UpdatePetByIDHandler))).Methods("PUT")
	r.Handle("/pets/{id}", mw.TokenAuthMiddleware(http.HandlerFunc(handler.DeletePetByIDHandler))).Methods("DELETE")
	r.Handle("/pets/{id}", mw.TokenAuthMiddleware(http.HandlerFunc(handler.PostPetByIDHandler))).Methods("POST")
	r.Handle("/pets/{id}/uploadImage", mw.TokenAuthMiddleware(http.HandlerFunc(handler.PostImageByIDHandler))).Methods("POST")
	r.Handle("/pets", mw.TokenAuthMiddleware(http.HandlerFunc(handler.PostPetHandler))).Methods("POST")
}

// GetPetByIDHandler возвращает информацию о питомце по его ID.
// swagger:route GET /pets/{id} pets getPetByID
// Возвращает информацию о питомце по его ID.
// Responses:
//
//	200: petResponse
//	404: errorResponse
func (h *PetHandler) GetPetByIDHandler(w http.ResponseWriter, r *http.Request) {
	IDStr := r.URL.Query().Get("id")
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	pet, err := h.service.GetPetByID(ID)
	if err != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
	log.Print(http.StatusOK)
}

func (h *PetHandler) GetPetByStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	pet, err := h.service.GetPetByStatus(status)
	if err != nil {
		http.Error(w, "Pet not found", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
	log.Print(http.StatusOK)
}

func (h *PetHandler) UpdatePetByIDHandler(w http.ResponseWriter, r *http.Request) {
	IDStr := r.URL.Query().Get("id")
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	var updatedPet petRepo.Pet
	err = json.NewDecoder(r.Body).Decode(&updatedPet)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	err = h.service.PutPetByID(ID, updatedPet)
	if err != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PetHandler) DeletePetByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	ID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	_, err = h.service.DeletePetByID(ID)
	if err != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}

// PostPetHandler добавляет нового питомца.
// swagger:route POST /pets pets postPet
// Добавляет нового питомца.
// Responses:
//
//	201: petResponse
//	400: errorResponse
func (h *PetHandler) PostPetByIDHandler(w http.ResponseWriter, r *http.Request) {
	IDStr := r.URL.Query().Get("id")
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse data", http.StatusBadRequest)
		return
	}
	formData := repository.FormData{
		Name:   r.Form.Get("name"),
		Status: r.Form.Get("status"),
	}

	err = h.service.PostPetByID(ID, formData)
	if err != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *PetHandler) PostImageByIDHandler(w http.ResponseWriter, r *http.Request) {
	IDStr := r.URL.Query().Get("id")
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	image, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	_, err = h.service.PostImageByID(ID, image)
	if err != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *PetHandler) PostPetHandler(w http.ResponseWriter, r *http.Request) {
	var newPet petRepo.Pet
	err := json.NewDecoder(r.Body).Decode(&newPet)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	_, err = h.service.PostPet(newPet)
	if err != nil {
		http.Error(w, "Failed to add pet", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
