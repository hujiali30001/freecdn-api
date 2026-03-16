package utils

import (
	teaconst "github.com/hujiali30001/freecdn-api/internal/const"
	"testing"
)

func TestServiceManager_Log(t *testing.T) {
	manager := NewServiceManager(teaconst.ProductName, teaconst.ProductName+" Server")
	manager.Log("Hello, World")
	manager.LogError("Hello, World")
}
