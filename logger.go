package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Logger - io.Writer struct for console and file logging
type Logger struct {
	vocal    bool
	filePath string
}

var Log *log.Logger

// New - Creates a new Logger instance with writable log file
func New(filePath string, vocal bool) (*Logger, error) {
	return &Logger{vocal: vocal, filePath: filePath}, nil
}

func (logger *Logger) File() (*os.File, error) {
	filePath := logger.filePath
	dir := filepath.Dir(filePath)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
	}

	if err != nil {
		return nil, err
	}

	_, err = os.Stat(filePath)
	var file *os.File
	if os.IsNotExist(err) {
		file, err = os.Create(filePath)
	} else {
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	}

	if err != nil {
		return nil, err
	}

	return file, nil

}

func (logger *Logger) Write(p []byte) (int, error) {
	if logger.vocal {
		fmt.Print(string(p))
	}

	file, err := logger.File()
	//file, err := os.OpenFile(logger.filePath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	return file.Write(p)
}
