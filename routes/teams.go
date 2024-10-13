package routes

import (
	"context"
	"encoding/json"
	"go_chin_rel/connection"
	"go_chin_rel/dto"
	"go_chin_rel/models"
	"go_chin_rel/utility"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	"github.com/gosimple/slug"
)

// Function to get all teams
func Teams_get(w http.ResponseWriter, r *http.Request) {
	var datos []models.Team

	//if err := conectar.Connect().FindAll(context.TODO(), &datos); err != nil {
	if err := connection.Connect().FindAll(context.TODO(), &datos, sort.Desc("id")); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}
	utility.ResponderJson(w, http.StatusOK, datos)
}

// Function to get a team by id
func Teams_get_by_id(w http.ResponseWriter, r *http.Request) {

	var team models.Team
	if err := connection.Connect().Find(context.TODO(), &team, where.Eq("id", chi.URLParam(r, "id"))); err != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, answer)
		return
	}
	utility.ResponderJson(w, http.StatusOK, team)
}

// Function to create a team
func Teams_post(w http.ResponseWriter, r *http.Request) {
	
	data := dto.TeamDto{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, respuesta)
		return
	}

	// Create a new team
	newData := models.Team{Name: data.Name, Slug: slug.Make(data.Name)}

	if errorInsert := connection.Connect().Insert(context.TODO(), &newData); errorInsert != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, answer)
		return
	}

	answer := map[string]string{
		"estado":  "ok",
		"mensaje": "Equipo creado correctamente",
	}

	utility.ResponderJson(w, http.StatusOK, answer)

}

// Function to update a team
func Teams_put(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := dto.TeamDto{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, answer)
		return
	}

	//Valid if the team exists
	var existingTeam models.Team
	if errLooking := connection.Connect().Find(context.TODO(), &existingTeam, where.Eq("id", id)); errLooking != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "El equipo no existe",
		}
		utility.ResponderJson(w, http.StatusBadRequest, answer)
		return
	}

	// Create a new team
	dataToUpdate := models.Team{Id: existingTeam.Id,Name: data.Name, Slug: slug.Make(data.Name)}

	if errorInsert := connection.Connect().Update(context.TODO(), &dataToUpdate); errorInsert != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, answer)
		return
	}

	answer := map[string]string{
		"estado":  "ok",
		"mensaje": "Equipo actualizado correctamente",
	}

	utility.ResponderJson(w, http.StatusOK, answer)

}

// Function to delete a team
func Teams_delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	//Valid if the team exists
	var existingTeam models.Team
	if errLooking := connection.Connect().Find(context.TODO(), &existingTeam, where.Eq("id", id)); errLooking != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "El equipo no existe",
		}
		utility.ResponderJson(w, http.StatusBadRequest, answer)
		return
	}

	if errorDelete := connection.Connect().Delete(context.TODO(), &existingTeam); errorDelete != nil {
		answer := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utility.ResponderJson(w, http.StatusBadRequest, answer)
		return
	}

	answer := map[string]string{
		"estado":  "ok",
		"mensaje": "Equipo eliminado correctamente",
	}

	utility.ResponderJson(w, http.StatusOK, answer)

}