package Model

import (
	"net/http"
	Database "tresure-hunt/Database"
	"unicode"
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
	if normalizarString(respuestaRecibida) == normalizarString(respuestaEsperada) {
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

func normalizarString(input string) string {
	var normalized string

	for _, r := range input {
		// Convertir a minúscula
		lower := unicode.ToLower(r)

		// Remover acentos
		switch lower {
		case 'á', 'à', 'â', 'ä', 'ã':
			normalized += "a"
		case 'é', 'è', 'ê', 'ë':
			normalized += "e"
		case 'í', 'ì', 'î', 'ï':
			normalized += "i"
		case 'ó', 'ò', 'ô', 'ö', 'õ':
			normalized += "o"
		case 'ú', 'ù', 'û', 'ü':
			normalized += "u"
		case 'ñ':
			normalized += "n"
		default:
			normalized += string(lower)
		}
	}

	return normalized
}

func calcularPuntos(id_pregunta string) int {
	database := Database.GetInstanceDatabase()
	numero, _ := database.NumeroPersonasQueRespondieron(id_pregunta)

	switch numero {
	case 0:
		return 100
	case 1:
		return 95
	case 2:
		return 90
	case 3:
		return 85
	case 4:
		return 80
	case 5:
		return 75
	case 6:
		return 70
	case 7:
		return 65
	case 8:
		return 60
	case 9:
		return 55
	case 10:
		return 50
	default:
		return 30
	}

}
