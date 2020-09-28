package main

import (
	"log"
	"os"
	"strings"

	"github.com/biplabpokhrel/avoxi/handlers"
	"github.com/biplabpokhrel/avoxi/internal"
	httpRepository "github.com/biplabpokhrel/avoxi/repository/iplocation"
	dbRepository "github.com/biplabpokhrel/avoxi/repository/pgql"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	var _internal *internal.Internal

	// Get the url of the api, that give us country name based on the IP address
	IPLocationAPIUrl := os.Getenv("IP_LOCATION_URL")

	if len(strings.TrimSpace(IPLocationAPIUrl)) != 0 {
		log.Println("IP to Country search will happen in ", IPLocationAPIUrl)
		_internal = internal.New(httpRepository.LocationNetworkRepository(IPLocationAPIUrl))
	} else {
		log.Println("IP to Country mapping will be check against the database")
		// Get database connection string from Environment file
		dbURL := os.Getenv("DATABASE_URL")

		if len(strings.TrimSpace(dbURL)) != 0 {

			// Set database connection
			db := dbRepository.PGConnect(dbURL)

			// Close database connection before going out of main
			defer db.Close()

			// Added Just for sql debugging purpose
			db.AddQueryHook(dbRepository.DBLogger{})

			// Pass connection to Postgres repository of Network Map
			_internal = internal.New(dbRepository.PgqlNetworkRepository(db))

		}

	}
	routerHandler := handlers.Route{Internal: _internal}

	// IT better to use POST method than GET
	// TODO: Convert into POST
	r.GET("/network", routerHandler.GetNetwork)
	r.Run("0.0.0.0:8052")
}
