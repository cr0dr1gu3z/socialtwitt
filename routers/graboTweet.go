package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cr0dr1gu3z/socialtwitt/bd"
	"github.com/cr0dr1gu3z/socialtwitt/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InserttoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar el registro tweet"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
