package Model

import (
	"net/http"
	Database "tresure-hunt/Database"
)

func verificarPreguntaRespondida(w http.ResponseWriter, r *http.Request) bool {
	id_pregunta := r.URL.Path
	cookieManager := &CookieManager{w, r}
	respondida, _ := cookieManager.verificarSiElCampoEsVerdadero(id_pregunta)
	if respondida {
		return true
	} else {
		return false
	}
}

func manejadorRespuestas(respuestaEsperada string, w http.ResponseWriter, r *http.Request) {
	id_pregunta := r.URL.Path
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}
	respuestaRecibida := r.FormValue("respuesta")
	if respuestaRecibida == respuestaEsperada {
		cookieManager := &CookieManager{w, r}
		cookieManager.marcarPreguntaComoRespondida(id_pregunta)

		database := Database.GetInstanceDatabase()
		database.AniadirPuntosUsuario(
			cookieManager.leerNombreUsuario(),
			cookieManager.leerMetodoContacto1(),
			cookieManager.leerMetodoContacto2(),
			calcularPuntos(id_pregunta),
		)
		database.MarcarPreguntaComoRespondida(cookieManager.leerNombreUsuario(), id_pregunta)
		database.IncrementarPersonasQueRespondieron(id_pregunta)

		MostrarModal(w, "¡Respuesta correcta!", "La pregunta se respondió correctamente.", "/ranking")
	} else {
		MostrarModal(w, "Respuesta incorrecta", "Por favor, intenta de nuevo.", id_pregunta)
	}
}

func calcularPuntos(id_pregunta string) int {
	database := Database.GetInstanceDatabase()
	numero, _ := database.NumeroPersonasQueRespondieron(id_pregunta)

	switch numero {
	case 0:
		return 100
	case 1:
		return 80
	case 2:
		return 60
	case 3:
		return 40
	default:
		return 10
	}

}
