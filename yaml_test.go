package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYaml(t *testing.T) {
	type TestStr struct {
		Content string `yaml:"content"`
		Events struct {
			Keys struct {
				KeyDown map[string][]string `yaml:"key_down"`
				KeyUp map[string][]string `yaml:"key_up"`
			} `yaml:"keys"`
		} `yaml:"events"`
	}


	f,err := os.Open("./test.yaml")
	if err != nil {
		panic(err)
	}

	b,err := io.ReadAll(f)

	v := TestStr{}

	err = yaml.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", v.Events.Keys.KeyDown["default"])
}