package conf

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-yaml/yaml"
)

// GetConf is a mothod
func GetConf(url string, out interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer res.Body.Close()

	reqBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(reqBody, out)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
