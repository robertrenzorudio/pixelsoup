package logger

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestFileLogging(t *testing.T) {
	fileName := "test_log.json"
	os.Remove(fileName)

	config := &LogConfig{Core: File, FileName: fileName}
	New(config)
	Log.Info("Hello from test")

	logFile, err := os.Stat(fileName)

	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected a log file but non is created")
	}

	if logFile.Size() == 0 {
		fmt.Println(logFile.Size())
		t.Errorf("expected a log entry but non is entered")
	}

	os.Remove(fileName)
}
