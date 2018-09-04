package omg

import (
	"log"
	"net/http"
)

// Run is start function
func Run() {

	log.Fatal(http.ListenAndServe("", router()))
}
