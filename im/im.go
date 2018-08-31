package im

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run is a start mothod
func Run() {
	// getconf()
	// beego.Run()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/im/user/:uid/sendmessage", index)
	router.HandleFunc("/api/v1/im/register", register)
	router.HandleFunc("/api/v1/im/session", session)
	router.HandleFunc("/api/v1/im/session/{sid}", sessionbyid)
	router.HandleFunc("/api/v1/im/message", message)

	log.Fatal(http.ListenAndServe(":18010", router))
}

func index(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {

}

func session(w http.ResponseWriter, r *http.Request) {

}

func sessionbyid(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sid := vars["sid"]
	io.WriteString(w, sid)

}

func message(w http.ResponseWriter, r *http.Request) {

}
