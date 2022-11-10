package middlew

import (
	"net/http"

	"github.com/cr0dr1gu3z/socialtwitt/bd"
)

/*chequeoDB es el middleware que permite conocer el estado de la DB*/
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
