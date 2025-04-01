package main

import (
        "net/http"
        "log"
	"github.com/go-chi/chi/v5"
	
)

// Función para obtener una serie por su ID
func GetSerieById(w http.ResponseWriter, r *http.Request){
	var serie Series

	// Obtenemos el ID de la serie desde los parámetros de la URL
	serieID := chi.URLParam(r,"id")

	// Si no se encuentra el ID en la URL, retornamos un error
	if serieID == "" {
		respondWithError(w, "Se requiere un id", http.StatusBadRequest)
		return
	}

	// Realizamos la consulta SQL para obtener la serie con el ID proporcionado
	result := db.Raw(
		"SELECT * FROM simple_login.series WHERE id = ?",serieID).Scan(&serie)

	// Si ocurre un error en la consulta, respondemos con un error
	if result.Error != nil {
                log.Println("Ocurrio un error la serie: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
        }

	// Respondemos con la serie encontrada en formato JSON// Respondemos con la serie encontrada en formato JSON
	respondWithJSON(w,serie)

}
