package omg

import "github.com/gorilla/mux"

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/privatestorage/file", putfile)
	router.HandleFunc("/api/v1/privatestorage/file/{fid}", getfile)
	return router
}
