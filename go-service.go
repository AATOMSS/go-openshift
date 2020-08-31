package main

import (
	"fmt"
	"net/http"
	"os"
)

func cobrosHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Inicio proceso..."
	}

	fmt.Fprintln(w, response)
	fmt.Println("Llamado al sitio...")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", cobrosHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9090"
	}
	go listenAndServe(port)

	port = os.Getenv("SECOND_PORT")
	if len(port) == 0 {
		port = "9999"
	}
	go listenAndServe(port)

	select {}
}