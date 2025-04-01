package main

import (

        "log"
        "net/http"
        "github.com/go-chi/chi/v5"
)

func DownRankingSeries(w http.ResponseWriter, r *http.Request) {

        // Obtengo el ID de la serie desde la URL usando chi
        serieID := chi.URLParam(r,"id")

        //Verifico si el ID está vacío, lo que indicaría que no se ha proporcionado el ID
        if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }

	// Realizo la actualización del ranking, disminuyendo su valor en 1
        result := db.Exec(
                "UPDATE simple_login.series SET ranking = ranking -1   WHERE id = ?",serieID,
        )

	// Si hay un error al ejecutar la consulta, lo manejo adecuadamente// Si hay un error al ejecutar la consulta, lo manejo adecuadamente
        if result.Error != nil {
                log.Println("Ocurrio un error con actualizar la base de datos: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
        }

	// Declaro una variable para almacenar los datos de la serie actualizada
        var serie Series

        // Consulto la serie actualizada de la base de datos para devolverla
        db.Raw("SELECT * FROM simple_login.series WHERE id = ?", serieID).Scan(&serie)
        log.Println(" Serie actualizada en la BD:", serie)

	// Retorno la respuesta con la serie actualizada y un mensaje de éxito
        respondWithJSON(w, ApiResponseAtributos{
                Success: true,
                Message: "Ranking Actualizado",
                Data:    serie,
        })

}
