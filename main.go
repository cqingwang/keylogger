package main

import (
	"fmt"
	"keylogger/keylogger"
)

func main() {
	fmt.Println("start")
	devices := keylogger.FindAllKeyboardDevices()
	fmt.Println("devices:", devices)
	// init keylogger with keyboard
	k, err := keylogger.New("/dev/input/event14")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:

			// if the state of key is pressed
			if e.KeyPress() {
				fmt.Println("[event] press key ", e.Time, e.Type, e.Code, e.KeyString())
			}

			// if the state of key is released
			if e.KeyRelease() {
				fmt.Println("[event] release key ", e.Time, e.Type, e.Code, e.KeyString())
			}
			break
		}
	}
}
