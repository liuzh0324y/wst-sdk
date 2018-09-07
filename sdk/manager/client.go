package manager

import (
	"bytes"
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

	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Println("failed to marshal.")
		return ResCode{
			Code: 4005,
		}
	}
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(reqBody))
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
	if err != nil {
		log.Println("failed to recv the response. err: ", err.Error())
		return ResCode{
			Code: 4006,
		}
	}

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
