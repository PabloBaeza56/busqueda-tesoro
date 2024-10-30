package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	controlador "tresure-hunt/Controller"
	database "tresure-hunt/Database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	serverAddr   = ":8080"
	readTimeout  = 2 * time.Second
	writeTimeout = 2 * time.Second
	idleTimeout  = 5 * time.Second
)

func init() {
	godotenv.Load()
}

func main() {

	router := mux.NewRouter()
	servirRecursos(router)

	miManejador := database.GetInstanceDatabase()
	defer miManejador.CerrarDb()

	activarVigilanteDeApagado(miManejador)

	controlador.ArranqueInicial(router)

	server := configurarServidor(router)

	log.Println("Servidor corriendo ... , " + serverAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}

func activarVigilanteDeApagado(manejador *database.SingleDatabase) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		manejador.CerrarDb()
		log.Fatalf("Programa Finalizado: %v", sig)
	}()
}

func configurarServidor(router *mux.Router) *http.Server {
	return &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}

func servirRecursos(r *mux.Router) {

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./Vista/css/"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./Vista/js/"))))

}
