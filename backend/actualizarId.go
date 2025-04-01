package main

import (
        "encoding/json"
        "log"
        "net/http"
	"github.com/go-chi/chi/v5"
)

func UpdateSerieID(w http.ResponseWriter, r *http.Request) {
	// Declaro la variable para recibir la serie que viene en el cuerpo de la petición
        var req Series

	//Obtener el 'id' de la serie desde la URL usando el router 'chi'
	serieID := chi.URLParam(r,"id")

	//Si no se encuentra el id en la URL, retorno un error
	if serieID == "" {
                respondWithError(w, "Se requiere un id", http.StatusBadRequest)
                return
        }

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

	// Realizo la actualización de la serie en la base de datos
        result := db.Exec(
                "UPDATE simple_login.series SET title = ?,status=?,last_episode_watched=?,total_episodes=?,ranking=?  WHERE id = ?",
		req.Title, req.Status, req.Last, req.Total, req.Ranking, serieID)

	// Verifico si hubo algún error al ejecutar la consulta
        if result.Error != nil {
                log.Println("Ocurrio un error la serie: ", result.Error)
                respondWithError(w, "No se pudo obtener la serie", http.StatusInternalServerError)
                return
	}

	// Registro en los logs la actualización realizada
        log.Printf("Insertando serie: %+v\n", req)

	// Retorno una respuesta de éxito al frontend
        respondWithJSON(w, ApiResponse{
                Success:true,
                Message: "Serie Registrada",
        })
}
