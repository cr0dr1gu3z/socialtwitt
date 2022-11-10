package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN es el objeto de conecciona la db de datos*/
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://cain:Rog19uni@socialtwitt.1rro1zq.mongodb.net/?retryWrites=true&w=majority")

/* ConectarDB() es la funcion para conectar a la DB mongo */
func ConectarDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexion exitosa con la DB")
	return client

}

/*Chequeo conexion Ping DB Mongo*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1

}
