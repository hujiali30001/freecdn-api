package models

import (
	"encoding/json"
	"github.com/hujiali30001/freecdn-api/internal/remotelogs"
)

// DecodeModules 解析模块
func (this *User) DecodeModules() []string {
	if len(this.Modules) == 0 {
		return nil
	}

	var result = []string{}
	err := json.Unmarshal(this.Modules, &result)
	if err != nil {
		remotelogs.Error("User.DecodeModules", err.Error())
	}

	return result
}
