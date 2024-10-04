package config

import "time"

const (
	ApplicationPort       = ":3300"
	ListenPort            = "3301"
	InstanceRemoveTimeout = 30 * time.Second
	RequestTimeout        = 5 * time.Second
	LanTimeout            = 2 * RequestTimeout
)
