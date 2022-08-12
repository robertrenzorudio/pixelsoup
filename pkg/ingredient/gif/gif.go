package gif

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/google/shlex"
	"github.com/robertrenzorudio/pixelsoup/pkg/ingredient/mediainfo"
)

type VidToGifInput struct {
	Duration   float64
	Fps        int
	OutGifName string
	Scale      int
	Start      float64
	InVidName  string
}

type Gif struct {
	MaxDuration float64
}

func New(maxDuration float64) *Gif {
	return &Gif{MaxDuration: maxDuration}
}

func (g *Gif) VidToGif(info *VidToGifInput) (outFileName string, err error) {
	err = checkInput(info, g.MaxDuration)

	if err != nil {
		return "", err
	}

	outFileName = info.OutGifName + ".gif"
	cmdStr := fmt.Sprintf("ffmpeg -hide_banner -loglevel error -y -ss %.2f -t %.2f -i %s -filter_complex '[0:v] fps=%d,scale=%d:-1,split [a][b];[a] palettegen [p];[b][p] paletteuse' %s",
		info.Start, info.Duration, info.InVidName, info.Fps, info.Scale, outFileName)

	args, err := shlex.Split(cmdStr)

	if err != nil {
		return "", err
	}

	cmd := exec.Command(args[0], args[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		os.Remove(outFileName)
		return "", err
	}

	cmdErr := stderr.String()
	if cmdErr != "" && cmdErr != "<nil>" {
		os.Remove(outFileName)
		return "", errors.New(cmdErr)
	}

	return outFileName, nil
}

func checkInput(info *VidToGifInput, maxDuration float64) error {
	format, err := mediainfo.GetFormat(info.InVidName)
	if err != nil {
		return err
	}

	vidDuration, err := strconv.ParseFloat(format.Duration, 32)
	if err != nil {
		return err
	}

	if info.Duration < 0 || info.Duration > maxDuration {
		return &ErrInputParameterValue{
			field:  "Duration",
			reason: fmt.Sprintf("got Duration = %.2f, expected 0 <= Duration <= %.2f", info.Duration, maxDuration),
		}
	}

	if info.Start < 0 || info.Start >= vidDuration {
		return &ErrInputParameterValue{
			field:  "Start",
			reason: fmt.Sprintf("got Start = %.2f, 0 <= Start < Input Video Duration = %.2f", info.Start, vidDuration),
		}
	}

	return nil
}
