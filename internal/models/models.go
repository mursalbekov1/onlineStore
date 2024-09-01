package models

import "time"

type User struct {
	Id      int       `gorm:"primaryKey" json:"id"`
	Name    string    `gorm:"size:100" json:"name"`
	Email   string    `gorm:"size:100;unique" json:"email"`
	Address string    `gorm:"size:255" json:"address"`
	Date    time.Time `gorm:"autoCreateTime" json:"date"`
	Role    string    `gorm:"size:50" json:"role"`
}

type Product struct {
	Id          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Price       string    `gorm:"size:50" json:"price"`
	Category    string    `gorm:"size:50" json:"category"`
	Count       int       `json:"count"`
	Date        time.Time `gorm:"autoCreateTime" json:"date"`
}

type Order struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	ProductID int       `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Price     string    `gorm:"size:50" json:"price"`
	Date      time.Time `gorm:"autoCreateTime" json:"date"`
	Status    string    `gorm:"size:50" json:"status"`
}

type Payment struct {
	Id      int       `gorm:"primaryKey" json:"id"`
	UserID  int       `gorm:"not null" json:"user_id"`
	User    User      `gorm:"foreignKey:UserID" json:"user"`
	OrderID int       `gorm:"not null" json:"order_id"`
	Order   Order     `gorm:"foreignKey:OrderID" json:"order"`
	Sum     string    `gorm:"size:50" json:"sum"`
	Date    time.Time `gorm:"autoCreateTime" json:"date"`
	Status  string    `gorm:"size:50" json:"status"`
}
