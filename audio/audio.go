package audio

import (
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
	"github.com/gopxl/beep/wav"
)

type Audio struct {
	OutputSampleRate beep.SampleRate
	OutputVolume float64
	AudioBuffers map[string]*beep.Buffer
}

func NewAudio(SampleRate beep.SampleRate) *Audio {
	audio := Audio{
		OutputSampleRate: SampleRate,
		AudioBuffers: make(map[string]*beep.Buffer),
		OutputVolume: 100.0,
	}
	audio.init()
	return &audio
}

func (a *Audio) init() {
	err := speaker.Init(a.OutputSampleRate, a.OutputSampleRate.N(time.Second/30))
	if err != nil {
		panic(err)
	}
}

func (a *Audio) calculateVolume() float64 {
	// I don't know if this calculation is correct at all
	// OutputVolume is a value from 0 - 100, where 0 is -24db in relative volume
	// beep is weird about how volume is set, and I don't know if the settings actually correlate
	// Either way, it feels relatively natural and adjustable with this
	return math.Abs((a.OutputVolume / 100.0) - 1) * -2.40
}

func (a *Audio) LoadFile(filePath string) (error) {
	var err error
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	var str beep.StreamSeekCloser
	var format beep.Format

	fileExt := path.Ext(filePath)
	switch fileExt {
		case ".wav":
			str, format, err = wav.Decode(f)
			break;
		case ".mp3":
			str, format, err = mp3.Decode(f)
		case ".ogg":
			str, format, err = vorbis.Decode(f)
		default:
			return fmt.Errorf("Could not determine file format for %s", f.Name())
	}
	if err != nil {
		return err
	}
	defer str.Close()

	buffer := beep.NewBuffer(format)
	buffer.Append(str)

	a.AudioBuffers[strings.Split(path.Base(f.Name()), path.Ext(f.Name()))[0]] = buffer
	return nil
}

func (a *Audio) LoadFileRaw(fileReader io.ReadCloser, soundName string, fileExt string) (*beep.Buffer,error) {
	var str beep.StreamSeekCloser
	var format beep.Format
	var err error

	switch fileExt {
		case ".wav":
			str, format, err = wav.Decode(fileReader)
			break;
		case ".mp3":
			str, format, err = mp3.Decode(fileReader)
		case ".ogg":
			str, format, err = vorbis.Decode(fileReader)
	}
	if err != nil {
		return nil, err
	}
	defer str.Close()

	buffer := beep.NewBuffer(format)
	buffer.Append(str)

	a.AudioBuffers[soundName] = buffer
	return buffer, nil
}

func (a *Audio) LoadDir(dir string) (error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _,f := range files {
		if err := a.LoadFile(path.Join(dir, f.Name())); err != nil {
			return err
		}
	}
	return nil
}

func (a *Audio) PlaySound(name string) (error) {
	buffer,ok := a.AudioBuffers[name]
	if ok == false {
		return fmt.Errorf("Sound name %s not cached", name)
	}

	resample := beep.Resample(
		4,
		buffer.Format().SampleRate,
		a.OutputSampleRate,
		buffer.Streamer(0, buffer.Len()),
	)

	vol := &effects.Volume{
		Streamer: resample,
		Base: 10,
		Volume: a.calculateVolume(),
		Silent: false,
	}

	speaker.Play(vol)
	return nil
}

func (a *Audio) Teardown() {
	speaker.Close()
}