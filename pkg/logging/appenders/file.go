package appenders

import (
	"fmt"
	"logger/pkg/logging"
	"os"
)

// FileAppender writes logs to a file
type FileAppender struct {
	filepath string
	file     *os.File
}

func NewFileAppender(filepath string) *FileAppender {
	return &FileAppender{
		filepath: filepath,
	}
}

func (f *FileAppender) Setup() error {
	var err error
	f.file, err = os.OpenFile(f.filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return err
}

func (f *FileAppender) Append(log *logging.Log) error {
	if f.file == nil {
		return fmt.Errorf("file appender not properly setup")
	}
	_, err := f.file.WriteString(log.String() + "\n")
	return err
}

func (f *FileAppender) Close() error {
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}
