package omigad

import (
	"net/http"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/cloudstorage/file", uploadLocalFile).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/cloudstorage/file", getFileInfo).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/cloudstorage/file", updataFileInfo).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/cloudstorage/file", deleteFile).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/cloudstorage/uploadinfo", getURL).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/cloudstorage/callback", callback).Methods(http.MethodPost)

	return router
}
