package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"time"
)

type chiRouter struct {
	router *chi.Mux
}

//NewChiRouter constructor function
func NewChiRouter() Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Use(middleware.Timeout(60 * time.Second))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	router.Use(c.Handler)

	s := &chiRouter{}
	s.router = router
	return s
}

//GET handler
func (s *chiRouter) GET(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	s.router.Get(uri, f)
}

//POST handler
func (s *chiRouter) POST(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	s.router.Post(uri, f)
}

//GET handler
func (s *chiRouter) PUT(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	s.router.Put(uri, f)
}

//POST handler
func (s *chiRouter) DELETE(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	s.router.Delete(uri, f)
}

//SERVE function creates a listener
func (s *chiRouter) SERVE(port string) {
	log.Println("Server listening on port: 8080")
	err := http.ListenAndServe(port, s.router)
	if err != nil {
		log.Println(err.Error())
	}
}
