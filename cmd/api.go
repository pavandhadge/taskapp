package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/pavandhadge/taskapp/internal/repository"
)

type application struct {
	config dbconfig
	addr   string
	store  repository.TaskRepo
}

type dbconfig struct {
	dbToken      string
	dbUrl        string
	maxopenConn  int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// r.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins: []string{"*"}, // Allow all origins (replace with specific domains if needed)
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
	// 	AllowedHeaders: []string{
	// 		"Accept",
	// 		"Authorization",
	// 		"Content-Type",
	// 		"X-CSRF-Token",
	// 		"X-Requested-With",
	// 		"X-Frame-Options",
	// 		"X-Content-Type-Options",
	// 		"X-XSS-Protection",
	// 		"Access-Control-Allow-Origin",
	// 		"Access-Control-Allow-Headers",
	// 		"Access-Control-Allow-Methods",
	// 		"Access-Control-Allow-Credentials",
	// 	},
	// 	ExposedHeaders: []string{
	// 		"Link",
	// 		"X-CSRF-Token",
	// 		"Authorization",
	// 		"Access-Control-Allow-Origin",
	// 		"Access-Control-Allow-Headers",
	// 		"Access-Control-Allow-Methods",
	// 	},
	// 	AllowCredentials: true, // Set to true if your frontend needs to send cookies/auth tokens
	// 	MaxAge:           300,  // Cache CORS results for 300 seconds
	// }))

	r.Route("/health", func(r chi.Router) {
		r.Get("/", app.healthCheckHandler)
	})
	r.Route("/v1/tasks", func(r chi.Router) {
		r.Get("/", app.getTaskHandler)
		r.Post("/", app.CreateTaskHandler)
		r.Put("/{id}", app.toggleTaskHandler)
		r.Delete("/{id}", app.deleteHandler)
	})
	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 20,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	fmt.Println("server is started at", app.addr)
	return srv.ListenAndServe()

}
