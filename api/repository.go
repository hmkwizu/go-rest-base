package api

// Repository queries the database and returns results to the controller
// Add your methods here for the handler functions your create in the controller
type Repository struct{}

// GetUsers gets users from the database. see README.md
func (r Repository) GetUsers() []User {

	// select
	rows, err := db.Query("select id,email from users")
	CheckError(err)

	results := make([]User, 0, 10)
	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Email)
		CheckError(err)

		results = append(results, user)

	}

	return results
}
