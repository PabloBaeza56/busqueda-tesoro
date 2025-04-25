package Model

import (
	"net/http"
	Database "tresure-hunt/Database"
)

type UserData struct {
	NombreUsuario  string `bson:"NombreUsuario"`
	Licenciatura   string `bson:"Licenciatura"`
	MetodoContacto string `bson:"MetodoContacto"`
	Puntos         int    `bson:"puntos"`
	Preguntas      map[string]bool
}

const (
	NOMBRE_USUARIO  = "NombreUsuario"
	LICENCIATURA    = "Licenciatura"
	METODO_CONTACTO = "MetodoContacto"
)

func HomeHandlerFormulario(preguntas []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		nombreUsuario := r.FormValue(NOMBRE_USUARIO)
		licenciatura := r.FormValue(LICENCIATURA)
		MetodoContacto := r.FormValue(METODO_CONTACTO)

		userData := UserData{
			NombreUsuario:  nombreUsuario,
			Licenciatura:   licenciatura,
			MetodoContacto: MetodoContacto,
			Puntos:         0,
		}

		cookieManager := &CookieManager{w, r}
		userData = cookieManager.agregarPreguntas(userData, preguntas)

		cookie, err := cookieManager.crearNuevoUsuario(userData)
		if err != nil {
			http.Error(w, "Error al crear la cookie", http.StatusInternalServerError)
			return
		}

		database := Database.GetInstanceDatabase()
		// Crear usuario en la base de datos
		err = database.CrearNuevoUsuario(userData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/ranking", http.StatusSeeOther)
	}
}
