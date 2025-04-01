package main

import (
        "net/http"
        "log"
        "github.com/go-chi/chi/v5"

)

func DeleteSerieById(w http.ResponseWriter, r *http.Request){
        var serie Series

        serieID := chi.URLParam(r,"id")
        if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }

        result := db.Raw(
                "DELETE FROM simple_login.series WHERE id = ?",serieID).Scan(&serie)

        if result.Error != nil {
                log.Println("Ocurrio un error la serie: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
        }

        respondWithJSON(w, ApiResponse{
                Success: true,
                Message: "Serie Eliminada",
             
        })

}
