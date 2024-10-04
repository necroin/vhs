package config

import (
	"flag"
	"os"
	"runtime"
	"time"
	"vhs/src/network/lan"
)

type Log struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

type Config struct {
	Url            string `yaml:"url"`
	Hostname       string `yaml:"hostname"`
	Platform       string `yaml:"platform"`
	TopologyPort   string `yaml:"listen_port"`
	Log            Log    `yaml:"log"`
	PluginsDir     string `yaml:"plugins_dir"`
	RequestTimeout time.Duration
}

func Load() (*Config, error) {
	url := flag.String("url", lan.GetMyLanAddr()+ApplicationPort, "server url")
	listenPort := flag.String("listen-port", ListenPort, "server topology listen port")

	logPath := flag.String("log-path", "logs/log_"+time.Now().Format("2006-01-02T15:04:05")+".txt", "path to logs file")
	logLevel := flag.String("log-level", "info", "logs level (error, info, verbose, debug)")

	pluginsDir := flag.String("plugins", "./plugins", "directory with plugins")

	flag.Parse()

	hostname, _ := os.Hostname()
	platform := flag.String("platform", runtime.GOOS, "OS platform ('windows', 'linux', 'darwin', etc.)")

	config := &Config{
		Url:          *url,
		Hostname:     hostname,
		Platform:     *platform,
		TopologyPort: *listenPort,
		Log: Log{
			Path:  *logPath,
			Level: *logLevel,
		},
		PluginsDir:     *pluginsDir,
		RequestTimeout: RequestTimeout,
	}

	return config, nil
}
