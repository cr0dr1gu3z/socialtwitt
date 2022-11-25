package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cr0dr1gu3z/socialtwitt/bd"
	"github.com/cr0dr1gu3z/socialtwitt/models"
)

// este es un metodo porque no regresa respuesta
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	var status bool

	status, erro := bd.ModificoRegistro(t, IDUsuario)

	if erro != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro reintente nuevamente"+erro.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
