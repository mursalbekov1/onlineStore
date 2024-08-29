package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"onlineStore/internal/models"
	"strconv"
)

func (h Handler) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h Handler) addProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	if err := h.DB.Create(&product).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h Handler) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateProduct models.Product
	json.Unmarshal(body, &updateProduct)

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	product.Name = updateProduct.Name
	product.Description = updateProduct.Description
	product.Category = updateProduct.Category
	product.Price = updateProduct.Price

	h.DB.Save(&product)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h Handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var product models.Product

	if err := h.DB.Delete(&product, id).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Product deleted")
}
