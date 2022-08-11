package logger

import (
	"errors"
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
		t.Errorf("expected log file, got none")
	}

	if logFile.Size() == 0 {
		t.Errorf("expected log entry, got none")
	}

	os.Remove(fileName)
}
