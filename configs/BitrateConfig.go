package configs

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

type AudioChannelLayout struct {
	MaxChannels int               `yaml:"max_channels"`
	Bitrates    map[string]string `yaml:"bitrates"`
}

type VideoResolution struct {
	MaxWidth     int               `yaml:"max_width"`
	MaxHeight    int               `yaml:"max_height"`
	MaxFrameRate float64           `yaml:"max_frame_rate"`
	Bitrates     map[string]string `yaml:"bitrates"`
}

type BitrateConfig struct {
	AudioChannelLayouts map[string]AudioChannelLayout `yaml:"audio_channel_layouts"`
	VideoResolutions    map[string]VideoResolution    `yaml:"video_resolutions"`
}

func BitrateConfigUnmarshal(path string) (ic BitrateConfig) {
	bitratesyml, err := ioutil.ReadFile(path)
	if err != nil {
		yaml.Unmarshal(bitratesyml, &ic)
	}

	if ic.AudioChannelLayouts == nil {
		ic.AudioChannelLayouts = DEFAULT_AUDIO_CHANNEL_LAYOUTS
	}

	if ic.VideoResolutions == nil {
		ic.VideoResolutions = DEFAULT_VIDEO_RESOLUTIONS
	}

	return
}
