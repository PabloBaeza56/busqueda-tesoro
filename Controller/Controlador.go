package Controlador

import (
	"net/http"
	Database "tresure-hunt/Database"
	modelo "tresure-hunt/Model"

	"github.com/gorilla/mux"
)

func verificarSiElUsuarioYaSeRegistro(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("userData")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type HandlerManager struct {
	Writer http.ResponseWriter
	Reader *http.Request
}

func ArranqueInicial(r *mux.Router) {

	r.NotFoundHandler = http.HandlerFunc(modelo.NotFoundHandler)

	subrouter := r.PathPrefix("/").Subrouter()

	subrouter.Use(verificarSiElUsuarioYaSeRegistro)

	r.HandleFunc("/", modelo.HomeHandler).Methods("GET")
	subrouter.HandleFunc("/ranking", modelo.RankingHandler).Methods("GET")

	modelo.InicializarPreguntas(subrouter)

	rutasPregunta := obtenerRutasPreguntas(r)

	r.HandleFunc("/", modelo.HomeHandlerFormulario(rutasPregunta)).Methods("POST")

	database := Database.GetInstanceDatabase()

	for _, pregunta := range rutasPregunta {
		database.AniadirPreguntasQueNoEstenEnLaDB(pregunta)
	}

}

func obtenerRutasPreguntas(r *mux.Router) []string {
	rutasExistentes := obtenerTodasLasRutas(r)
	elementosAEliminar := []string{"/css/", "/", "/js/", "/ranking"}

	resultado := filtrarYEliminarRepetidos(rutasExistentes, elementosAEliminar)
	return resultado
}

func obtenerTodasLasRutas(r *mux.Router) []string {
	rutas := []string{}

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}

		rutas = append(rutas, pathTemplate)

		return nil
	})

	return rutas
}

func filtrarYEliminarRepetidos(arr []string, elementosAEliminar []string) []string {

	elementosVistos := make(map[string]bool)
	elementosAEliminarMap := make(map[string]bool)

	for _, elem := range elementosAEliminar {
		elementosAEliminarMap[elem] = true
	}

	resultado := []string{}

	for _, elem := range arr {
		// Verificar si el elemento ya se ha visto o si está en la lista de eliminación
		if !elementosVistos[elem] && !elementosAEliminarMap[elem] {
			// Añadir el elemento al resultado si es único y no está en la lista de eliminación
			resultado = append(resultado, elem)
			// Marcar el elemento como visto
			elementosVistos[elem] = true
		}
	}

	return resultado
}
