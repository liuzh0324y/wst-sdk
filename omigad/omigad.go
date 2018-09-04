package omigad

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/mux"
	"github.com/wst-libs/wst-sdk/conf"
	"github.com/wst-libs/wst-sdk/errors"
	"github.com/wst-libs/wst-sdk/sdk/manager"
	"github.com/wst-libs/wst-sdk/utils"
)

var filechan chan FileInfo = make(chan FileInfo, 1024)
var config OMGConfig

const (
	configurl = "http://39.105.53.16:48888/im-dev.yml"
)

// Run is start function
func Run() {
	go uploadChan()
	var addr string
	err := getconfig()
	if err != nil {
		addr = ":18010"
	} else {
		addr = ":" + config.Server.Httpport
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/cloudstorage/file", file)
	router.HandleFunc("/api/v1/cloudstorage/callback", callback)
	router.HandleFunc("/api/v1/cloudstorage/uploadinfo", uploadinfo)

	log.Fatal(http.ListenAndServe(addr, router))
}

func getconfig() error {
	err := conf.GetConf(configurl, &config)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func file(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setHeader(w.Header())
	var outbody []byte

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		object := vars["object"]
		outbody = getfileHandler(bucket, object, body)
	} else if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		object := vars["object"]
		outbody = postfileHandler(bucket, object, body)
	} else if r.Method == http.MethodPut {
		outbody = createHandler(body)
	} else if r.Method == http.MethodDelete {
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		object := vars["object"]
		outbody = delfileHandler(bucket, object, body)
	} else {
		outbody = errors.NotSupportMethod()
	}

	w.Write(outbody)

}

func callback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setHeader(w.Header())
	var outbody []byte

	if r.Method == "POST" {

	} else {

	}
	w.Write(outbody)

}

func uploadinfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setHeader(w.Header())
	var outbody []byte

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var req ReqBody
	if json.Unmarshal(body, req) != nil {
		log.Println("Failed to json unmarshal.")
		// return ResponseFailed()
	}

	manager.Update(beego.AppConfig.String("managerurl")+"/"+req.Id, 3)

	// return ResponseSuccess()
	w.Write(outbody)
}

func setHeader(h http.Header) {
	h.Add("Content-Type", "application/json")
	h.Add("Connection", "close")
}

func uploadChan() {
	for {
		f := <-filechan
		log.Println("upload file: ", f.FileName)

		filepath := f.FilePath + "/" + f.FileName
		fileinfo, err := os.Stat(filepath)
		if err != nil {
			logs.Error("PutFile error: ", err.Error())
			continue
			// return FileNotExist(request.SeqNum)
		}
		if fileinfo.IsDir() != false {
			logs.Error("PutFile error: This is a dir.")
			continue
		}

		aliyun, err := NewAliyunObject(beego.AppConfig.String("endpoint"), beego.AppConfig.String("accesskey"), beego.AppConfig.String("secretkey"), f.Bucket)
		if err != nil {
			logs.Error(err.Error())
			continue
		}

		err = aliyun.PutFile(f.Object, filepath)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		isDel, _ := beego.AppConfig.Bool("remove")
		if isDel == true {
			err = os.Remove(filepath)

			if err != nil {
				logs.Error(err.Error())
				continue
			}
		}

		res := ResUpdateFileToP{
			utils.RequestCommon{
				Version: utils.Version,
				SeqNum:  1,
				From:    "omigad",
				To:      "file manager",
				Type:    "omigad",
				Number:  "XXXX-XXXX-XXXX-XXXX",
				Uid:     f.Id,
			},
			PostInfo{
				Name:     f.FileName,
				Type:     f.FileType,
				Url:      "https://sample.com/sample.mp4",
				Key:      "",
				Secret:   "",
				Bucket:   f.Bucket,
				Object:   f.Object,
				Region:   "",
				Endpoint: beego.AppConfig.String("endpoint"),
				Desc:     "",
			},
		}
		out, _ := json.Marshal(res)
		posturl := "http://" + beego.AppConfig.String("puthost") + ":" + beego.AppConfig.String("putport") + beego.AppConfig.String("putpath") + "/" + f.Id

		resp, err := http.Post(posturl, "application/json", strings.NewReader(string(out)))
		if err != nil {
			log.Println("Upload file error: ", err.Error())
		} else {
			resp.Body.Close()
		}
	}
}
