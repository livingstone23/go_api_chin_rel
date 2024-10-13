package routes

import (
	
	"context"
	"go_chin_rel/connection"
	"go_chin_rel/models"
	"go_chin_rel/utility"
	"net/http"
	"io"
	"os"
	"strings"
	"strconv"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	
)

// Function to get all pictures
func Players_pictures_get(w http.ResponseWriter, r *http.Request) {
	var datos []models.PlayerPicture

	if err := connection.Connect().FindAll(context.TODO(), &datos, sort.Desc("id")); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}

	utility.ResponderJson(w, http.StatusOK, datos)
}

// Function to get all pictures by player id
func Players_pictures_get_by_player_id(w http.ResponseWriter, r *http.Request) {
	var datos []models.PlayerPicture
	player_id := chi.URLParam(r, "player_id")

	if err := connection.Connect().FindAll(context.TODO(), &datos, where.Eq("player_id", player_id)); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
			"player_id": player_id,
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}

	utility.ResponderJson(w, http.StatusOK, datos)
}

// Function to upload a picture for a player
func Players_pictures_post(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado2",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	//fmt.Println(time[4])
	foto := string(time[4][6:14]) + "." + extension
	var archivo string = "public/uploads/players/" + foto
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	//acá insertamos el registro de la BD
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	datos := models.PlayerPicture{Picture: foto, PlayerID: id}
	if errInsert := connection.Connect().Insert(context.TODO(), &datos); errInsert != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	//retornamos
	respuesta := map[string]string{
		"state":  "ok",
		"message": "Se creó el registro exitosamente",
	}
	utility.ResponderJson(w, http.StatusCreated, respuesta)
}

// Function to delete a picture
func Players_pictures_delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var data models.PlayerPicture
	if errBuscar := connection.Connect().Find(context.TODO(), &data, where.Eq("id", id)); errBuscar != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	e := os.Remove("public/uploads/players/" + data.Picture)
	if e != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	if err := connection.Connect().Delete(context.TODO(), &data); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se eliminó el registro exitosamente",
	}
	utility.ResponderJson(w, http.StatusOK, respuesta)
}