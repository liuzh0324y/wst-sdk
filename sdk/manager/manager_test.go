package manager_test

import (
	"testing"

	"github.com/wst-libs/wst-sdk/sdk/manager"
)

const (
	url = "http://10.33.48.20:8083/api/v1/storage/file"
)

func Test_Manager(t *testing.T) {
	obj := manager.NewManager()
	if obj != nil {

	}
}

func Test_Manager_Add(t *testing.T) {
	obj := manager.NewManager()
	ret := obj.Add(url, nil)
	if ret.Code != 0 {
		t.Log("?????")
	}
}

func Test_Manager_Del(t *testing.T) {
	obj := manager.NewManager()
	obj.Del()
}

func Test_Manager_Update(t *testing.T) {
	obj := manager.NewManager()
	obj.Update(url)
}

func Test_Manager_Get(t *testing.T) {
	obj := manager.NewManager()
	obj.Get()
}
