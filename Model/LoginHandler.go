package Model

import (
	"net/http"
	Database "tresure-hunt/Database"
)

type UserData struct {
	Username       string `bson:"username"`
	ContactMethod1 string `bson:"contact_method_1"`
	ContactMethod2 string `bson:"contact_method_2"`
	Puntos         int    `bson:"puntos"`
	Preguntas      map[string]bool
}

const (
	NOMBRE_USUARIO    = "username"
	METODO_CONTACTO_1 = "contact_method_1"
	METODO_CONTACTO_2 = "contact_method_2"
)

func HomeHandlerFormulario(preguntas []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}

		username := r.FormValue(NOMBRE_USUARIO)
		contactMethod1 := r.FormValue(METODO_CONTACTO_1)
		contactMethod2 := r.FormValue(METODO_CONTACTO_2)

		userData := UserData{
			Username:       username,
			ContactMethod1: contactMethod1,
			ContactMethod2: contactMethod2,
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
		database.CrearNuevoUsuario(userData)

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/ranking", http.StatusSeeOther)
	}
}
