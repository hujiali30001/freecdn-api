// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package models

import (
	"github.com/hujiali30001/freecdn-common/pkg/nodeconfigs"
	"github.com/iwind/TeaGo/dbs"
)

// 节点值变更Hook
func (this *NodeValueDAO) nodeValueHook(tx *dbs.Tx, role nodeconfigs.NodeRole, nodeId int64, item nodeconfigs.NodeValueItem, valueJSON []byte) error {
	return nil
}
