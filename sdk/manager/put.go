package manager

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Put(url string) {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPut, url, strings.NewReader("asdflkjkljasdklfjskladfjklasdfjklsafjklasdfj"))
	if err != nil {
		log.Println(err.Error())
	}

	res, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}

	out, _ := ioutil.ReadAll(res.Body)
	log.Println(string(out))
}
