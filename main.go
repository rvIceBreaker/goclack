package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/rvIceBreaker/goclack/audio"
	"github.com/rvIceBreaker/goclack/config"
	"github.com/rvIceBreaker/goclack/eventhandler"
	"github.com/rvIceBreaker/goclack/eventhook"
	"github.com/rvIceBreaker/goclack/keycodes"
	"github.com/rvIceBreaker/goclack/keyset"
	"golang.org/x/exp/slices"
)

func CalcKeyCombinationStr(keys []*keycodes.VKey) (string) {
	str := ""
	for i,key := range keys {
		if i <= 0 {
			str = key.GetSimpleName()
		} else {
			str = fmt.Sprintf("%s+%s",str, key.GetSimpleName())
		}
	}

	return str
}

var CurrentKeySet *keyset.KeySet

func main() {
	conf,err := config.LoadConfig()
	if err != nil {
		conf = config.NewAppConfig()
		conf.SaveConfig()
	}

	audio := audio.NewAudio(44100)
	audio.OutputVolume = conf.App.GlobalVolume
	defer audio.Teardown()

	CurrentKeySet,err = keyset.LoadKeySet(conf.App.DefaultSet, audio)
	if err != nil {
		panic(err)
	}

	//It seems that win32 message pumps need to happen on the same thread as the hook
	//Doing otherwise seems to break the message pump handling and we don't get key events
	go func() {
		keys := []*keycodes.VKey{}
		eventHandler := eventhook.NewEventHook()
		defer eventHandler.Teardown()
		eventHandler.RegisterKeyboardCallback(func(ev *eventhandler.KeyEvent) {
			var sounds []string
			switch ev.Type {
				case eventhandler.KEYEVENT_DOWN:
					//There might be a better way to do this, but this provides a do-once trap
					if !slices.Contains(keys, ev.Key) {
						keys = append(keys, ev.Key)
					} else {
						return
					}
					sounds = CurrentKeySet.Config.GetKeyEventSounds(CalcKeyCombinationStr(keys), ev.Key.GetSimpleName(), eventhandler.KEYEVENT_DOWN)
				case eventhandler.KEYEVENT_UP:
					sounds = CurrentKeySet.Config.GetKeyEventSounds(CalcKeyCombinationStr(keys), ev.Key.GetSimpleName(), eventhandler.KEYEVENT_UP)
					keys = slices.DeleteFunc(keys, func(vk *keycodes.VKey) bool {
						return vk == ev.Key
					})
			}

			if len(sounds) > 0 {
				err := audio.PlaySound(sounds[rand.Intn(len(sounds))])
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		})

		eventHandler.RegisterKeyboardHook()
		eventHandler.ListenForEvents()
	}()

	sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <-sigs
		fmt.Println("Received ", sig.String(), ", exiting...")
        done <- true
    }()

    <-done
}