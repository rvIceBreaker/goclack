package main

import (
	"math/rand"
	"os"
	"path"
	"sync"
	"testing"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func TestHelloBeep(t *testing.T) {
	var outSampleRate = beep.SampleRate(44100)

	err := speaker.Init(outSampleRate, 2048)
	if err != nil {
		panic(err)
	}

	cwd,_ := os.Getwd()
	f, err := os.Open(path.Join(cwd, "aud", "KeyboardPress002_002.wav"))
	if err != nil {
		panic(err)
	}

	str, format, err := wav.Decode(f)
	if err != nil {
		panic(err)
	}
	defer str.Close()

	speaker.Play(beep.Resample(4, format.SampleRate, outSampleRate, str))

	time.Sleep(1 * time.Second)
}

func TestRandomPlayback(t *testing.T) {
	var outSampleRate = beep.SampleRate(44100)
	cwd,_ := os.Getwd()

	err := speaker.Init(outSampleRate, outSampleRate.N(time.Second/30))
	if err != nil {
		panic(err)
	}

	//Load and decode a .wav and return a buffer for later playback
	loadFile := func(filePath string) (*beep.Buffer) {
		f, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		str, format, err := wav.Decode(f)
		if err != nil {
			panic(err)
		}
		defer str.Close()

		buffer := beep.NewBuffer(format)
		buffer.Append(str)

		return buffer
	}

	playBuffer := func(wg *sync.WaitGroup, delay time.Duration, buffer *beep.Buffer) {
		defer wg.Done()
		time.Sleep(delay)

		speaker.Play(
			beep.Resample(
				4,
				buffer.Format().SampleRate,
				outSampleRate,
				buffer.Streamer(0, buffer.Len()),
				))
	}

	var fileBuffers []*beep.Buffer

	files, err := os.ReadDir(path.Join(cwd, "aud"))
	if err != nil {
		panic(err)
	}

	for _,f := range files {
		fileBuffers = append(fileBuffers, loadFile(path.Join(cwd, "aud", f.Name())))
	}

	wg := new(sync.WaitGroup)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go playBuffer(wg, time.Duration(rand.Float32() * float32(time.Second)), fileBuffers[rand.Intn(len(fileBuffers) - 1)])
	}

	wg.Wait()
}