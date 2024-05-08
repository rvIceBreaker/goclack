package keyset

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/gopxl/beep"
	"github.com/rvIceBreaker/goclack/audio"
	"github.com/rvIceBreaker/goclack/config"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type KeySet struct {
	Config *config.KeySetConfig
	AudioFiles map[string]*beep.Buffer
}

func LoadKeySet(p string, audio *audio.Audio) (*KeySet, error) {
	switch path.Ext(p) {
		case ".zip":
			return implLoadKeySetFromPack(p, audio)
		case ".yaml":
			return implLoadKeySetFromConfig(p, audio)
	}
	return nil, fmt.Errorf("Attempt to load unknown file as KeySet, file was not a .yaml config or a .zip pack")
}

func implLoadKeySetFromPack(p string, audio *audio.Audio) (*KeySet, error) {
	zf,err := zip.OpenReader(p)
	if err != nil {
		return nil, err
	}
	defer zf.Close()

	keyset := KeySet{}
	keyset.AudioFiles = make(map[string]*beep.Buffer)

	keysetConfigFile, err := zipGetFile(zf, "keyset.yaml")
	if err != nil {
		return nil, err
	}
	configBytes,err := zipReadFile(keysetConfigFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configBytes, &keyset.Config)
	if err != nil {
		return nil, err
	}

	for _,f := range zf.File {
		if f.Name == "keyset.yaml" {
			continue
		}
		if f.FileInfo().IsDir() {
			continue
		}

		soundName := strings.Split(path.Base(f.FileInfo().Name()), path.Ext(f.FileInfo().Name()))[0]

		reader,err := f.Open()
		if err != nil {
			return nil, err
		}

		buffer,err := audio.LoadFileRaw(reader, soundName, path.Ext(f.Name))
		if err != nil {
			return nil, err
		}
		reader.Close()

		keyset.AudioFiles[soundName] = buffer
	}

	return &keyset, nil
}

func implLoadKeySetFromConfig(p string, audio *audio.Audio) (*KeySet, error) {
	keyset := KeySet{}
	keyset.AudioFiles = make(map[string]*beep.Buffer)
	rootPath := strings.Split(p, path.Base(p))[0]

	confFile,err := os.Open(p)
	if err != nil {
		return nil, err
	}
	confBytes,err := io.ReadAll(confFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(confBytes, &keyset.Config)
	if err != nil {
		return nil, err
	}

	audioFiles,err := os.ReadDir(path.Join(rootPath, keyset.Config.ContentPath))
	if err != nil {
		return nil, err
	}

	for _,file := range audioFiles {
		if file.IsDir() { continue }

		soundExt := path.Ext(file.Name())

		if soundExt != ".wav" && soundExt != ".mp3" && soundExt != ".ogg" {
			continue
		}
		
		soundPath := path.Join(rootPath, keyset.Config.ContentPath, file.Name())
		soundName := strings.Split(path.Base(file.Name()), path.Ext(file.Name()))[0]

		fileHandle,err := os.Open(soundPath)
		if err != nil {
			return nil, err
		}

		buffer,err := audio.LoadFileRaw(fileHandle, soundName, soundExt)
		if err != nil {
			return nil, err
		}

		keyset.AudioFiles[soundName] = buffer
	}

	return &keyset, nil
}

func zipReadFile(file *zip.File) ([]byte, error) {
	fc,err := file.Open()
	if err != nil {
		return nil, err
	}

	cont,err := io.ReadAll(fc)
	if err != nil {
		return nil, err
	}

	return cont, nil
}

func zipGetFile(zf *zip.ReadCloser, file string) (*zip.File, error) {
	idx := slices.IndexFunc(zf.File, func(f *zip.File) bool {
		return f.Name == file
	})

	if idx < 0 {
		return nil, fmt.Errorf("File '%s' not found in zip", file)
	}

	return zf.File[idx], nil
}