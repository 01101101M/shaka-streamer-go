package controler

import (
	"fmt"
	"shaka-streamer/configs"
	"shaka-streamer/internal"
)

type Args struct {
	InputConfig    string `short:"i" long:"input_config" description:"The path to the input config file (required)." required:"true"`
	PipelineConfig string `short:"p" long:"pipeline_config" description:"The path to the pipeline config file (required)." required:"true"`
	BitrateConfig  string `short:"b" long:"bitrate_config" description:"The path to a config file which defines custom bitrates and resolutions for transcoding.  (optional, see example in config/config_files/bitrate_config.yaml)"`
	//CloudUrl       string `short:"c" long:"cloud_url" description:"The Google Cloud Storage or Amazon S3 URL to  upload to.  (Starts with gs:// or s3://)"`
	Output        string `short:"o" long:"output" description:"The output folder to write files to. Used even if uploading to cloud storage."`
	SkipDepsCheck bool   `short:"s" long:"skip_deps_check" description:"Skip checks for dependencies and their versions.  This can be useful for testing pre-release versions of FFmpeg or Shaka Packager."`
}

func (args Args) Process() (err error) {

	if !args.SkipDepsCheck {
		// Check that ffmpeg version is 4.1 or above.
		err := internal.CheckVersion("FFmpeg", []string{"ffmpeg", "-version"}, "4.1.0")
		if err != nil {
			panic(err)
		}

		// Check that ffprobe version (used for autodetect features) is 4.1 or
		// above.
		err = internal.CheckVersion("ffprobe", []string{"ffprobe", "-version"}, "4.1.0")
		if err != nil {
			panic(err)
		}

		// Check that Shaka Packager version is 2.4.2 or above.
		err = internal.CheckVersion("Shaka Packager", []string{"packager", "-version"}, "2.4.2")
		if err != nil {
			panic(err)
		}
	}

	// Define resolutions and bitrates before parsing other configs.
	bitrateConfig := configs.BitrateConfigUnmarshal(args.BitrateConfig)
	fmt.Println(bitrateConfig, err)
	if err != nil {
		fmt.Println(err)
	}

	inputConfig, err := configs.InputConfigUnmarshal(args.InputConfig)
	if err != nil {
		fmt.Println(err)
	}

	pipelineConfig, err := configs.PipelineConfigUnmarshal(args.PipelineConfig)
	if err != nil {
		fmt.Println(err)
	}

	_ = pipelineConfig

	for _, input := range inputConfig.Inputs {
		switch input.MediaType {
		case "video":
			for _, videoCodec := range pipelineConfig.VideoCodecs {
				for _, outputResolution := range pipelineConfig.GetResolutions() {
					_ = videoCodec
					videoRes := bitrateConfig.VideoResolutions[outputResolution]
					fmt.Println(input.GetResolution(), videoRes)
				}
			}
		case "audio":
			for _, audioCodec := range pipelineConfig.AudioCodecs {
				// path, err := internal.CreatePipe("")
				_ = audioCodec
				// fmt.Println(path, err, audioCodec)
			}
		default:
			// fmt.Println("Nai: ", input)
		}
	}

	return nil
}
