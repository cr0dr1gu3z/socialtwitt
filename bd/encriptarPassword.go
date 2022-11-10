package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(pass string) (string, error) {

	//cantidad de pasadas por el texto para encriptarla mientra mas cantidad mayor cantidad de pasadas tiene para encriptadas
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
