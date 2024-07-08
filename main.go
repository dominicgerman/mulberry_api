package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dominicgerman/mulberry_api/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{
		DB: database.New(db),
	}

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // 5 minutes
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(corsOptions))

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healtz", handlerReadiness)
	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/tasks", apiCfg.handlerTasksGet)

	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
	v1Router.Post("/tasks", apiCfg.middlewareAuth(apiCfg.handlerTasksCreate))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Serving on port %s", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
