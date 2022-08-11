package mediainfo

import (
	"errors"
	"testing"

	"github.com/robertrenzorudio/pixelsoup/test"
)

func TestGetFormat(t *testing.T) {
	someError := errors.New("some error")
	tests := []struct {
		input       string
		expected    *Format
		expectedErr error
	}{
		{"test_video.mov", &Format{}, nil},
		{"FILE_DNE.mp4", nil, someError},
		{"test_invalid_file.txt", nil, someError},
	}

	for _, tt := range tests {
		testFile := test.GetTestDataDir(tt.input)
		expected := tt.expected
		expectedErr := tt.expectedErr

		got, gotErr := GetFormat(testFile)

		if expectedErr == nil && gotErr != nil {
			t.Error("expected no error")
		}

		if expected != nil && got == nil {
			t.Error("expected Format struct, got nil")
		}
	}
}
