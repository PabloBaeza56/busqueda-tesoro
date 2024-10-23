package Model

import "net/http"

type ModalData struct {
	Title       string
	Message     string
	RedirectURL string
}

func MostrarModal(w http.ResponseWriter, title string, message string, redirectURL string) {

	data := ModalData{
		Title:       title,
		Message:     message,
		RedirectURL: redirectURL,
	}

	Templates.ExecuteTemplate(w, "Modal", data)
}
