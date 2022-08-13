package gif

import (
	"fmt"
	"os"
	"strconv"

	"github.com/robertrenzorudio/pixelsoup/pkg/ingredient/mediainfo"
	"github.com/robertrenzorudio/pixelsoup/pkg/runcommand"
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

	base := "ffmpeg -hide_banner -loglevel error -y"
	input := fmt.Sprintf("-ss %.2f -t %.2f -i %s", info.Start, info.Duration, info.InVidName)
	filterComplex := fmt.Sprintf(
		"-filter_complex '[0:v] fps=%d,scale=%d:-1,split [a][b];[a] palettegen [p];[b][p] paletteuse'",
		info.Fps, info.Scale,
	)
	outFileName = info.OutGifName + ".gif"

	command := fmt.Sprintf("%s %s %s %s", base, input, filterComplex, outFileName)

	_, stderr, err := runcommand.Run(command)

	if err != nil {
		os.Remove(outFileName)
		return "", fmt.Errorf("%s: %s", err.Error(), stderr.String())
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
