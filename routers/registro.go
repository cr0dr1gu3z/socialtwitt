package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cr0dr1gu3z/socialtwitt/bd"
	"github.com/cr0dr1gu3z/socialtwitt/models"
)

/*Registro es la funcion para crear en la DB el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	//el body de un httpRequest es un stream lo cual cuando se utiliza por primera ves el dato se destruye por lo cual solo se puede ocupar una sola vez
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "DE especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "YA existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error en al intentar registrar el usuario"+err.Error(), 400)
		return

	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
