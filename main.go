package main


import (
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"
	"go_chin_rel/routes"
	
)



func main() {
	fmt.Println("Hello, from Api with CHIN and REL!!!")

	// Definir prefijo de la API
	var prefij string = "/api/v1"

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		//For static files
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/public/*", http.StripPrefix("/public", fs))

	// Example routes
	router.Get(prefij+"/example", routes.Ejemplo_get)
	router.Get(prefij+"/example/{id}", routes.Ejemplo_get_con_parametro)
	router.Post(prefij+"/example", routes.Ejemplo_post)
	router.Put(prefij+"/example", routes.Ejemplo_put)
	router.Delete(prefij+"/example", routes.Ejemplo_delete)
	router.Post(prefij+"/example_query_string", routes.Ejemplo_query_string)


	// Team routes
	router.Get(prefij+"/teams", routes.Teams_get)
	router.Get(prefij+"/teams/{id}", routes.Teams_get_by_id)
	router.Post(prefij+"/teams", routes.Teams_post)
	router.Put(prefij+"/teams/{id}", routes.Teams_put)
	router.Delete(prefij+"/teams/{id}", routes.Teams_delete)

	// Player routes
	router.Get(prefij+"/players", routes.Players_get)
	router.Get(prefij+"/players/{id}", routes.Players_get_by_id)
	router.Post(prefij+"/players", routes.Players_post)
	router.Put(prefij+"/players/{id}", routes.Players_put)
	router.Delete(prefij+"/players/{id}", routes.Players_delete)

	// Player picture routes
	router.Get(prefij+"/players_pictures", routes.Players_pictures_get)
	router.Get(prefij+"/players_pictures/{player_id}", routes.Players_pictures_get_by_player_id)
	router.Post(prefij+"/players_pictures/{id}", routes.Players_pictures_post)
	router.Delete(prefij+"/players_pictures/{id}", routes.Players_pictures_delete)



	// Cargar variables de entorno
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }


	http.ListenAndServe(":"+os.Getenv("API_PORT"), router)

}
