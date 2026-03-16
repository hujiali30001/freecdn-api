// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package tasks

import (
	"github.com/hujiali30001/freecdn-api/internal/db/models"
	"github.com/hujiali30001/freecdn-api/internal/remotelogs"
)

type BaseTask struct {
}

func (this *BaseTask) logErr(taskType string, errString string) {
	remotelogs.Error("TASK", "run '"+taskType+"' failed: "+errString)
}

func (this *BaseTask) IsPrimaryNode() bool {
	return models.SharedAPINodeDAO.CheckAPINodeIsPrimaryWithoutErr()
}
