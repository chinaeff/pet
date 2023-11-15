package storeHandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"pet/mw"
	"pet/repository/storeRepo"
	"pet/service/storeService"
	"strconv"
)

type StoreHandler struct {
	service *storeService.StoreService
}

func NewStoreHandler(service *storeService.StoreService) *StoreHandler {
	return &StoreHandler{service: service}
}

func RegisterStoreHandlers(r *mux.Router, service *storeService.StoreService) {
	handler := NewStoreHandler(service)

	r.Handle("/store/inventory", mw.TokenAuthMiddleware(http.HandlerFunc(handler.InventoryHandler))).Methods("GET")
	r.Handle("/store/order", mw.TokenAuthMiddleware(http.HandlerFunc(handler.OrderHandler))).Methods("POST")
	r.Handle("/store/order/{id}", mw.TokenAuthMiddleware(http.HandlerFunc(handler.GetOrderHandler))).Methods("GET")
	r.Handle("/store/order/{id}", mw.TokenAuthMiddleware(http.HandlerFunc(handler.DeleteOrderHandler))).Methods("DELETE")

}

func (h *StoreHandler) InventoryHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	result, err := h.service.Inventory(status)
	if err != nil {
		http.Error(w, "Failed to get inventory", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(result)
}
func (h *StoreHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	IDStr := r.URL.Query().Get("id")
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetOrder(ID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *StoreHandler) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	IDStr := r.URL.Query().Get("id")
	ID, err := strconv.ParseInt(IDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteOrder(ID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *StoreHandler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	var newOrder storeRepo.Store
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	_, err = h.service.Order(newOrder)
	if err != nil {
		http.Error(w, "Failed to place order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
