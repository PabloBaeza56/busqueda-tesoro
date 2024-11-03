package Database

import (
	"context"
	"fmt"
	"log"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
)

type RankingData struct {
	Posicion             int    `json:"posicion"`
	NombreUsuario        string `json:"NombreUsuario"`
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
			NombreUsuario string          `bson:"NombreUsuario"`
			Puntos        int             `bson:"puntos"`
			Preguntas     map[string]bool `bson:"preguntas"` // Map de preguntas
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

		// Ingresar al usuario con un ranking temporal
		ranking = append(ranking, RankingData{
			Posicion:             0,
			NombreUsuario:        userData.NombreUsuario,
			Puntos:               userData.Puntos,
			PreguntasRespondidas: preguntasRespondidas,
		})
	}

	// Ordenar al ranking (por puntos) de mayor a menor
	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i].Puntos > ranking[j].Puntos
	})

	// Truncar a ranking a un máximo de 10 usuarios
	if len(ranking) > 10 {
		ranking = ranking[0:9]
	}

	//Asignar la posición de los usuarios de acuerdo a su orden en el arreglo
	for index := range ranking {
		ranking[index].Posicion = index + 1
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre el cursor: %w", err)
	}

	return ranking, nil
}
