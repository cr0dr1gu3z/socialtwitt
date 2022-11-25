package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cr0dr1gu3z/socialtwitt/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {

	var ID string = r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe de enviar el parametro id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
