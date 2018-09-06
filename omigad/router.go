package omigad

import (
	"net/http"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/cloudstorage/file", uploadLocalFileToCloud).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/cloudstorage/file", getURLOfFileForCloud).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/cloudstorage/file", updateFileInfoOfCloud).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/cloudstorage/file", deleteFileOfCloud).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/cloudstorage/uploadinfo", getURLOfUploadFile).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/cloudstorage/callback", callback).Methods(http.MethodPost)

	return router
}
