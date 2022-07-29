package domain

import "time" // Dipakai untuk CreatedAt & UpdatedAt

// Struct Untuk Products
type Products struct {
	ID          int
	Name        string //Nama Produk
	Price       int
	Qty         int //Jumlah Produk
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
