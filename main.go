package main

import (
	"fmt"

	"shaka-streamer/configs"
	"shaka-streamer/controler"

	"github.com/jessevdk/go-flags"
)

var (
	args controler.Args
)

func main() {
	_, err := flags.Parse(&args)
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

	err = controler.Process(args)

	fmt.Println(inputConfig, pipelineConfig)
}
