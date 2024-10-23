package Database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	NOMBRE_PREGUNTA          = "nombre"
	PERSONA_QUE_LA_RESPONDIO = "personas_respondidas"
)

type Pregunta struct {
	Nombre              string `bson:"nombre"`
	PersonasRespondidas int    `bson:"personas_respondidas"`
}

func (a *SingleDatabase) AniadirPreguntasQueNoEstenEnLaDB(nombrePregunta string) error {
	collection := a.accederColeccion("Busqueda-Tesoro", "Preguntas")

	// Filtro para verificar si la pregunta ya existe
	filtro := bson.D{{Key: NOMBRE_PREGUNTA, Value: nombrePregunta}}

	var preguntaExistente Pregunta
	err := collection.FindOne(context.TODO(), filtro).Decode(&preguntaExistente)
	if err == nil {
		// Si no hay error, significa que la pregunta ya existe
		return fmt.Errorf("la pregunta '%s' ya existe en la base de datos", nombrePregunta)
	}
	if err != mongo.ErrNoDocuments {
		// Si ocurrió otro error al buscar, retornar ese error
		return fmt.Errorf("error al verificar la pregunta en la base de datos: %w", err)
	}

	// Si no existe, procedemos a insertarla
	pregunta := Pregunta{
		Nombre:              nombrePregunta,
		PersonasRespondidas: 0,
	}

	_, err = collection.InsertOne(context.TODO(), pregunta)
	if err != nil {
		return fmt.Errorf("error al insertar la pregunta en la base de datos: %w", err)
	}

	return nil
}

func (a *SingleDatabase) NumeroPersonasQueRespondieron(nombrePregunta string) (int, error) {
	// Filtro para buscar la pregunta por su nombre
	collection := a.accederColeccion("Busqueda-Tesoro", "Preguntas")
	filtro := bson.D{primitive.E{Key: NOMBRE_PREGUNTA, Value: nombrePregunta}}

	// Estructura para almacenar el resultado
	var pregunta Pregunta

	// Buscar el documento
	err := collection.FindOne(context.TODO(), filtro).Decode(&pregunta)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, fmt.Errorf("no se encontró la pregunta con el nombre: %s", nombrePregunta)
		}
		return 0, fmt.Errorf("error al buscar la pregunta: %w", err)
	}

	// Retornar el número de personas que han respondido
	return pregunta.PersonasRespondidas, nil
}

func (a *SingleDatabase) IncrementarPersonasQueRespondieron(nombrePregunta string) error {
	// Filtro para buscar la pregunta por su nombre
	collection := a.accederColeccion("Busqueda-Tesoro", "Preguntas")
	filtro := bson.D{primitive.E{Key: NOMBRE_PREGUNTA, Value: nombrePregunta}} // Usando nombres de campos

	// Actualización para incrementar el campo "personas_respondidas" en 1
	actualizacion := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: PERSONA_QUE_LA_RESPONDIO, Value: 1}, // Usando nombres de campos
		}},
	}

	// Actualizar el documento
	_, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no se encontró la pregunta con el nombre: %s", nombrePregunta)
		}
		return fmt.Errorf("error al actualizar la pregunta: %w", err)
	}

	return nil
}

func (a *SingleDatabase) MarcarPreguntaComoRespondida(username, pregunta string) error {
	collection := a.accederColeccion("Busqueda-Tesoro", "Usuarios")

	// Filtro para buscar al usuario por su nombre de usuario
	filtro := bson.D{
		{Key: NOMBRE_USUARIO, Value: username},
	}

	// Actualización para establecer la pregunta como true
	actualizacion := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: fmt.Sprintf("preguntas.%s", pregunta), Value: true}, // Utiliza el nombre de la pregunta para el mapa
		}},
	}

	// Actualizar el documento
	_, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)
	if err != nil {
		return fmt.Errorf("error al actualizar la pregunta: %w", err)
	}

	return nil
}
