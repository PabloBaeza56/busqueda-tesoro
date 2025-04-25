package Model

import "github.com/gorilla/mux"

func InicializarPreguntas(subrouter *mux.Router) {
	preguntas := []ContenidoPregunta{
		{
			URLsecreta:        "a9XfB72Lqe5Y",
			PreguntaTexto:     "Nombre del primer rector de la UADY.",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Dr. Eduardo Urzaiz Rodríguez",
		},
		{
			URLsecreta:        "M7pL2nXv04Re",
			PreguntaTexto:     "Nombre completo del campus en el que se encuentra la Facultad de Educación.",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Campus de Ciencias Sociales Económico-Administrativas y Humanidades",
		},
		{
			URLsecreta:        "V39dLpTs91qF",
			PreguntaTexto:     "Verdadero o falso: La Facultad de Educación es una de las más antiguas en la UADY.",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Falso",
		},
		{
			URLsecreta:        "Kz81NpTfR02a",
			PreguntaTexto:     "¿En qué año se formalizó la creación de la FEDU? (2 puntos extra si das la fecha exacta)",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "1984",
		},
		{
			URLsecreta:        "Lx4MeN28vQoZ",
			PreguntaTexto:     "¿Cuál es el nombre de la primera directora?",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Ing. Julia González Cervera",
		},
		{
			URLsecreta:        "Px93AjWfLmT7",
			PreguntaTexto:     "¿Cuántas personas han sido directores de la FEDU?",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "7",
		},
		{
			URLsecreta:        "Ue28Xp7KsQwB",
			PreguntaTexto:     "¿Cuál es el nombre completo del actual director?",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Dr. Pedro José Canto Herrera",
		},
		{
			URLsecreta:        "Zn45TcBvM92d",
			PreguntaTexto:     "¿Actualmente, qué función tiene el anterior edificio de la Facultad de Educación?",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Centro Institucional de Lenguas",
		},
		{
			URLsecreta:        "Jt3vL8NrXpQ1",
			PreguntaTexto:     "¿Desde qué año se imparte la LEII?",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "2005",
		},
		{
			URLsecreta:        "Fy9NaM2rZwQ8",
			PreguntaTexto:     "¿Desde qué año se imparte la LE?",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "1984",
		},
		{
			URLsecreta:        "Bp6qAfT0YuJx",
			PreguntaTexto:     "¿Cuáles son los tres posgrados que se imparten en la FEDU?",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Maestría en docencia, Maestría en innovación educativa, Maestría en investigación",
		},
		{
			URLsecreta:        "Dx1VpQfZ7rLy",
			PreguntaTexto:     "¿En qué año se creó la UMT?",
			MultimediaURL:     "https://goo.su/pObEnd",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "2000",
		},
		{
			URLsecreta:        "tL4cjnZGmCe3",
			PreguntaTexto:     "Sustantivo, animal, mascota de la Facultad de Educación:",
			MultimediaURL:     "https://cdn-icons-png.flaticon.com/512/44/44091.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "hormiga",
		},
		{
			URLsecreta:        "UpXMf3Af0rLC",
			PreguntaTexto:     "Después de 2 siglos Mozart estrenó una pieza llena de misterios, los secretos más oscuros solo se revelan a quien escucha lento y con atención",
			MultimediaURL:     "https://www.youtube.com/embed/f097MHp9TLQ",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "Amadeus",
		},
		{
			URLsecreta:        "PhLINcMyK1oB",
			PreguntaTexto:     "La iluminación llega si sabes observar detenidamente y leer entre líneas:",
			MultimediaURL:     "https://i.ibb.co/QdcNpLG/C.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "CREA EL CAMBIO",
		},
		{
			URLsecreta:        "npU2UDb7lcz1",
			PreguntaTexto:     "A veces las respuestas están justo frente a tus ojos, el arte es ver desde otra perspectiva",
			MultimediaURL:     "https://i.ibb.co/JszjdPZ/L.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "NO_TE_RINDAS",
		},
		{
			URLsecreta:        "Fj21hTVqknuH",
			PreguntaTexto:     "Una noticia del año 2027 es tu siguiente pista: <a href=\"https://drive.google.com/uc?export=download&id=1k2rZdmcPG_9mbOH6SJt-JMfGoKFSb782\" target=\"_blank\">Ver</a>",
			MultimediaURL:     "https://uvn-brightspot.s3.amazonaws.com/assets/vixes/btg/curiosidades.batanga.com/files/Es-posible-viajar-a-traves-del-tiempo-6.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "CREER",
		},
		{
			URLsecreta:        "PhvePOyte9DU",
			PreguntaTexto:     "¿piedra, papel o tijera?, prueba tu suerte, si fallas no lo intentes nuevamente (y no hables de esto)",
			MultimediaURL:     "https://elprofehora.com/wp-content/uploads/2024/05/Imagen-ilustrativa-Variaciones-de-un-juego-popular-4.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "tijera",
		},
		{
			URLsecreta:        "s0qGdorHbxzI",
			PreguntaTexto:     "He hackeado la página para ayudarte, te regalo la clave, pero sin el <i>regalo</i> (? : Cr0eN7RAgS3aN1l4o",
			MultimediaURL:     "https://www.campusciberseguridad.com/media/k2/items/cache/9267284e7733f4bec00d2e114d3f3ba1_L.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "C0N7RAS3N14",
		},
		{
			URLsecreta:        "aTiYVLoePNJB",
			PreguntaTexto:     "Puedes verlo, si lo conoces escribe su nombre... (usa minusculas)",
			MultimediaURL:     "https://i.ibb.co/Qd2fvcy/P2.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "pedro",
		},
	}

	for _, pregunta := range preguntas {
		GenerarPreguntas(pregunta, subrouter)
	}
}
