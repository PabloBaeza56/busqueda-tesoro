package Database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	lock                   = &sync.Mutex{}
	singleDatabaseInstance *SingleDatabase
)

type SingleDatabase struct {
	clienteDirecto *mongo.Client
}

func GetInstanceDatabase() *SingleDatabase {
	if singleDatabaseInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleDatabaseInstance == nil {
			fmt.Println("Creating singleDatabase instance now.")

			mongoURI := os.Getenv("MONGO_URI")
			if mongoURI == "" {
				log.Fatal("La variable de entorno MONGO_URI no está definida.")
			}
			clientOptions := options.Client().ApplyURI(mongoURI)
			clientValue, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Fatal(err)
			}

			singleDatabaseInstance = &SingleDatabase{clienteDirecto: clientValue}
		} else {
			//fmt.Println("Single singleDatabase already created.")
		}
	} else {
		//fmt.Println("Single singleDatabase already created.")
	}

	return singleDatabaseInstance
}

func (a *SingleDatabase) CerrarDb() {
	fmt.Println("Base de datos cerrada correctamente")
	a.clienteDirecto.Disconnect(context.TODO())
}

func (a *SingleDatabase) accederColeccion(database string, nombreColeccion string) *mongo.Collection {
	if a == nil {
		log.Fatal("El puntero a SingleDatabase es nulo.")
	}
	collection := a.clienteDirecto.Database(database).Collection(nombreColeccion)
	return collection
}
