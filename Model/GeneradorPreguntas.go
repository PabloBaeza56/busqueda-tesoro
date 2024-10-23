package Model

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TipoMultimedia int

const (
	TipoImagen TipoMultimedia = iota
	TipoVideo
	TipoTexto
)

type ContenidoPregunta struct {
	URLsecreta        string
	PreguntaTexto     string
	MultimediaURL     string
	RespuestaCorrecta string
	Tipo              TipoMultimedia
}

func GenerarPreguntas(pf ContenidoPregunta, subrouter *mux.Router) {
	subrouter.HandleFunc("/"+pf.URLsecreta, pf.PreguntaHandler).Methods("GET")
	subrouter.HandleFunc("/"+pf.URLsecreta, pf.RespuestaHandler).Methods("POST")
}

func (pf *ContenidoPregunta) PreguntaHandler(w http.ResponseWriter, r *http.Request) {
	estaRespondida := verificarPreguntaRespondida(w, r)
	if estaRespondida {
		MostrarModal(w, "Pregunta ya respondida", "Sigue buscando...", "/ranking")
	} else {

		cookieManager := &CookieManager{w, r}

		data := struct {
			NombreUsuario string
			Pregunta      string
			MultimediaURL string
		}{
			NombreUsuario: cookieManager.leerNombreUsuario(),
			Pregunta:      pf.PreguntaTexto,
			MultimediaURL: pf.MultimediaURL,
		}

		var templateName string
		switch pf.Tipo {
		case TipoImagen:
			templateName = "Pregunta-Imagen"
		case TipoVideo:
			templateName = "Pregunta-Video"
		case TipoTexto:
			templateName = "Pregunta-Texto"
		default:
			http.Error(w, "Tipo de pregunta no soportado", http.StatusBadRequest)
			return
		}

		Templates.ExecuteTemplate(w, templateName, data)

	}
}

func (pf *ContenidoPregunta) RespuestaHandler(w http.ResponseWriter, r *http.Request) {
	manejadorRespuestas(pf.RespuestaCorrecta, w, r)
}
