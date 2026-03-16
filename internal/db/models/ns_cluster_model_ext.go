package models

import (
	"encoding/json"
	"github.com/hujiali30001/freecdn-api/internal/remotelogs"
	"github.com/hujiali30001/freecdn-common/pkg/serverconfigs/ddosconfigs"
)

// DecodeDDoSProtection 解析DDOS Protection设置
func (this *NSCluster) DecodeDDoSProtection() *ddosconfigs.ProtectionConfig {
	if IsNull(this.DdosProtection) {
		return nil
	}

	var result = &ddosconfigs.ProtectionConfig{}
	err := json.Unmarshal(this.DdosProtection, &result)
	if err != nil {
		remotelogs.Error("NSCluster.DecodeDDoSProtection", "decode failed: "+err.Error())
	}
	return result
}

// HasDDoSProtection 检查是否有DDOS设置
func (this *NSCluster) HasDDoSProtection() bool {
	var config = this.DecodeDDoSProtection()
	if config != nil {
		return config.IsOn()
	}
	return false
}
