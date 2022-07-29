package domain

import "time" // Dipakai untuk CreatedAt & UpdatedAt

// Struct Untuk User
type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
