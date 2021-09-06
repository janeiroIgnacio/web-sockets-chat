package app

import (
	"LefkasChat/handlers"
	"net/http"
	"os"
)

func Start(){

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handlers.GetInstance().SetupRoutes()
	http.ListenAndServe(":"+port, nil)

}

