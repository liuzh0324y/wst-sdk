package omg

import (
	"log"
	"time"

	minio "github.com/minio/minio-go"
)

func fileHandler(bucket, object string) []byte {
	c, err := minio.New(endpoint, accesskey, secretkey, false)
	if err != nil {
		log.Println("error: ", err.Error())
	}
	// obj, err := c.GetObject(bucket, object, minio.GetObjectOptions{})
	url, err := c.PresignedGetObject(bucket, object, time.Second*60*60*24, nil)
	if err != nil {
		log.Println("error: ", err.Error())
	}
	log.Println(url)
	return []byte{}
}
