package main

import (
        "net/http"
        "log"
        "github.com/go-chi/chi/v5"

)

func DeleteSerieById(w http.ResponseWriter, r *http.Request){
        var serie Series

	// Obtengo el ID de la serie desde la URL usando chi
        serieID := chi.URLParam(r,"id")

	// Verifico si el ID está vacío, lo que indicaría que no se ha proporcionado el ID
        if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }

	// Realizo la eliminación de la serie usando el ID proporcionado// Realizo la eliminación de la serie usando el ID proporcionado
        result := db.Raw(
                "DELETE FROM simple_login.series WHERE id = ?",serieID).Scan(&serie)

	// Si hay un error al ejecutar la consulta, lo manejo adecuadamente
        if result.Error != nil {
                log.Println("Ocurrio un error la serie: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
        }

	// Retorno una respuesta de éxito si la serie fue eliminada correctamente
        respondWithJSON(w, ApiResponse{
                Success: true,
                Message: "Serie Eliminada",
             
        })

}
