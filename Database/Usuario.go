package Database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	NOMBRE_USUARIO  = "NombreUsuario"
	LICENCIATURA    = "Licenciatura"
	METODO_CONTACTO = "MetodoContacto"
)

func (a *SingleDatabase) CrearNuevoUsuario(data any) error {
	collection := a.accederColeccion("Busqueda-Tesoro", "Usuarios")

	// Convertir cualquier struct (como UserData) a bson.M
	raw, err := bson.Marshal(data)
	if err != nil {
		return fmt.Errorf("error al serializar datos del usuario: %w", err)
	}

	var doc bson.M
	err = bson.Unmarshal(raw, &doc)
	if err != nil {
		return fmt.Errorf("error al convertir a bson.M: %w", err)
	}

	filtro := bson.D{
		{Key: NOMBRE_USUARIO, Value: doc[NOMBRE_USUARIO]},
		{Key: LICENCIATURA, Value: doc[LICENCIATURA]},
		{Key: METODO_CONTACTO, Value: doc[METODO_CONTACTO]},
	}

	var usuarioExistente bson.M
	err = collection.FindOne(context.TODO(), filtro).Decode(&usuarioExistente)
	if err == nil {
		return fmt.Errorf("el usuario ya existe")
	}

	_, err = collection.InsertOne(context.TODO(), data)
	if err != nil {
		return fmt.Errorf("error al insertar usuario: %w", err)
	}

	return nil
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
