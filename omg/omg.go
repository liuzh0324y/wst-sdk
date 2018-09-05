package omg

import (
	"log"
	"net/http"
)

// Run is start function
func Run() {

	log.Fatal(http.ListenAndServe(":18013", router()))
}
