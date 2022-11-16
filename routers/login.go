package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cr0dr1gu3z/socialtwitt/bd"
	"github.com/cr0dr1gu3z/socialtwitt/jwt"
	"github.com/cr0dr1gu3z/socialtwitt/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "Usuario o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token correspondiente"+err.Error(), 400)
		return
	}

	res := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

	//como grabar una cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
