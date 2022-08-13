package runcommand

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type TestType struct {
	input          string
	expectedStdout string
	expectedErr    error
}

func TestRun(t *testing.T) {
	tests := createTests()

	for _, tt := range tests {
		gotStdout, _, gotErr := Run(tt.input)

		if tt.expectedErr == nil && gotErr != nil {
			t.Error("expected no error, got ", gotErr)
		} else if tt.expectedStdout == "" && gotStdout != nil {
			t.Errorf("expected stdout to be nil, got %s", tt.expectedStdout)
		} else if strings.TrimSpace(gotStdout.String()) != strings.TrimSpace(tt.expectedStdout) {
			t.Errorf("expected stdout to be %s, got %s", tt.expectedStdout, gotStdout)
		}
	}
}

func createTests() []TestType {
	currDir, _ := os.Getwd()
	t1 := TestType{
		input:          "pwd",
		expectedStdout: currDir,
		expectedErr:    nil,
	}

	t2 := TestType{
		input:          "ls -123",
		expectedStdout: "<nil>",
		expectedErr:    fmt.Errorf(""),
	}

	return []TestType{t1, t2}
}
