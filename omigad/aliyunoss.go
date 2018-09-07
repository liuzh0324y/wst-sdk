package omigad

import (
	"io"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego/logs"
)

// AliyunOss is a type
type AliyunOss struct {
	client *oss.Client
	bucket *oss.Bucket
}

// Create OSS instance
func NewAliyunObject(endpoint_, accessKey_, secretKey_, name string) (*AliyunOss, error) {
	// logs.Info("accessKey: %s, secretKey: %s, bucket: %s, endpoint: %s", accessKey_, secretKey_, name, endpoint_)

	c, err := oss.New(endpoint_, accessKey_, secretKey_)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	bucket, err := c.Bucket(name)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	return &AliyunOss{
		client: c,
		bucket: bucket,
	}, nil
}

// Check if the file exist
func (this AliyunOss) IsFileExist(name string) (bool, error) {
	isExist, err := this.bucket.IsObjectExist(name)
	if err != nil {
		logs.Error(err.Error())
		return isExist, err
	}
	return isExist, nil
}

// Upload local file to the OSS server
func (t AliyunOss) PutFile(name, file string) error {
	options := []oss.Option{
		// oss.Callback("ewogICAgImNhbGxiYWNrVXJsIjoiaHR0cDovLzEyMC43OS4yNS4xMTg6MTgwMTIvYXBpL3YxL2Nsb3Vkc3RvcmFnZS9jYWxsYmFjayIsCiAgICAiY2FsbGJhY2tCb2R5Ijoie1wiY29kZVwiOjB9Igp9Cg=="),
	}
	return t.bucket.PutObjectFromFile(name, file, options...)
}

// PutFileWithURL Create the URL that want to upload file to the OSS server
func (t AliyunOss) PutFileWithURL(filename string) (string, error) {
	signedURL, err := t.bucket.SignURL(filename, oss.HTTPPut, 24*60*60)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}

	return signedURL, nil
}

// Get the file from OSS server, and save local disk
func (this AliyunOss) GetFile(file string) {
	logs.Debug("filename: ", file)
	body, err := this.bucket.GetObject(file)
	if err != nil {
		logs.Error(err.Error())
	}

	defer body.Close()

	fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		logs.Error(err.Error())
	}
	defer fd.Close()

	io.Copy(fd, body)
}

// Create the URL that want to access file
func (this AliyunOss) GetFileWithURL(name string, timeout int64) (string, error) {
	signedURL, err := this.bucket.SignURL(name, oss.HTTPGet, timeout)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}

	// _, err = this.bucket.GetObjectWithURL(signedURL)
	// if err != nil {
	// 	logs.Error(err.Error())
	// 	return "", err
	// }
	logs.Debug(signedURL)
	return signedURL, nil
}

// Delete the file from the OSS server by filename
func (this AliyunOss) DeleteFile(file string) {
	// isExist, err := this.bucket.IsObjectExist(file)
	// if err != nil {
	// 	logs.Error(err.Error())
	// }
	// if isExist == false {
	// 	logs.Debug(file + " already exist!")
	// 	return true
	// }

	err := this.bucket.DeleteObject(file)
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Debug("delete ", file)

}

// Update remote file infomation
func (this AliyunOss) UpdateFile(object, key, value string) {
	err := this.bucket.SetObjectMeta(object, oss.Meta(key, value))
	if err != nil {
		logs.Error(err.Error())
	}
}
