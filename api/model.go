package api

// Add your models to this file

// User holds information about application users
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
