package bd

import (
	"context"
	"time"

	"github.com/sebaz143/socialMedia/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email y chequea en la base de datos si existe o no*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCon.Database("socialMediaGO")
	coleccion := db.Collection("user")

	condicion := bson.M{"email": email} //formatea el email de entrada a la funcion en formato bson , asignando a la clave email el valo email

	var resultado models.Usuario //aca voy a grabar el resultado del find

	err := coleccion.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
