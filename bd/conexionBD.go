package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCon es el objeto de conecion a la BD, esta variable es global por estar en mayuscula*/
var MongoCon = ConectarBD()

/*clientOptions es la URL a la BD en atlas*/
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@sebasmongodb.tp0w7.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD Conectar a a base de datos, restorna una variable de tipo mongo.Client*/
func ConectarBD() *mongo.Client {

	//context.TODO() es el contexto por defecto y no hace nada, un contexto es un entorno de
	//ejecucucion, tambien pueden ponerse contextos de TIMEOUT por ejemplo. Nos sirven para
	//comunicar informacion entre ejecuciones.

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion existosa a la DB")
	return client
}

/*CheckConnection() chequea la conexion*/
func CheckConnection() int {
	err := MongoCon.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return 0
	} else {
		return 1
	}
}
