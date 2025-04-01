package main

import (
        "encoding/json"
        "log"
        "net/http"
	"github.com/go-chi/chi/v5"
)

func UpdateSerieID(w http.ResponseWriter, r *http.Request) {
	//obtengo la base de datos
        var req Series
	//Obtener la serie por el id con chi de la url
	serieID := chi.URLParam(r,"id")

	//Si no se encuentra el id, aunque dudo que pase por el frontend
	if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
                respondWithError(w, "Formato invalido", http.StatusBadRequest)
                return
        }

        if req.Title == ""  {
                respondWithError(w, "Ingrese una serie", http.StatusBadRequest)
                return

        }


        result := db.Exec(
                "UPDATE simple_login.series SET title = ?,status=?,last_episode_watched=?,total_episodes=?,ranking=?  WHERE id = ?",
		req.Title, req.Status, req.Last, req.Total, req.Ranking, serieID)

        if result.Error != nil {
                log.Println("Ocurrio un error la serie: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
	}

        log.Printf("Insertando serie: %+v\n", req)

        respondWithJSON(w, ApiResponse{
                Success:true,
                Message: "Serie Registrada",
        })
}
