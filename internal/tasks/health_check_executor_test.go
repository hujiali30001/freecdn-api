//go:build plus
// +build plus

package tasks_test

import (
	teaconst "github.com/hujiali30001/freecdn-api/internal/const"
	"github.com/hujiali30001/freecdn-api/internal/tasks"
	"github.com/iwind/TeaGo/dbs"
	"testing"
)

func TestHealthCheckExecutor_Run(t *testing.T) {
	teaconst.IsPlus = true
	dbs.NotifyReady()

	var executor = tasks.NewHealthCheckExecutor(42)
	results, err := executor.Run()
	if err != nil {
		t.Fatal(err)
	}
	for _, result := range results {
		t.Log(result.Node.Name, "addr:", result.NodeAddr, "isOk:", result.IsOk, "error:", result.Error)
	}
}
