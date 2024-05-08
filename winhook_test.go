package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/rvIceBreaker/goclack/eventhandler"
	"github.com/rvIceBreaker/goclack/eventhook"
)

func TestKeyboardHook(t *testing.T) {
	fmt.Println("Starting up...")

	go func() {
		eventhook := eventhook.NewEventHook()
		eventhook.RegisterKeyboardCallback(func(ev *eventhandler.KeyEvent) {
			//fmt.Println(ev.KeyCode)
		})

		eventhook.RegisterKeyboardHook()
		eventhook.ListenForEvents()
	}()

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	select{}
}