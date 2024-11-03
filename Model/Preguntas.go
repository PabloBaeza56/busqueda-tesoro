package Model

import "github.com/gorilla/mux"

func InicializarPreguntas(subrouter *mux.Router) {
	preguntas := []ContenidoPregunta{
		{
			URLsecreta:        "Qv63cRZKDv0B",
			PreguntaTexto:     "Pídele ayuda a César para descifrarlo: nytbevgzb_qry_pnzovb",
			MultimediaURL:     "https://www.gub.uy/centro-nacional-respuesta-incidentes-seguridad-informatica/sites/centro-nacional-respuesta-incidentes-seguridad-informatica/files/imagenes/noticias/Cifrado_720x350.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "algoritmo_del_cambio",
		},
		{
			URLsecreta:        "tL4cjnZGmCe3",
			PreguntaTexto:     "Sustantivo, animal, mascota de la Facultad de Matemáticas:",
			MultimediaURL:     "https://cdn-icons-png.flaticon.com/512/44/44091.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "lobo",
		},
		{
			URLsecreta:        "vae212t2uRJV",
			PreguntaTexto:     "¿Cómo se llama el actual director de la facultad? (nombre completo)",
			MultimediaURL:     "https://fastly.4sqi.net/img/general/600x600/YKHDyTPS-wazfy8ZOid8oF49X23jqzGhczQhO64MyWQ.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "Ernesto Antonio Guerrero Lara",
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
			URLsecreta:        "iAuQ4BVIUbua",
			PreguntaTexto:     "La iluminación llega si sabes observar detenidamente y leer entre líneas:",
			MultimediaURL:     "https://i.ibb.co/3pNwp2P/C2.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "FMAT UNIDA",
		},
		{
			URLsecreta:        "npU2UDb7lcz1",
			PreguntaTexto:     "A veces las respuestas están justo frente a tus ojos, el arte es ver desde otra perspectiva",
			MultimediaURL:     "https://i.ibb.co/JszjdPZ/L.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "NO_TE_RINDAS",
		},
		{
			URLsecreta:        "gNhK6XCD3xEP",
			PreguntaTexto:     "Vamos a dar un paseo: 29X4+9G Mérida, Yucatán, llegaremos a la __________",
			MultimediaURL:     "https://www.youtube.com/embed/Zcf5nyappYk?si=iFc1ay6exxJJ1AMG",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "biblioteca",
		},
		{
			URLsecreta:        "ylHrFR0vNteY",
			PreguntaTexto:     "Sólo PIensa, y PIensa y PIensa, esa es la PIsta: 141592_____897932384626",
			MultimediaURL:     "https://www.youtube.com/embed/3HRkKznJoZA?si=mHt706_bxtAO6JX6",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "6535",
		},
		{
			URLsecreta:        "zSWuYc4JSj8r",
			PreguntaTexto:     "Escribe en número la forma de que al agregarle un uno al veinte, este sea un número más pequeño. (piensa afuera de la caja, en mayúsculas)",
			MultimediaURL:     "https://www.youtube.com/embed/E4WlUXrJgy4",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "XIX",
		},
		{
			URLsecreta:        "ac52LRUFvcrq",
			PreguntaTexto:     "Una especie de otro planeta envió este video a nuestros servidores, ¿Qué materia es su preferida?",
			MultimediaURL:     "https://www.youtube.com/embed/83gdQe0Ij5k?si=Ry7KobWMEztStHXf",
			Tipo:              TipoVideo,
			RespuestaCorrecta: "trigonometría",
		},
		{
			URLsecreta:        "Fj21hTVqknuH",
			PreguntaTexto:     "Una noticia del año 2027 es tu siguiente pista: <a href=\"https://drive.google.com/uc?export=download&id=1k2rZdmcPG_9mbOH6SJt-JMfGoKFSb782\" target=\"_blank\">Ver</a>",
			MultimediaURL:     "https://uvn-brightspot.s3.amazonaws.com/assets/vixes/btg/curiosidades.batanga.com/files/Es-posible-viajar-a-traves-del-tiempo-6.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "CREER",
		},
		{
			URLsecreta:        "puNhgatFVuZJ",
			PreguntaTexto:     "¿Qué número sigue la serie: 1, 11, 21, 1211, 111221, 312211,...?",
			MultimediaURL:     "https://siadcon.com/wp-content/uploads/2022/07/la-logica.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "13112221",
		},
		{
			URLsecreta:        "PhvePOyte9DU",
			PreguntaTexto:     "¿piedra, papel o tijera?, prueba tu suerte, si fallas no lo intentes nuevamente (y no hables de esto)",
			MultimediaURL:     "https://elprofehora.com/wp-content/uploads/2024/05/Imagen-ilustrativa-Variaciones-de-un-juego-popular-4.png",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "tijera",
		},
		{
			URLsecreta:        "eej4jiSYn1fE",
			PreguntaTexto:     "escribe en números la cantidad de veces que se puede restar 1 a la cifra 1111",
			MultimediaURL:     "https://img.freepik.com/vector-gratis/ilustracion-fondo-joker_1284-8531.jpg",
			Tipo:              TipoImagen,
			RespuestaCorrecta: "1",
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
