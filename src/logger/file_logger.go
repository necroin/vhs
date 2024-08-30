package logger

import (
	"fmt"
	"os"
	"path"
)

func NewFileLogger(filePath string, level int) (*Logger, error) {
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return nil, fmt.Errorf("[File Logger] failed create logs dicrectory: %s", err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("[File Logger] failed open file: %s", err)
	}

	return NewLogger(file, level), nil
}
