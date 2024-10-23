package Model

import "github.com/gorilla/mux"

func InicializarPreguntas(subrouter *mux.Router) {
	preguntas := []ContenidoPregunta{
		{
			URLsecreta:        "pregunta3",
			PreguntaTexto:     "¿Cómo se llama este lenguaje de programación?",
			MultimediaURL:     "https://th.bing.com/th/id/OIP.IQ9Q_hzG39tKELQhYxiOLQAAAA?w=400&h=400&rs=1&pid=ImgDetMain",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Go",
		},
		{
			URLsecreta:        "pregunta4",
			PreguntaTexto:     "¿Quién es el cantante?",
			MultimediaURL:     "https://youtu.be/E4WlUXrJgy4",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "Ricky",
		},
	}

	for _, pregunta := range preguntas {
		GenerarPreguntas(pregunta, subrouter)
	}
}
