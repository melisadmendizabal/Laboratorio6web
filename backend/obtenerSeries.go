package main

import (
	"net/http"
	"log"

)

func GetAllSeries(w http.ResponseWriter, r *http.Request){
	var series []Series

	result := db.Raw("SELECT * FROM simple_login.series").Scan(&series)

	if result.Error != nil {
		log.Println("Ocurrio un error con todas las series: ", result.Error)
		respondWithError(w, "No se pudo obtener las series", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w,series)



}
