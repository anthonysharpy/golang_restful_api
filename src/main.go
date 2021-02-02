/* Copyright Anthony Sharp 2021 */

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting JDI RESTful API server...\n")

	// Initialise MySQL
	fmt.Printf("Initialising MySQL database... \n")
	if(!InitialiseDatabase()) {
		log.Fatal(errors.New("ERROR INITIALISING DATABASE. STOPPING.\n"))
	}
	defer CloseDatabase()
	fmt.Printf("Done!\n")

	// Listen to port 8080 for requests
	fmt.Printf("Server now starting.\n")
	http.HandleFunc("/", RequestHandlerFunction)

	var err = http.ListenAndServe(":8080", nil)

	if (err != nil) {
		fmt.Printf("Fatal error: " + err.Error())
	}

	// TODO: implement way to shut down server without force-closing it.
}