package api

import (
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Controller receives calls from the router
// Here is where you write your application code.
// Add your handlers and connect them to the router in the routes.go file
type Controller struct {
	Repository Repository
}

// Test this is a test handler, it displays a test message in json format
// for testing purposes it doesn't connect to a database
func (c *Controller) Test(w http.ResponseWriter, r *http.Request) {
	//Prepare the response
	response := make(map[string]interface{})
	response["message"] = "Congratulations! You made it!"

	render.JSON(w, r, response)
}

// GetUsers gets users from the database
// you will need to export CONNECTION_STRING for this to work. see README.md
func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {

	results := c.Repository.GetUsers()
	render.JSON(w, r, results)
}

// GetProduct gets a particular product from the database.
// this is a mock method, you will need to implement it yourself
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {

	//Read product Id from the route /product/{id:[0-9]+}
	idStr := chi.URLParam(r, "id")

	//Convert id to int and ignore err
	id, _ := strconv.ParseInt(idStr, 10, 64)

	//Do some cool stuff for your API
	//YOUR CODE HERE

	//Prepare the response
	response := make(map[string]interface{})
	response["message"] = "Success"
	response["id"] = id

	render.JSON(w, r, response)
}

// GetProducts gets products from the database.
// this is a mock method, you will need to implement it yourself
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {

	//Read query params from the route /products?start=0&count=50
	start := r.FormValue("start")
	count := r.FormValue("count")

	//Do some cool stuff for your API
	//YOUR CODE HERE

	//Prepare the response
	response := make(map[string]interface{})
	response["start"] = start
	response["count"] = count

	render.JSON(w, r, response)
}

// Login returns access_token if correct username and password are given.
// this is a mock method, you will need to implement it yourself
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {

	//Normally you will read these values from a post body
	username := "test"
	//password := "testpass"

	//YOUR CODE HERE
	//Step 1. Verify that the username and password are correct
	//Step 2. If incorrect send the response with the appropriate HTTP status code
	//Step 3. If correct generate a token and send it to the user

	//Generate JWT Token
	// NOTE: Don't add password to the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      EpochNow(),                             //issued at NOW!
		"exp":      ExpireIn(time.Duration(3) * time.Hour), //expires in 3 hours
	})

	//Sign the Token
	tokenString, error := token.SignedString(Config.SignKey)
	if error != nil {
		log.Println(error)
	}

	//Prepare the response
	response := make(map[string]string)
	response["access_token"] = tokenString
	render.JSON(w, r, response) // Return some demo response
}
