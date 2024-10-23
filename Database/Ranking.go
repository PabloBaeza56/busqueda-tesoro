package Database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type RankingData struct {
	Posicion             int    `json:"posicion"`
	Username             string `json:"username"`
	Puntos               int    `json:"puntos"`
	PreguntasRespondidas int    `json:"personas_respondidas"`
}

func (a *SingleDatabase) ObtenerRanking() ([]RankingData, error) {

	collection := a.accederColeccion("Busqueda-Tesoro", "Usuarios")
	var ranking []RankingData
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error al obtener el ranking: %w", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var userData struct {
			Username  string          `bson:"username"`
			Puntos    int             `bson:"puntos"`
			Preguntas map[string]bool `bson:"preguntas"` // Map de preguntas
		}
		if err := cursor.Decode(&userData); err != nil {
			log.Fatal(err)
		}

		// Contar preguntas respondidas
		preguntasRespondidas := 0
		for _, respondido := range userData.Preguntas {
			if respondido {
				preguntasRespondidas++
			}
		}

		ranking = append(ranking, RankingData{
			Posicion:             len(ranking) + 1, // O ajusta según la lógica de tu ranking
			Username:             userData.Username,
			Puntos:               userData.Puntos,
			PreguntasRespondidas: preguntasRespondidas,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre el cursor: %w", err)
	}

	return ranking, nil
}
