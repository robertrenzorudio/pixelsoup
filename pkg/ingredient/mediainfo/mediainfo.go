package mediainfo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	"github.com/google/shlex"
)

type Format struct {
	BitRate        string `json:"bit_rate"`
	Duration       string `json:"duration"`
	Filename       string `json:"filename"`
	FormatLongName string `json:"format_long_name"`
	FormatName     string `json:"format_name"`
	NbPrograms     int64  `json:"nb_programs"`
	NbStreams      int64  `json:"nb_streams"`
	ProbeScore     int64  `json:"probe_score"`
	Size           string `json:"size"`
	StartTime      string `json:"start_time"`
}

type outerFormat struct {
	Outer Format `json:"format"`
}

func GetFormat(fileName string) (*Format, error) {
	cmdStr := fmt.Sprintf("ffprobe -v error -print_format json -show_format %s", fileName)

	args, err := shlex.Split(cmdStr)

	if err != nil {
		return nil, err
	}

	cmd := exec.Command(args[0], args[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		return nil, err
	}

	cmdErr := stderr.String()
	if cmdErr != "" && cmdErr != "<nil>" {
		return nil, errors.New(cmdErr)
	}

	outer := &outerFormat{}
	json.Unmarshal(stdout.Bytes(), outer)
	format := outer.Outer
	return &format, nil
}
