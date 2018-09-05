package omg

import (
	"net/http"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/privatestorage/file", uploadLocalFile).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/privatestorage/file", getFileInfo).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/privatestorage/file", updataFileInfo).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/privatestorage/file", deleteFile).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/privatestorage/uploadinfo", getURL).Methods(http.MethodGet)

	return router
}
