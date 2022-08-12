package gif

import (
	"errors"
	"math"
	"os"
	"testing"

	"github.com/robertrenzorudio/pixelsoup/config"
	"github.com/robertrenzorudio/pixelsoup/test"
)

func TestVidToGif(t *testing.T) {
	tests := getVidToGifTestsValid()
	vtg := New(config.MAX_GIF_DURATION)
	for _, tt := range tests {
		input := tt.input
		expected := tt.expected
		expectedErr := tt.expectedErr

		got, gotErr := vtg.VidToGif(input)

		if expectedErr == nil && gotErr != nil {
			t.Error("expected no error, got ", gotErr)
		}

		if expected != got {
			t.Errorf("expected gif file name to be %s, got %s", expected, got)
		}

		_, err := os.Stat(expected)

		if errors.Is(err, os.ErrNotExist) {
			t.Errorf("expected gif file, got none")
		}

		os.Remove(got)

	}
}

func TestVidToGifInvalid(t *testing.T) {
	tests := getVidToGifTestsInvalid()
	vtg := New(config.MAX_GIF_DURATION)
	for _, tt := range tests {
		input := tt.input
		expectedErr := tt.expectedErr

		got, gotErr := vtg.VidToGif(input)

		if expectedErr != nil && gotErr == nil {
			t.Error("expected error, got none")
			os.Remove(got)
		}
	}
}

func getVidToGifTestsValid() []struct {
	input       *VidToGifInput
	expected    string
	expectedErr error
} {
	valid0 := VidToGifInput{
		InVidName:  test.GetTestDataDir("test_video.mov"),
		Start:      0,
		Duration:   2,
		Fps:        15,
		Scale:      480,
		OutGifName: "out",
	}

	tests := []struct {
		input       *VidToGifInput
		expected    string
		expectedErr error
	}{
		{&valid0, valid0.OutGifName + ".gif", nil},
	}

	return tests
}

func getVidToGifTestsInvalid() []struct {
	input       *VidToGifInput
	expectedErr error
} {
	// duration < 0
	valid := VidToGifInput{
		InVidName:  test.GetTestDataDir("test_video.mov"),
		Start:      0,
		Duration:   config.MAX_GIF_DURATION,
		Fps:        15,
		Scale:      480,
		OutGifName: "shouldNE",
	}

	// duration < 0
	invalid0 := valid
	invalid0.Duration = -1

	// duration > max
	invalid1 := valid
	invalid1.Duration = math.Inf(1)

	// start > exceed duration
	invalid2 := valid
	invalid2.Start = math.Inf(1)

	// start < 0
	invalid3 := valid
	invalid3.Start = -1

	// invalid file
	invalid4 := valid
	invalid4.InVidName = "test_invalid_file.txt"

	inputError := &ErrInputParameterValue{}
	tests := []struct {
		input       *VidToGifInput
		expectedErr error
	}{
		{&invalid0, inputError},
		{&invalid1, inputError},
		{&invalid2, inputError},
		{&invalid3, inputError},
		{&invalid4, inputError},
	}

	return tests
}
