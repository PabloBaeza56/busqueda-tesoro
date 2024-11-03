package Model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const CookieName = "userData"

type CookieManager struct {
	Writer http.ResponseWriter
	Reader *http.Request
}

func (a *CookieManager) crearNuevoUsuario(userData UserData) (*http.Cookie, error) {
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	encodedData := base64.StdEncoding.EncodeToString(jsonData)

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    encodedData,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	return cookie, nil
}

func (a *CookieManager) leerNombreUsuario() string {
	cookie, _ := a.Reader.Cookie(CookieName)

	decodedData, _ := base64.StdEncoding.DecodeString(cookie.Value)

	var userData UserData
	erro := json.Unmarshal(decodedData, &userData)
	if erro != nil {
		return ""
	}

	return userData.NombreUsuario
}

func (a *CookieManager) leerMetodoContacto1() string {
	cookie, _ := a.Reader.Cookie(CookieName)

	decodedData, _ := base64.StdEncoding.DecodeString(cookie.Value)

	var userData UserData
	erro := json.Unmarshal(decodedData, &userData)
	if erro != nil {
		return ""
	}

	return userData.Licenciatura
}

func (a *CookieManager) leerMetodoContacto2() string {
	cookie, _ := a.Reader.Cookie(CookieName)

	decodedData, _ := base64.StdEncoding.DecodeString(cookie.Value)

	var userData UserData
	erro := json.Unmarshal(decodedData, &userData)
	if erro != nil {
		return ""
	}

	return userData.MetodoContacto
}

func (a *CookieManager) marcarPreguntaComoRespondida(preguntaField string) {
	cookie, err := a.Reader.Cookie(CookieName)
	if err != nil {
		http.Error(a.Writer, "No se pudo obtener la cookie", http.StatusBadRequest)
		return
	}

	decodedValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		http.Error(a.Writer, "Error al decodificar la cookie", http.StatusInternalServerError)
		return
	}

	var userData UserData
	err = json.Unmarshal(decodedValue, &userData)
	if err != nil {
		http.Error(a.Writer, "Error al deserializar el JSON", http.StatusInternalServerError)
		return
	}

	// Establece el campo de la pregunta a true
	userData.Preguntas[preguntaField] = true

	updatedJSON, err := json.Marshal(userData)
	if err != nil {
		http.Error(a.Writer, "Error al serializar los datos actualizados", http.StatusInternalServerError)
		return
	}

	encodedData := base64.StdEncoding.EncodeToString(updatedJSON)

	updatedCookie := &http.Cookie{
		Name:     CookieName,
		Value:    encodedData,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(a.Writer, updatedCookie)
}

func (a *CookieManager) verificarSiElCampoEsVerdadero(campo string) (bool, error) {
	cookie, err := a.Reader.Cookie(CookieName)
	if err != nil {
		return false, fmt.Errorf("no se pudo obtener la cookie: %w", err)
	}

	decodedValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return false, fmt.Errorf("error al decodificar la cookie: %w", err)
	}

	var userData UserData
	err = json.Unmarshal(decodedValue, &userData)
	if err != nil {
		return false, fmt.Errorf("error al deserializar el JSON: %w", err)
	}

	// Verifica si el campo existe en el mapa de respuestas
	valor, existe := userData.Preguntas[campo]
	if !existe {
		return false, fmt.Errorf("campo no válido: %s", campo)
	}

	return valor, nil
}

func (a *CookieManager) agregarPreguntas(userData UserData, preguntas []string) UserData {
	if userData.Preguntas == nil {
		userData.Preguntas = make(map[string]bool)
	}

	for _, pregunta := range preguntas {
		userData.Preguntas[pregunta] = false
	}

	return userData
}
