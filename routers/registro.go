package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sebaz143/socialMedia/bd"
	"github.com/sebaz143/socialMedia/models"
)

/*Registro es la funcion para crear en la BD el resitro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t) //El body de un http es un stream, lo que quiere decir que una vez el objeto se lea se destruye

	if err != nil {
		http.Error(w, "Error en los datos recividos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password debe tener mas de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {

		http.Error(w, "Ya existe un usuario registrado con ese Email", 400)
		return
	}

	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al registral el usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
