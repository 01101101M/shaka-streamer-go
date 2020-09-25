package configs

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

type InputConfig struct {
	Inputs []struct {
		Name      string `yaml:"name"`
		MediaType string `yaml:"media_type"`
		Language  string `yaml:"language"`
	} `yaml:"inputs"`
}

func InputConfigUnmarshal(path string) (ic InputConfig, err error) {
	inputyml, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	yaml.Unmarshal(inputyml, &ic)
	return
}
