package connection

import (
	"os"
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Function to connect to the database
func Connect() rel.Repository {
	// Connect to the database
	errorVariables := godotenv.Load()
	if errorVariables != nil {

		panic(errorVariables)

	}
	
	adapter, err := mysql.Open(os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?clientFoundRows=true&charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return rel.New(adapter)
} 