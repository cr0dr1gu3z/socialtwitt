package bd

import (
	"context"
	"time"

	"github.com/cr0dr1gu3z/socialtwitt/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*fucion es la parada final con la BD para insertar los datos del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer se ejecuta al final
	defer cancel()

	db := MongoCN.Database("socialtwitt")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
