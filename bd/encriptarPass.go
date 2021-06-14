package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword encripta la contrasena de usuario*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //es la cantidad 2^costo de veces que se va a llamar de forma recursiva la encriptacion (6 para un usuario normal, 8 para un super usuario)
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
