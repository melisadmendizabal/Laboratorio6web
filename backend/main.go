package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
)

var db *gorm.DB

func main() {
	dns := "app_user:app_password@tcp(localhost:3306)/simple_login?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Problemas", err)
	}



	r := chi.NewRouter()
        r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
                AllowedOrigins:   []string{"*"},
                AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
                AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
                AllowCredentials: true,
        }))

	r.Post("/api/series", HandleNewSeries)
	r.Get("/api/series", GetAllSeries)
	r.Get("/api/series/{id}", GetSerieById)
	r.Delete("/api/series/{id}", DeleteSerieById)
	r.Put("/api/series/{id}",UpdateSerieID)
	r.Patch("/api/series/{id}/status", StatusSerieID)
	r.Patch("/api/series/{id}/episode", EpisodeSerieID)
	r.Patch("/api/series/{id}/upvote", UpRankingSeries)
	r.Patch("/api/series/{id}/downvote", DownRankingSeries)
	if err := http.ListenAndServe(":8080", r); err != nil {
 	   log.Fatal("Error al iniciar el servidor: ", err)
	}
}
