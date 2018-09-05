package omg

import (
	"log"
	"net/http"
	"net/url"
)

func uploadLocalFile(w http.ResponseWriter, r *http.Request) {

}

func getFileInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	raw, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println(err.Error())
	}
	bucket := raw.Get("bucket")
	object := raw.Get("object")

	log.Println("bucket: ", bucket)
	log.Println("object: ", object)
	w.Write(fileHandler(bucket, object))
}

func updataFileInfo(w http.ResponseWriter, r *http.Request) {

}

func deleteFile(w http.ResponseWriter, r *http.Request) {

}

func getURL(w http.ResponseWriter, r *http.Request) {

}
