package config

import (
	"strings"

	"github.com/rvIceBreaker/goclack/eventhandler"
)

type KeySetConfig struct {
	ContentPath string `yaml:"content_root"`
	Events      struct {
		Keys struct {
			KeyDown map[string][]string `yaml:"key_down"`
			KeyUp   map[string][]string `yaml:"key_up"`
		} `yaml:"keys"`
	} `yaml:"events"`
}

func (c *KeySetConfig) GetKeyEventSounds(keycombo string, currentKey string, eventType eventhandler.KeyEventType) []string {
	keycombo = strings.ToLower(keycombo)
	currentKey = strings.ToLower(currentKey)

	//Check for keycombos first, then the new key from the event, then defaults if they exist
	switch eventType {
		case eventhandler.KEYEVENT_DOWN:
			if len(c.Events.Keys.KeyDown) > 0 {
				sounds,ok := c.Events.Keys.KeyDown[keycombo]
				if ok {
					return sounds
				}

				sounds,ok = c.Events.Keys.KeyDown[currentKey]
				if ok {
					return sounds
				}

				sounds,ok = c.Events.Keys.KeyDown["default"]
				if ok {
					return sounds
				}

				return nil
			}
		case eventhandler.KEYEVENT_UP:
			if len(c.Events.Keys.KeyUp) > 0 {
				sounds,ok := c.Events.Keys.KeyUp[keycombo]
				if ok {
					return sounds
				}

				sounds,ok = c.Events.Keys.KeyUp[currentKey]
				if ok {
					return sounds
				}

				sounds,ok = c.Events.Keys.KeyUp["default"]
				if ok {
					return sounds
				}

				return nil
			}
	}

	return nil
}