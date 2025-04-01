package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleNewSeries(w http.ResponseWriter, r *http.Request) {

	var req Series
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		respondWithError(w, "Formato invalido", http.StatusBadRequest)
		return
	}

	if req.Title == ""  {
		respondWithError(w, "Ingrese una serie", http.StatusBadRequest)
		return

	}

	log.Printf("Insertando serie: %+v\n", req)

	queryResult := db.Exec(
		"INSERT INTO simple_login.series(title,status,last_episode_watched,total_episodes,ranking) VALUES (?,?,?,?,?)",
		req.Title, req.Status, req.Last, req.Total, req.Ranking)

	if queryResult.Error != nil {
		log.Println("Algo salio mal creando la serie: ", queryResult.Error)
		respondWithError(w, "Error al crear serie, criterios no apropiados o base de datos no conectada", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, ApiResponse{
		Success:true,
		Message: "Serie Registrada",
	})
}
