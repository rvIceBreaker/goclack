# goclack
A silly input noisemaker written in Go, using [faiface/beep](https://github.com/faiface/beep) for audio output

Written with `Go v1.20.5`

Goclack hooks into low level input events and uses them to trigger sounds. Current only Windows events are implemented, but Linux and MacOS support is planned.

## Using Goclack

Upon starting up, Goclack will create a settings.yaml next to the executable if one does not exist, which will be structed like so

```yaml
app:
  global_volume: 100.0              #The global volume modifier for sounds played, float value from 0.0 to 100.0
  default_set: "./path/to/keyset"   #A path to a default keyset, either a .yaml config file or a .zip pack file
```

## Keysets

Goclack uses `KeySets` to define what sounds should play. KeySet configuration files are yaml files defined like so

```yaml
content_root: /aud/keys/        #A path relative to the config file (in a .zip or otherwise) where sound files are located
events:
  keys:                         #'keys' contains configuration for keyboard events
    key_down:                   #'key_down' triggers when a key is pressed
      "home":                   #You can define raw key names to define groups of sounds for that key
        - SoundName1
      "ctrl+c":                 #You can also define groups of sounds for key combinations
        - SoundName1
      default:                  #'default' is a group of sounds that will play if no specific key or combo group is defined
        - SoundName1
        - SoundName2
        - SoundName3
    key_up:                     #'key_up' triggers when a key is released
      default:
        - SoundName1
        - SoundName2
        - SoundName3            #Sound names should be the name of the audio file without the extension
```

When an input event matches a group, a sound from that group is selected and played at random

#### Keyset Packs

KeySets can be packed into .zip files to easily bundle together the assets and configuration for a KeySet.

KeySet packs do not have a specific structure other than a `keyset.yaml` must exist in the root directory.

## Future plans

The following features are planned

* Cross-platform support
* * Linux and MacOS input events
* A TUI graphical interface
* * Possibily using [bubbletea](https://github.com/charmbracelet/bubbletea)
* Expanded input event handling
* * Mouse input events
* * * Movement
* * * Buttons
* * * Scrolling

## Known issues

Due to the lack of support in [faiface/beep](https://github.com/faiface/beep), the output audio device can not be specified.

