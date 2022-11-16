package middlew

import(
	"net/http"
	"github.com/cr0dr1gu3z/socialtwitt/routers"
)


func ValidoJWT(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		//obtiene el valor del token del header Autorization y lo manda al metodo ProcespToken
		_,_,_, err := routers.ProcesoToken(r.Header.Get("Authorization"))

		if err != nil{
			http.Error(w, "Error en el token !"+ err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w,r)
	
	}
}