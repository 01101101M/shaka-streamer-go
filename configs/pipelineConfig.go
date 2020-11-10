package configs

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

type PipelineConfig struct {
	StreamingMode      string   `yaml:"streaming_mode"`
	Quiet              bool     `yaml:"quiet"`
	DebugLogs          bool     `yaml:"debug_logs"`
	HwaccelApi         string   `yaml:"hwaccel_api"`
	Resolutions        []string `yaml:"resolutions"`
	Channels           int64    `yaml:"channels"`
	AudioCodecs        []string `yaml:"audio_codecs"`
	VideoCodecs        []string `yaml:"video_codecs"`
	ManifestFormat     []string `yaml:"manifest_format"`
	DashOutput         string   `yaml:"dash_output"`
	HlsOutput          string   `yaml:"hls_output"`
	SegmentFolder      string   `yaml:"segment_folder"`
	SegmentSize        float64  `yaml:"segment_size"`
	SegmentPerFile     bool     `yaml:"segment_per_file"`
	AvailabilityWindow int64    `yaml:"availability_window"`
	PresentationDelay  int64    `yaml:"presentation_delay"`
	UpdatePeriod       int64    `yaml:"update_period"`
	Encryption         struct {
		Enable           bool   `yaml:"enable"`
		ContentId        string `yaml:"content_id"`
		KeyServerUrl     string `yaml:"key_server_url"`
		Signer           string `yaml:"signer"`
		SigningKey       string `yaml:"signing_key"`
		SigningIv        string `yaml:"signing_iv"`
		ProtectionScheme string `yaml:"protection_scheme"`
		ClearLead        int64  `yaml:"clear_lead"`
	} `yaml:"encryption"`
}

func (pc PipelineConfig) GetResolutions() []string {
	return pc.Resolutions
}

func PipelineConfigUnmarshal(path string) (ic PipelineConfig, err error) {
	pipelineyml, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	yaml.Unmarshal(pipelineyml, &ic)
	return
}
