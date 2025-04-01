package main

import (
	"net/http"
	"log"

)

// Función para obtener todas las series de la base de datos
func GetAllSeries(w http.ResponseWriter, r *http.Request){
	var series []Series

	// Realizamos la consulta SQL para obtener todas las series
	result := db.Raw("SELECT * FROM simple_login.series").Scan(&series)

	// Si ocurre un error en la consulta, respondemos con un error
	if result.Error != nil {
		log.Println("Ocurrio un error con todas las series: ", result.Error)
		respondWithError(w, "No se pudo obtener las series", http.StatusInternalServerError)
		return
	}

	// Respondemos con todas las series obtenidas en formato JSON
	respondWithJSON(w,series)

}
