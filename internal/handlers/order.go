package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"onlineStore/internal/models"
	"strconv"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func (h Handler) addOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var order models.Orders
	err = json.Unmarshal(body, &order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if result := h.DB.Create(&order); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h Handler) getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var order models.Orders
	if result := h.DB.First(&order, id); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Orders

	if result := h.DB.Find(&orders); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h Handler) updateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var updateOrder models.Orders
	err = json.Unmarshal(body, &updateOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var order models.Orders
	if result := h.DB.First(&order, id); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	order.Price = updateOrder.Price
	order.Status = updateOrder.Status

	h.DB.Save(&order)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h Handler) deleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var order models.Orders

	if result := h.DB.Delete(&order, id); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("deleted")
}
