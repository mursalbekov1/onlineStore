package repository

import (
	"fmt"
	"gorm.io/gorm"
	"onlineStore/internal/models"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (r *ProductRepo) GetByID(id int) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get product by id: %d: %w", id, err)
	}

	return &product, nil
}

func (r *ProductRepo) GetAll() ([]models.Product, error) {
	var products []models.Product

	err := r.db.Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}

	return products, nil
}

func (r *ProductRepo) Save(product *models.Product) error {
	err := r.db.Save(&product).Error
	if err != nil {
		return fmt.Errorf("failed to save product: %w", err)
	}

	return nil
}

func (r *ProductRepo) Update(id int, product *models.Product) error {
	existingProduct := &models.Product{}
	err := r.db.First(*existingProduct, id).Error
	if err != nil {
		fmt.Printf("failed to get product by id: %d\n", id)
	}

	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.Category = product.Category
	existingProduct.Count = product.Count

	if err := r.db.Save(existingProduct).Error; err != nil {
		return fmt.Errorf("failed to save product: %w", err)
	}

	return nil
}
