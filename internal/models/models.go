package models

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Date    string `json:"date"`
	Role    string `json:"role"`
}

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Count       int    `json:"count"`
	Date        string `json:"date"`
}

type Orders struct {
	Id      int     `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
	Price   string  `json:"price"`
	Date    string  `json:"date"`
	Status  string  `json:"status"`
}

type Payments struct {
	Id     int    `json:"id"`
	User   User   `json:"user"`
	Order  Orders `json:"orders"`
	Sum    string `json:"sum"`
	Date   string `json:"date"`
	Status string `json:"status"`
}
