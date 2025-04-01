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
	// Defino la cadena de conexión a la base de datos MySQL
	dns := "app_user:app_password@tcp(localhost:3306)/simple_login?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	 // Conecto a la base de datos usando GORM
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Problemas", err)
	}

	// Configuración del enrutador con Chi
	r := chi.NewRouter()

	// Uso el middleware Logger para loguear cada solicitud
        r.Use(middleware.Logger)

	// Configuro CORS para permitir peticiones desde cualquier origen
	r.Use(cors.Handler(cors.Options{
                AllowedOrigins:   []string{"*"},
                AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
                AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
                AllowCredentials: true,
        }))

	// Defino las rutas y sus respectivos controladores
	r.Post("/api/series", HandleNewSeries)
	r.Get("/api/series", GetAllSeries)
	r.Get("/api/series/{id}", GetSerieById)
	r.Delete("/api/series/{id}", DeleteSerieById)
	r.Put("/api/series/{id}",UpdateSerieID)
	r.Patch("/api/series/{id}/status", StatusSerieID)
	r.Patch("/api/series/{id}/episode", EpisodeSerieID)
	r.Patch("/api/series/{id}/upvote", UpRankingSeries)
	r.Patch("/api/series/{id}/downvote", DownRankingSeries)

	// Inicia el servidor en el puerto 8080 y maneja las solicitudes con el enrutador
	if err := http.ListenAndServe(":8080", r); err != nil {
 	   log.Fatal("Error al iniciar el servidor: ", err)
	}
}
