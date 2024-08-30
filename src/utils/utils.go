package utils

import "strings"

func GetProcessRunCommand(platform, path string, args ...string) (string, []string) {
	if platform == "windows" {
		return "cmd", []string{"/C", strings.ReplaceAll(path, "/", "\\")}
	}
	return "open", []string{path}
}
