package repository

import (
	"fmt"
	"gorm.io/gorm"
	"onlineStore/internal/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) GetByID(id int) (*models.User, error) {
	user := models.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by ID %d: %w", id, err)
	}
	return &user, nil
}

func (r *UserRepo) Save(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (r *UserRepo) GetAll() ([]*models.User, error) {
	var users []*models.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func (r *UserRepo) Update(user *models.User) error {
	existingUser := &models.User{}
	if err := r.db.First(existingUser, user.Id).Error; err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Role = user.Role

	if err := r.db.Save(existingUser).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
