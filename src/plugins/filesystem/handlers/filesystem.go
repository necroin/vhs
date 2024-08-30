package filesystem_handlers

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"vhs/src/message"
	"vhs/src/network/http"
	plugins_core "vhs/src/plugins/core"
)

type FileInfo struct {
	ModTime  string `json:"mod_time"`
	Size     int64  `json:"size"`
	Url      string `json:"url"`
	Platform string `json:"platform"`
	Hostname string `json:"hostname"`
}

type FilesystemDirectory struct {
	Directories map[string]FileInfo   `json:"directories"`
	Files       map[string][]FileInfo `json:"files"`
}

func CollectSelfFilesystem(self *message.HostInfo, walkPath string) *FilesystemDirectory {
	walkPath = path.Clean(walkPath)
	if walkPath == "" {
		walkPath = "/"
	}

	filesystemDirectory := &FilesystemDirectory{
		Directories: map[string]FileInfo{},
		Files:       map[string][]FileInfo{},
	}

	entries, err := os.ReadDir(walkPath)
	if err != nil {
		return filesystemDirectory
	}

	for _, entry := range entries {
		stat, err := os.Stat(path.Join(walkPath, entry.Name()))
		if err != nil {
			continue
		}

		size := stat.Size() / 1000
		if size == 0 {
			size = 1
		}

		info := FileInfo{
			ModTime:  stat.ModTime().Format("02.01.2006 15:04"),
			Size:     size,
			Url:      self.Url,
			Platform: self.Platform,
			Hostname: self.Hostname,
		}

		if entry.IsDir() {
			filesystemDirectory.Directories[entry.Name()] = info
		} else {
			filesystemDirectory.Files[entry.Name()] = append(filesystemDirectory.Files[entry.Name()], info)
		}
	}

	return filesystemDirectory
}

func CollectHostFileSystem(host *message.HostInfo, data []byte) *FilesystemDirectory {
	result := &FilesystemDirectory{}

	response, err := http.SendPostRequest(
		host.Url+"/filesystem/self",
		data,
	)
	if err != nil {
		return result
	}

	json.NewDecoder(response.Body).Decode(result)
	return result
}

func FilesystemSelfHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	result := CollectSelfFilesystem(clusterInfo.Self, string(data))
	return json.NewEncoder(out).Encode(result)
}

func FilesystemAllHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	result := FilesystemDirectory{
		Directories: map[string]FileInfo{},
		Files:       map[string][]FileInfo{},
	}

	for _, host := range clusterInfo.Hosts {
		storageFilesystem := CollectHostFileSystem(
			host,
			data,
		)
		for directory, info := range storageFilesystem.Directories {
			result.Directories[directory] = info
		}

		for file, info := range storageFilesystem.Files {
			result.Files[file] = append(result.Files[file], info...)
		}
	}

	return json.NewEncoder(out).Encode(result)
}
