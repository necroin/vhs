package plugins_core

import "flag"

type PluginFlags struct {
	Call string
}

func ParseFlags() PluginFlags {
	callFlag := flag.String("call", "", "")
	flag.Parse()
	return PluginFlags{
		Call: *callFlag,
	}
}
