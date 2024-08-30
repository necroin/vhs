package plugins_core

import "vhs/src/message"

type ClusterInfo struct {
	Self  *message.HostInfo
	Hosts map[string]*message.HostInfo
}
