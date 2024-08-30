package filesystem_utils

import (
	"os"
)

func CreateNewDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CreateNewFile(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
