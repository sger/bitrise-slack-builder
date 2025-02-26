package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bitrise-Slack Builder is running!")
	})

	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
