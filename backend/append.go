package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleNewSeries(w http.ResponseWriter, r *http.Request) {
	// Declaro la variable para recibir los datos de la nueva serie que vienen en el cuerpo de la petición
	var req Series

	// Decodifico el cuerpo de la petición en la estructura 'req'
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		respondWithError(w, "Formato invalido", http.StatusBadRequest)
		return
	}

	// Verifico que el título de la serie no esté vacío
	if req.Title == ""  {
		respondWithError(w, "Ingrese una serie", http.StatusBadRequest)
		return

	}

	// Registro en los logs la serie que se está intentando insertar
	log.Printf("Insertando serie: %+v\n", req)

	// Realizo la inserción de la nueva serie en la base de datos// Realizo la inserción de la nueva serie en la base de datos
	queryResult := db.Exec(
		"INSERT INTO simple_login.series(title,status,last_episode_watched,total_episodes,ranking) VALUES (?,?,?,?,?)",
		req.Title, req.Status, req.Last, req.Total, req.Ranking)

	// Verifico si hubo algún error al ejecutar la consulta
	if queryResult.Error != nil {
		log.Println("Algo salio mal creando la serie: ", queryResult.Error)
		respondWithError(w, "Error al crear serie, criterios no apropiados o base de datos no conectada", http.StatusInternalServerError)
		return
	}

	// Retorno una respuesta de éxito al frontend
	respondWithJSON(w, ApiResponse{
		Success:true,
		Message: "Serie Registrada",
	})
}
