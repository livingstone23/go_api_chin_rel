package routes

import (
	"encoding/json"
	"go_chin_rel/dto"
	"go_chin_rel/utility"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func Ejemplo_get(w http.ResponseWriter, r *http.Request) {

	//Asnwer before the utility
	//w.Write([]byte("Método GET"))

	//Answer with utility
	answer := map[string]string{
		"state":         "ok",
		"message":       "Método GET",
		"Authorization": r.Header.Get("Authorization"),
	}

	utility.ResponderJson(w, http.StatusOK, answer)

}

func Ejemplo_get_con_parametro(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	answer := map[string]string{
		"state":   "ok",
		"message": "Método GET | id=" + id,
	}
	utility.ResponderJson(w, http.StatusOK, answer)

}

func Ejemplo_query_string(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	slug := r.URL.Query().Get("slug")
	answer := map[string]string{
		"estado":  "ok",
		"mensaje": "Método querystring | id=" + id + " | slug=" + slug,
	}
	utility.ResponderJson(w, http.StatusOK, answer)
	//w.Write([]byte("Método querystring | id=" + id + " | slug=" + slug))
}

func Ejemplo_post(w http.ResponseWriter, r *http.Request) {
	data := dto.TeamDto{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		answer := map[string]string{
			"state":   "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusOK, answer)
		return
	}
	answer := map[string]string{
		"state":   "ok",
		"message": "Método POST",
		"name":    data.Name,
	}
	utility.ResponderJson(w, http.StatusOK, answer)
}

/*
	func Ejemplo_post(w http.ResponseWriter, r *http.Request) {
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Método POST",
		}
		utilidades.ResponderJson(w, http.StatusOK, respuesta)
	}
*/
func Ejemplo_put(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	answer := map[string]string{
		"estado":  "ok",
		"mensaje": "Método PUT | id=" + id,
	}
	utility.ResponderJson(w, http.StatusOK, answer)
	//w.Write([]byte("Método PUT | id=" + id))
}
func Ejemplo_delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	answer := map[string]string{
		"estado":  "ok",
		"mensaje": "Método DELETE | id=" + id,
	}
	utility.ResponderJson(w, http.StatusOK, answer)
	//w.Write([]byte("Método DELETE | id=" + id))
}

func Ejemplo_upload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("foto")
	if err != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusOK, answer)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	//fmt.Println(time[4])
	foto := string(time[4][6:14]) + "." + extension
	var archivo string = "public/uploads/pictures/" + foto
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusOK, answer)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusOK, answer)
		return
	}
	//acá insertamos el registro de la BD

	//retornamos
	answer := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro exitosamente",
		"foto":    foto,
	}
	utility.ResponderJson(w, http.StatusOK, answer)
}
