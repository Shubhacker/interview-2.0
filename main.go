package main

import (
	"fmt"
	"log"
	"net/http"

	apis "github.com/Shubhacker/interview-2.0.git~/APIS"
)

func main() {
	r := apis.APIS()
	port := "8080"
	fmt.Println("Server running at :", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
