package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"onlineStore/internal/models"
	"onlineStore/internal/repository"
	"strconv"
)

type ProductHandler struct {
	ProductRepo *repository.ProductRepo
}

func NewProductHandler(ur *repository.ProductRepo) *ProductHandler {
	return &ProductHandler{ProductRepo: ur}
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	product, err := h.ProductRepo.GetByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductRepo.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) SaveProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.ProductRepo.Save(&product); err != nil {
		http.Error(w, "Failed to save product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	var updateProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updateProduct.Id = id

	if err := h.ProductRepo.Update(id, &updateProduct); err != nil {
		http.Error(w, "Failed to update prodyct", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Product updated successfully")
}
