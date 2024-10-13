package routes


import (
	"context"
	"go_chin_rel/connection"
	"go_chin_rel/models"
	"encoding/json"
	"go_chin_rel/dto"
	"go_chin_rel/utility"
	"net/http"
	"github.com/go-rel/rel/where"
	"github.com/go-chi/chi/v5"
	"github.com/go-rel/rel/sort"

)


// Function to get all players
func Players_get(w http.ResponseWriter, r *http.Request) {
	var datos []models.Player

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

// Function to get a player by id
func Players_get_by_id(w http.ResponseWriter, r *http.Request) {
	
	var player models.Player
	if err := connection.Connect().Find(context.TODO(), &player, where.Eq("id", chi.URLParam(r, "id"))); err != nil {
		answer := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, answer)
		return
	}
	utility.ResponderJson(w, http.StatusOK, player)
}

// Function to create a player
func Players_post(w http.ResponseWriter, r *http.Request) {
	
	data := dto.PlayerDto{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	newPlayer := models.Player{
		Name: data.Name,
		Description: data.Description,
		TeamID: data.TeamId,
	}
	if err := connection.Connect().Insert(context.TODO(), &newPlayer); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}

	answer := map[string]string{
		"state":  "ok",
		"message": "Jugador creado",
	}
	utility.ResponderJson(w, http.StatusOK, answer)
}

// Function to update a player
func Players_put(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := dto.PlayerDto{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}

	//Validate if the player exists
	var player models.Player
	if err := connection.Connect().Find(context.TODO(), &player, where.Eq("id", id)); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Jugador no encontrado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}


	playerUpdate := models.Player{
		Id: player.Id,
		Name: data.Name,
		Description: data.Description,
		TeamID: data.TeamId,
	}
	if err := connection.Connect().Update(context.TODO(), &playerUpdate ); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}

	answer := map[string]string{
		"state":  "ok",
		"message": "Jugador actualizado",
	}
	utility.ResponderJson(w, http.StatusOK, answer)
}

// Function to delete a player
func Players_delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	//Validate if the player exists
	var player models.Player
	if err := connection.Connect().Find(context.TODO(), &player, where.Eq("id", id)); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Jugador no encontrado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}

	if err := connection.Connect().Delete(context.TODO(), &player); err != nil {
		respuesta := map[string]string{
			"state":  "error",
			"message": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}

	answer := map[string]string{
		"state":  "ok",
		"message": "Jugador eliminado",
	}
	utility.ResponderJson(w, http.StatusOK, answer)
}
