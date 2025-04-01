package main

// Definir la estructura para representar una serie
type Series struct {
	ID	uint	`gorm:"column:id;primaryKey" json:"id"`
	Title   string  `gorm:"column:title;notnull"    json:"title"`
	Status  string  `gorm:"column:status;notnull"    json:"status"`
	Last    int    `gorm:"column:last_episode_watched;notnull"    json:"lastEpisodeWatched"`
	Total   int    `gorm:"column:total_episodes;notnull"    json:"totalEpisodes"`
	Ranking int    `gorm:"column:ranking" json:"ranking"`
}

// Estructura para la respuesta estándar de la API (con una serie)
type ApiResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Series   *Series `json:"series,omitempty"`

}

// Estructura para la respuesta estándar de la API con atributos genéricos
type ApiResponseAtributos struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Data    interface{} `json:"data"` // Aquí agregamos un campo 'Data' para almacenar cualquier tipo de datos
}
