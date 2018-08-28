package manager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Client struct {
}

// Add file record
func (t *Client) Add(url string, body *ReqPutFile) ResCode {
	if len(url) == 0 {
		return ResCode{
			Code: 404,
		}
	}
	reqStr := &ReqPutFile{}
	reqBody, err := json.Marshal(reqStr)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, strings.NewReader(string(reqBody)))
	if err != nil {
		log.Println("failed to new request object")
		return ResCode{
			Code: 404,
		}
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "close")

	res, err := client.Do(req)

	defer func() {
		res.Body.Close()
	}()

	// stdout := os.Stdout
	// _, err = io.Copy(stdout, res.Body)

	status := res.StatusCode
	var resStr ResPutFile
	resBody, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(resBody, &resStr)
	if err != nil {
		log.Println("error: ", err.Error())
	}
	log.Println(status)
	return ResCode{
		Code: 0,
		Msg:  "success",
		Id:   resStr.Id,
	}
}

// Delete file record
func (t *Client) Del() {

}

// Update file record
func (t *Client) Update(url string, body *ReqPutFile) {
	var resp ResPutFile
	b, _ := json.Marshal(body)
	res, err := http.Post(url, "application/json", strings.NewReader(string(b)))
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}
	defer res.Body.Close()

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(out, &resp)
	if err != nil {
		return
	}
}

// Get file record
func (t *Client) Get() {

}
