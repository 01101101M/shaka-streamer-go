package configs

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var streamName = map[string]string{
	"video": "v",
	"audio": "a",
	"text":  "s",
}

func (in Input) getStreamSpecifier() string {
	return fmt.Sprintf("%s:%v", streamName[in.MediaType], in.trackNum)
}

func (in Input) ffprobe(show_entries string) (string, error) {

	fffpropPath, _ := exec.LookPath("ffprobe")
	arsg := []string{fffpropPath, in.Name,
		// Specifically, this stream
		"-select_streams", in.getStreamSpecifier(),
		// Show the needed metadata only
		"-show_entries", show_entries,
		// Print the metadata in a compact form, which is easier to parse
		"-of", "compact=p=0:nk=1",
	}

	c := exec.Cmd{
		Path: fffpropPath,
		Args: arsg,
		//Stderr: os.Stderr,
		//Stdout: os.Stdout,
	}
	fmt.Println(c.String())

	out, err := c.Output()

	return strings.TrimSpace(string(out)), err
}

func (in Input) IsInterlaced() (interlaced bool) {
	fieldOrder, _ := in.ffprobe("stream=field_order")
	_, interlaced = map[string]interface{}{"tt": nil,
		"bb": nil,
		"tb": nil,
		"bt": nil,
	}[fieldOrder]
	return
}

func (in Input) FrameRate() (frameRate float32) {

	frameRateString, err := in.ffprobe("stream=r_frame_rate")
	if err != nil {
		return
	}
	fraction := strings.Split(frameRateString, "/")
	if len(fraction) == 1 {
		i, _ := strconv.ParseFloat(fraction[0], 32)
		frameRate = float32(i)
	} else {
		a, _ := strconv.ParseFloat(fraction[0], 32)
		b, _ := strconv.ParseFloat(fraction[1], 32)

		frameRate = float32(a / b)
	}
	if in.Resolution.IsInterlaced {
		frameRate /= 2.0
	}

	return
}

func (in Input) VideResolution() (resolution string) {
	frameRateString, err := in.ffprobe("stream=width,height")
	if err != nil {
		return
	}
	fraction := strings.Split(frameRateString, "|")
	w, _ := strconv.ParseInt(fraction[0], 10, 32)
	h, _ := strconv.ParseInt(fraction[1], 10, 32)
	for k, res := range Bitrates.VideoResolutions {
		if int(w) <= res.MaxWidth && int(h) <= res.MaxHeight &&
			in.Resolution.FrameRate <= res.MaxFrameRate {
			return k
		}
	}
	return
}
