package main

import (
	"net/http"

	"github.com/jaideep-penikalapati/go-mini-projects/user-web-service/controller"
)

func main() {

	controller.RegisterControllers()
	err := http.ListenAndServe(":3000", http.DefaultServeMux)
	if err != nil {
		panic("Error Starting the service")
	}

}
