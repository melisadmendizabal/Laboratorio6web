package main

import (
       
        "log"
        "net/http"
        "github.com/go-chi/chi/v5"
)

func EpisodeSerieID(w http.ResponseWriter, r *http.Request) {

        //Obtener la serie por el id con chi de la url
        serieID := chi.URLParam(r,"id")

        //Si no se encuentra el id, aunque dudo que pase por el frontend
        if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }


        
        result := db.Exec(
                "UPDATE simple_login.series SET last_episode_watched=last_episode_watched +1   WHERE id = ?",serieID,
	)

        if result.Error != nil {
                log.Println("Ocurrio un error con actualizar la base de datos: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
        }

        var serie Series
        // Usamos un SELECT para obtener la serie actualizada
        db.Raw("SELECT * FROM simple_login.series WHERE id = ?", serieID).Scan(&serie)
	log.Println(" Serie actualizada en la BD:", serie)

	respondWithJSON(w, ApiResponseAtributos{
        	Success: true,
        	Message: "Episodio actualizado",
        	Data:    serie,
	})        

}
