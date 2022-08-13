package mediainfo

import (
	"encoding/json"
	"fmt"

	"github.com/robertrenzorudio/pixelsoup/pkg/runcommand"
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
	command := fmt.Sprintf("ffprobe -v error -print_format json -show_format %s", fileName)

	stdout, stderr, err := runcommand.Run(command)

	if err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), stderr.String())
	}

	outer := &outerFormat{}
	json.Unmarshal(stdout.Bytes(), outer)
	format := outer.Outer
	return &format, nil
}
