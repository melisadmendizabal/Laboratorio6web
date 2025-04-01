package main

import (
        "encoding/json"
        "log"
        "net/http"
        "github.com/go-chi/chi/v5"
)

func StatusSerieID(w http.ResponseWriter, r *http.Request) {
        
        //Obtener la serie por el id con chi de la url
        serieID := chi.URLParam(r,"id")

        //Si no se encuentra el id, aunque dudo que pase por el frontend
        if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }

	var statusUpdate struct {
		Status string `json:"status"`
	}

	// Decodificar el cuerpo de la petición
        if err := json.NewDecoder(r.Body).Decode(&statusUpdate); err != nil {
        	respondWithError(w, "Formato inválido", http.StatusBadRequest)
      		 return
        }

        result := db.Exec(
                "UPDATE simple_login.series SET status=? WHERE id = ?",
                statusUpdate.Status, serieID)

        if result.Error != nil {
                log.Println("Ocurrio un error la serie: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
        }

	var serie Series
        // Usamos un SELECT para obtener la serie actualizada
        db.Raw("SELECT * FROM simple_login.series WHERE id = ?", serieID).Scan(&serie)
        log.Printf("Estado de la serie ID %s actualizado a: %s\n", serieID, statusUpdate.Status)

        respondWithJSON(w, ApiResponseAtributos{
                Success:true,
                Message: "Estado actualizado",
		Data:    serie,
        })
}
