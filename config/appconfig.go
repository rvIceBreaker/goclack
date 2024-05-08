package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	App AppConfigApp `yaml:"app"`
}

type AppConfigApp struct {
	GlobalVolume float64 `yaml:"global_volume"`
	DefaultSet string `yaml:"default_set"`
}

func NewAppConfig() *AppConfig {
	config := AppConfig{
		App: AppConfigApp{
			GlobalVolume: 100.0,
		},
	}

	return &config
}

func LoadConfig() (*AppConfig, error) {
	config := AppConfig{}
	cwd,_ := os.Getwd()

	f, err := os.ReadFile(path.Join(cwd, "settings.yaml"))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *AppConfig) SaveConfig() (error) {
	cwd,_ := os.Getwd()

	d, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(cwd, "settings.yaml"))
	if err != nil {
		return err
	}
	defer f.Close()

	n,err := f.Write(d)
	if err != nil || n != len(d) {
		return err
	}

	return nil
}