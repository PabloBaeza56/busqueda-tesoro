package Model

import "github.com/gorilla/mux"

func InicializarPreguntas(subrouter *mux.Router) {
	preguntas := []ContenidoPregunta{
		{
			URLsecreta:        "pregunta1",
			PreguntaTexto:     "Pídele ayuda a César para descifrarlo: nytbevgzb_qry_pnzovb",
			MultimediaURL:     "https://www.gub.uy/centro-nacional-respuesta-incidentes-seguridad-informatica/sites/centro-nacional-respuesta-incidentes-seguridad-informatica/files/imagenes/noticias/Cifrado_720x350.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "algoritmo_del_cambio",
		},
		{
			URLsecreta:        "pregunta2",
			PreguntaTexto:     "Sustantivo, animal, mascota de la Facultad de Matemáticas:",
			MultimediaURL:     "https://cdn-icons-png.flaticon.com/512/44/44091.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "lobo",
		},
		{
			URLsecreta:        "pregunta3",
			PreguntaTexto:     "¿Cómo se llama el actual director de la facultad? (sólo 1 nombre y 1 apellido)",
			MultimediaURL:     "https://fastly.4sqi.net/img/general/600x600/YKHDyTPS-wazfy8ZOid8oF49X23jqzGhczQhO64MyWQ.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Ernesto Guerrero",
		},
		{
			//TODO
			URLsecreta:        "pregunta4",
			PreguntaTexto:     "Los secretos más oscuros solo se revelan a quien escucha con atención",
			MultimediaURL:     "https://www.youtube.com/embed/E4WlUXrJgy4",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "Ricky",
		},
		{
			//TODO: Imagen con el código
			URLsecreta:        "pregunta5",
			PreguntaTexto:     "a veces las respuestas están justo frente a tus ojos, pero no las puedes ver…",
			MultimediaURL:     "https://www.youtube.com/embed/E4WlUXrJgy4",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "Ricky",
		},
		{
			//TODO: imagen con el código
			URLsecreta:        "pregunta6",
			PreguntaTexto:     "La iluminación llega si sabes observar detenidamente en el momento correcto",
			MultimediaURL:     "https://www.youtube.com/embed/E4WlUXrJgy4",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "FMAT",
		},
		{
			//TODO: mona lisa
			URLsecreta:        "pregunta7",
			PreguntaTexto:     "Algo sobre el arte",
			MultimediaURL:     "Model/nffrgf/L.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Ricky",
		},
		{
			URLsecreta:        "pregunta8",
			PreguntaTexto:     "Vamos a dar un paseo: 29X4+9G Mérida, Yucatán, ¿sabes a dónde iremos?",
			MultimediaURL:     "https://www.youtube.com/embed/E4WlUXrJgy4",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "Ricky",
		},
		{
			URLsecreta:        "pregunta9",
			PreguntaTexto:     "Sólo PIensa, y PIensa y PIensa, esa es PIsta: 141592_____897932384626",
			MultimediaURL:     "https://www.youtube.com/embed/E4WlUXrJgy4",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "6535",
		},
		{
			URLsecreta:        "pregunta10",
			PreguntaTexto:     "Escribe en número la forma de que al agregarle un uno al veinte, este sea un número más pequeño. (piensa afuera de la caja, en mayúsculas)",
			MultimediaURL:     "https://www.youtube.com/embed/E4WlUXrJgy4",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "XIX",
		},
		{
			URLsecreta:        "pregunta11",
			PreguntaTexto:     "Una especie de otro planeta envió este video a nuestros servidores, ¿Qué materia es su preferida?",
			MultimediaURL:     "https://www.youtube.com/embed/83gdQe0Ij5k?si=Ry7KobWMEztStHXf",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "trigonometría",
		},
		{
			//TODO: noticia del canva
			URLsecreta:        "pregunta12",
			PreguntaTexto:     "Una noticia del año 2027: <a href=\"https://drive.google.com/uc?export=download&id=1k2rZdmcPG_9mbOH6SJt-JMfGoKFSb782\" target=\"_blank\">Aquí</a>",
			MultimediaURL:     "",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Ricky",
		},
		{
			URLsecreta:        "pregunta13",
			PreguntaTexto:     "¿Qué número sigue la serie: 1, 11, 21, 1211, 111221, 312211,...?",
			MultimediaURL:     " ",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "13112221",
		},
		{
			URLsecreta:        "pregunta14",
			PreguntaTexto:     "¿piedra, papel o tijera?, prueba tu suerte, si fallas no lo intentes nuevamente (y no hables de esto)",
			MultimediaURL:     "tijera",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "tijera",
		},
		{
			URLsecreta:        "pregunta15",
			PreguntaTexto:     "escribe en números la cantidad de veces que se puede restar 1 a la cifra 1111",
			MultimediaURL:     "",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "1",
		},
		{
			URLsecreta:        "pregunta16",
			PreguntaTexto:     "He hackeado la página para ayudarte, te regalo la clave, pero sin el _regalo_ (? : Cr0eN7RAgS3aN1l4o",
			MultimediaURL:     "https://www.campusciberseguridad.com/media/k2/items/cache/9267284e7733f4bec00d2e114d3f3ba1_L.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "C0N7RAS3N14",
		},
	}

	for _, pregunta := range preguntas {
		GenerarPreguntas(pregunta, subrouter)
	}
}
