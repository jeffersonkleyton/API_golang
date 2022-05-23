package routes

import (
	"testando/api/controllers"

	"github.com/gorilla/mux"
)

/* func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		h.ServeHTTP(w, r)
	})
} */
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/host").Subrouter()
	s.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	s.Use(LoggingMiddleware)
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user", controllers.NewUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	return r
}
