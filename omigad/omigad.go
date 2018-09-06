package omigad

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/wst-libs/wst-sdk/conf"
	"github.com/wst-libs/wst-sdk/sdk/manager"
	"github.com/wst-libs/wst-sdk/utils"
)

var filechan chan FileInfo = make(chan FileInfo, 1024)
var config httpconfig

const (
	configurl = "http://39.105.53.16:48888/omigad-dev.yml"
)

// Run is start function
func Run() {
	// go uploadChan()
	var addr string
	err := getconfig()
	if err != nil {
		addr = ":18012"
	} else {
		addr = ":" + config.Server.Httpport
	}

	log.Fatal(http.ListenAndServe(addr, router()))
}

func getconfig() error {
	err := conf.GetConf(configurl, &config)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Printf("%v\n", config)
	return nil
}

// uploadinfo get a upload url of oss and create the record to manager.
func uploadinfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Parse url

	m, _ := url.ParseQuery(r.URL.RawQuery)
	log.Println(m)
	bucket := m.Get("bucket")
	object := m.Get("object")
	if len(bucket) == 0 {
		bucket = defaultBucket
	}
	if len(object) == 0 {
		w.Write(InvalidParams())
		return
	}
	// New object for oss
	obj, err := NewAliyunObject(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println("GetUrlFromFileHandler")
		w.Write(BucketNotFound())
		return
	}

	// check is file exist
	isExist, err := obj.IsFileExist(object)
	if err != nil {
		log.Println("GetUrlForFileHandler error: ", err.Error())
		w.Write(InternalError())
		return
	}
	if isExist != false {
		log.Println("file not exist")
		w.Write(FileAlreadyExist())
		return
	}

	// get id for manager
	res := manager.Add(config.Server.ManagerURL)
	if res.Code != 0 {
		w.Write(CreateRecordFailed())
		return
	}
	url, err := obj.PutFileWithURL(object)
	if err != nil {
		log.Println("GetUrlForFileHandler error: ", err.Error())
		w.Write(InternalError())
		return
	}

	w.Write(GetUrlForFileResponse(res.Id, url))
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

		aliyun, err := NewAliyunObject(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, f.Bucket)
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
		posturl := "http://" + config.Server.ManagerURL + "/" + f.Id

		resp, err := http.Post(posturl, "application/json", strings.NewReader(string(out)))
		if err != nil {
			log.Println("Upload file error: ", err.Error())
		} else {
			resp.Body.Close()
		}
	}
}
