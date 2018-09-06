package omigad

import (
	"log"

	"github.com/wst-libs/wst-sdk/sdk/manager"
)

func uploadLocalFileToCloudHandler(body []byte) []byte {
	s, err := PutFileRequest(body)
	if err != nil {
		return JsonFormatErr()
	}
	info := FileInfo{
		Id:       s.Data.Id,
		FilePath: s.Data.Path,
		FileName: s.Data.Name,
		FileType: s.Data.Type,
		Bucket:   s.Data.Bucket,
		Object:   s.Data.Object,
	}
	filechan <- info

	return PutFileResponse()
}

func getURLOfFileForCloudHandler(bucket, object string) []byte {

	obj, err := NewAliyunObject(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println("GetFileHandler error: ", err.Error())
		return BucketNotFound()
	}

	isExist, err := obj.IsFileExist(object)
	if err != nil {
		log.Println("GetFileHandler Error: ", err.Error())
		return InternalError()
	}
	if isExist != true {
		log.Println("file not exist.")
		return FileNotFound()
	}
	returl, err := obj.GetFileWithURL(object, 3600)
	if err != nil {
		log.Println("GetFileHandler error: ", err.Error())
		return InternalError()
	}
	log.Println("url: ", returl)
	return GetFileResponse(returl)
}

func updateFileInfoOfCloudHandler(bucket, object string, body []byte) []byte {

	s, err := UpdateFileRequest(body)
	if err != nil {
		return JsonFormatErr()
	}

	obj, err := NewAliyunObject(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return BucketNotFound()
	}

	isExist, err := obj.IsFileExist(object)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return InternalError()
	}
	if isExist != true {
		log.Println("file not exist.")
		return FileNotFound()
	}

	obj.UpdateFile(object, "description", s.Data.Desc)

	return UpdateFileResponse()
}

func deleteFileOfCloudHandler(bucket, object string, body []byte) []byte {
	obj, err := NewAliyunObject(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return BucketNotFound()
	}

	isExist, err := obj.IsFileExist(object)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return InternalError()
	}
	if isExist != true {
		log.Println("file not exist.")
		return FileNotFound()
	}

	obj.DeleteFile(object)

	return DeleteFileResponse()
}

func getURLOfUploadFileHandler(bucket, object string) []byte {
	// New object for oss
	log.Println("endpoint: ", config.Server.Endpoint)
	log.Println("accesskey: ", config.Server.AccessKey)
	log.Println("secretkey: ", config.Server.SecretKey)
	obj, err := NewAliyunObject(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println("GetUrlFromFileHandler")
		return BucketNotFound()
	}

	// check is file exist
	isExist, err := obj.IsFileExist(object)
	if err != nil {
		log.Println("GetUrlForFileHandler error: ", err.Error())
		return InternalError()
	}
	if isExist != false {
		log.Println("file not exist")
		return FileAlreadyExist()
	}

	url, err := obj.PutFileWithURL(object)
	if err != nil {
		log.Println("GetUrlForFileHandler error: ", err.Error())
		return InternalError()
	}

	// get id for manager
	log.Println("manager url: ", config.Server.ManagerURL)
	res := manager.Add(config.Server.ManagerURL)
	if res.Code != 0 {
		return CreateRecordFailed()
	}

	return GetUrlForFileResponse(res.Id, url)
}
