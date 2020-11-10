package main

import (
	"fmt"

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

	err = args.Process()

}
