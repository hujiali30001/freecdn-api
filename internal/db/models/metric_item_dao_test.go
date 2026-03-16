package models_test

import (
	"github.com/hujiali30001/freecdn-api/internal/db/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"testing"
)

func TestMetricStatDAO_Clean(t *testing.T) {
	var dao = models.NewMetricStatDAO()
	t.Log(dao.Clean(nil))
}
