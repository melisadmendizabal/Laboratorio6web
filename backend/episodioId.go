package main

import (

        "log"
        "net/http"
        "github.com/go-chi/chi/v5"
)

func EpisodeSerieID(w http.ResponseWriter, r *http.Request) {

        // Obtengo el ID de la serie desde la URL usando chi
        serieID := chi.URLParam(r,"id")

        // Verifico si el ID está vacío, lo que indicaría que no se ha proporcionado el IDVerifico si el ID está vacío, lo que indicaría que no se ha proporcionado el ID
        if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }

	// Realizo la actualización del episodio de la serie, incrementando el número del último episodio visto
        result := db.Exec(
                "UPDATE simple_login.series SET last_episode_watched=last_episode_watched +1   WHERE id = ?",serieID,
	)

	// Si hay un error al ejecutar la consulta de actualización, lo manejo adecuadamente
        if result.Error != nil {
                log.Println("Ocurrio un error con actualizar la base de datos: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
        }

        var serie Series
        //SELECT para obtener la serie actualizada
        db.Raw("SELECT * FROM simple_login.series WHERE id = ?", serieID).Scan(&serie)
	log.Println(" Serie actualizada en la BD:", serie)

	// Retorno la respuesta con los datos de la serie actualizada
	respondWithJSON(w, ApiResponseAtributos{
        	Success: true,
        	Message: "Episodio actualizado",
        	Data:    serie,
	})        

}
