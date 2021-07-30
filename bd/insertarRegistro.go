package bd

import (
	"context"
	"time"

	"github.com/sebaz143/socialMedia/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarRegistro Insertar registro a la base de datos*/
func InsertarRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //en la caja de ejecuciones de context se cancela la operacion de timeout, para que no quede la operacion abierta

	db := MongoCon.Database("socialMediaGO")
	coleccion := db.Collection("user")
	u.Password, _ = EncriptarPassword(u.Password)

	result, err := coleccion.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	//Esto no se usa en el desarrollo pero es interesante conocerlo:
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
