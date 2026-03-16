package models

import (
	"encoding/json"
	"github.com/hujiali30001/freecdn-api/internal/remotelogs"
)

func (this *UpdatingServerList) DecodeServerIds() []int64 {
	if len(this.ServerIds) == 0 {
		return nil
	}

	var serverIds = []int64{}
	err := json.Unmarshal(this.ServerIds, &serverIds)
	if err != nil {
		remotelogs.Error("UpdatingServerList", "DecodeServerIds(): "+err.Error())
	}

	return serverIds
}
