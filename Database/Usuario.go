package Database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	NOMBRE_USUARIO  = "NombreUsuario"
	LICENCIATURA    = "Licenciatura"
	METODO_CONTACTO = "MetodoContacto"
)

func (a *SingleDatabase) CrearNuevoUsuario(data any) {
	collection := a.accederColeccion("Busqueda-Tesoro", "Usuarios")
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *SingleDatabase) AniadirPuntosUsuario(nombreUsuario, licenciatura, contacto string, puntosSumar int) error {
	collection := a.accederColeccion("Busqueda-Tesoro", "Usuarios")
	filtro := bson.D{
		{Key: NOMBRE_USUARIO, Value: nombreUsuario},
		{Key: LICENCIATURA, Value: licenciatura},
		{Key: METODO_CONTACTO, Value: contacto},
	}

	actualizacion := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "puntos", Value: puntosSumar},
		}},
	}

	resultado, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)
	if err != nil {
		return fmt.Errorf("error al actualizar los puntos: %w", err)
	}

	// Verificar si el documento fue encontrado y actualizado
	if resultado.MatchedCount == 0 {
		return fmt.Errorf("no se encontró el documento para el usuario %s con los métodos de contacto proporcionados", nombreUsuario)
	}

	return nil
}
