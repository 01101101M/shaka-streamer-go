package configs

import (
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

type Input struct {
	Name       string `yaml:"name"`
	MediaType  string `yaml:"media_type"`
	Language   string `yaml:"language"`
	Resolution InputResolution
	trackNum   int
}

type InputConfig struct {
	Inputs []Input `yaml:"inputs"`
}
type InputResolution struct {
	IsInterlaced bool
	FrameRate    float32
	Resolution   string
}

func (in Input) prob() {
	if _, notSupported := DEFAULT_AUDIO_CHANNEL_LAYOUTS[in.MediaType]; notSupported {
		return
	}
	in.Resolution.IsInterlaced = in.IsInterlaced()
	in.Resolution.FrameRate = in.FrameRate()
	in.Resolution.Resolution = in.VideResolution()
	fmt.Println("prob", in)
}

func (in Input) GetResolution() (resolution InputResolution) {
	if in.Resolution == (InputResolution{}) {
		in.prob()
	}
	return
}

func InputConfigUnmarshal(path string) (ic InputConfig, err error) {
	inputyml, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	yaml.Unmarshal(inputyml, &ic)
	return
}
