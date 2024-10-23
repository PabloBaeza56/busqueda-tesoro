package Database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	NOMBRE_USUARIO    = "username"
	METODO_CONTACTO_1 = "contact_method_1"
	METODO_CONTACTO_2 = "contact_method_2"
)

func (a *SingleDatabase) CrearNuevoUsuario(data any) {
	collection := a.accederColeccion("Busqueda-Tesoro", "Usuarios")
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *SingleDatabase) AniadirPuntosUsuario(username, contactMethod1, contactMethod2 string, puntosSumar int) error {
	collection := a.accederColeccion("Busqueda-Tesoro", "Usuarios")
	filtro := bson.D{
		{Key: NOMBRE_USUARIO, Value: username},
		{Key: METODO_CONTACTO_1, Value: contactMethod1},
		{Key: METODO_CONTACTO_2, Value: contactMethod2},
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
		return fmt.Errorf("no se encontró el documento para el usuario %s con los métodos de contacto proporcionados", username)
	}

	return nil
}
