package Model

import (
	"net/http"
	"text/template"
	Database "tresure-hunt/Database"
)

var Templates = template.Must(template.ParseGlob("Vista/html/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("userData")
	if err == nil {
		MostrarModal(w, "Usuario ya creado", "El camino no es por aquí...", "/ranking")
		return
	}

	Templates.ExecuteTemplate(w, "Registro", nil)
}

func Top10RanckingHandler(w http.ResponseWriter, r *http.Request) {
	database := Database.GetInstanceDatabase()
	data, _ := database.ObtenerTop10Ranking()
	Templates.ExecuteTemplate(w, "Ranking", data)
}

func CompleteRankingHandler(w http.ResponseWriter, r *http.Request) {
	database := Database.GetInstanceDatabase()
	data, _ := database.ObtenerRankingCompleto()
	Templates.ExecuteTemplate(w, "Ranking", data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "NotFound", nil)
}
