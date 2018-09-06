package manager_test

import (
	"testing"

	"github.com/wst-libs/wst-sdk/sdk/manager"
)

const (
	url = "http://39.105.49.69:48083/api/v1/storage/file"
)

// func TestAdd(t *testing.T) {
// 	manager.Add(url)
// }
func TestPut(t *testing.T) {
	manager.Put(url)
}
