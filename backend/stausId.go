package main

import (
        "encoding/json"
        "log"
        "net/http"
        "github.com/go-chi/chi/v5"
)

func StatusSerieID(w http.ResponseWriter, r *http.Request) {

        // Obtener el id de la serie desde la URL con chi
        serieID := chi.URLParam(r,"id")

        // Verificar que el id no esté vacío
        if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }

	// Definimos la estructura para actualizar el estado
	var statusUpdate struct {
		Status string `json:"status"`
	}

	// Decodificar el cuerpo de la petición JSON
        if err := json.NewDecoder(r.Body).Decode(&statusUpdate); err != nil {
        	respondWithError(w, "Formato inválido", http.StatusBadRequest)
      		 return
        }

	// Ejecutamos la actualización del estado en la base de datos
        result := db.Exec(
                "UPDATE simple_login.series SET status=? WHERE id = ?",
                statusUpdate.Status, serieID)

	// Si ocurre un error en la actualización, respondemos con un error
        if result.Error != nil {
                log.Println("Ocurrio un error la serie: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
        }

	var serie Series
        // SELECT para obtener la serie actualizada
        db.Raw("SELECT * FROM simple_login.series WHERE id = ?", serieID).Scan(&serie)
        log.Printf("Estado de la serie ID %s actualizado a: %s\n", serieID, statusUpdate.Status)

	// Respondemos con la serie actualizada y un mensaje de éxito
        respondWithJSON(w, ApiResponseAtributos{
                Success:true,
                Message: "Estado actualizado",
		Data:    serie,
        })
}
