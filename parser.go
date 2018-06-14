package gkin

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Gkin is top level
type Gkin struct {
	Pipeline []Pipe `yaml:"pipeline"`
}

// Pipe is one of pipeline
type Pipe struct {
	Image    string   `yaml:"image"`
	Name     string   `yaml:"name"`
	Commands []string `yaml:"commands"`
}

// Parse do parsing .gkin.yml file to struct
func Parse(name string) (Gkin, error) {
	var gkin Gkin
	in, err := os.Open(name)
	if err != nil {
		return gkin, err
	}
	decoder := yaml.NewDecoder(in)
	err = decoder.Decode(&gkin)
	return gkin, err
}
