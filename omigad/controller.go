package omigad

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/wst-libs/wst-sdk/sdk/manager"
)

type Controller struct {
	beego.Controller
}

func SetOutPutHeaders(ctx *context.Context) {
	ctx.Output.Header("Connection", "close")
	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Header("Server", "omigad:V1.0")
}
func (t *Controller) PutFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(PutFileHandler(t.Ctx))
}

func (t *Controller) GetFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(GetFileHandler(t.Ctx))
}

func (t *Controller) UpdateFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(UpdateFileHandler(t.Ctx))
}

func (t *Controller) DeleteFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(DeleteFileHandler(t.Ctx))
}

func (t *Controller) CallBack() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(CallBackHandler(t.Ctx))
}

func (t *Controller) GetUrlForFile() {
	SetOutPutHeaders(t.Ctx)
	t.Ctx.Output.Body(GetUrlForFileHandler(t.Ctx))
}

func uploadLocalFile(w http.ResponseWriter, r *http.Request) {

	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {

	}

	bucket := u.Get("bucket")
	object := u.Get("object")
	log.Println("bucket: ", bucket)
	log.Println("object: ", object)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	outbody := getfileHandler(bucket, object, body)

	w.Write(outbody)
}

func getFileInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}

func updataFileInfo(w http.ResponseWriter, r *http.Request) {
	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {

	}

	bucket := u.Get("bucket")
	object := u.Get("object")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	}
	outbody := postfileHandler(bucket, object, body)
	w.Write(outbody)
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
	}
	bucket := u.Get("bucket")
	object := u.Get("object")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	}
	outbody := delfileHandler(bucket, object, body)
	w.Write(outbody)
}

func getURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Parse url

	m, _ := url.ParseQuery(r.URL.RawQuery)
	log.Println(m)
	bucket := m.Get("bucket")
	object := m.Get("object")
	if len(bucket) == 0 {
		bucket = beego.AppConfig.String("bucket")
	}
	if len(object) == 0 {
		w.Write(InvalidParams())
		return
	}
	// New object for oss
	obj, err := NewAliyunObject(beego.AppConfig.String("endpoint"), beego.AppConfig.String("accesskey"), beego.AppConfig.String("secretkey"), bucket)
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
	res := manager.Add(beego.AppConfig.String("managerurl"))
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

func callback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setHeader(w.Header())
	var outbody []byte

	w.Write(outbody)

}
