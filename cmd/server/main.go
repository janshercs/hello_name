package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

const defaultName = "default server"

var configName string

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": fmt.Sprintf("hello from: %s", configName)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	flag.StringVar(&configName, "name", "", "The name to use in the greeting")
	flag.Parse()

	if configName == "" {
		fmt.Println("name not provided, using default-server")
		configName = defaultName
	}

	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
