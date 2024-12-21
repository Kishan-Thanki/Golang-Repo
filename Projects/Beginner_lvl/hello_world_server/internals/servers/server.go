package servers

import (
	"fmt"
	"net/http"

	"github.com/Kishan-Thanki/Go_Projects/hello-world-server/internals/config"
	"github.com/Kishan-Thanki/Go_Projects/hello-world-server/internals/handlers"
)

func setupRoutes() {
	http.HandleFunc("/", handlers.HelloWorldHandler)
}

func Start() {
	port := config.GetPort()

	setupRoutes()

	fmt.Printf("Server running at http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
