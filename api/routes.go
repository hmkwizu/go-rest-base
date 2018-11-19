package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var controller = &Controller{Repository: Repository{}}

var tokenAuth *jwtauth.JWTAuth

// Routes creates the routes for our application
// Add your application routes below
func Routes() *chi.Mux {

	tokenAuth = jwtauth.New("HS256", Config.SignKey, nil)

	r := chi.NewRouter()

	//Public routes - Add unsecured routes here
	r.Get("/test", controller.Test)
	r.Get("/users", controller.GetUsers)
	//r.Post("/getToken", controller.GetToken) //return jwt token on success
	r.Get("/getToken", controller.GetToken) //return jwt token on success

	// Protected routes - Add routes that you want to secure inside this group
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens
		r.Use(jwtauth.Authenticator)

		//router.Get("/product/{id}", controller.GetExam)
		r.Get("/product/{id:[0-9]+}", controller.GetProduct) //ensure id is integer
		r.Get("/products", controller.GetProducts)
	})

	return r
}
