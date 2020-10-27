package usage

import (
	"errors"
	"fmt"
	"time"
	"usb_keyboard/keyboard"
)

func Watch(complete func(self *KeyStor)) {
	if HasWatching() {
		return
	}
	setWatching(true)
	go func() {
		for {
			if !HasWatching() {
				fmt.Println("Break keyboard watching")
				break
			}

			if HasListening() {
				time.Sleep(time.Second * 3)
				continue
			}

			devices := keyboard.FindAllKeyboardDevices()
			if len(devices) <= 0 {
				time.Sleep(time.Second * 3)
				continue
			}

			DeviceBind(devices[0], complete)
			time.Sleep(time.Second * 3)
		}
	}()

}
func DevicesFind() ([]string, error) {
	devices := keyboard.FindAllKeyboardDevices()
	if len(devices) <= 0 {
		return nil, errors.New("not one keyboard")
	}
	fmt.Println("devices:", devices)
	return devices, nil
}

func DeviceBind(devPath string, listener func(self *KeyStor)) {
	go func() {
		dev, err := keyboard.New(devPath) ///dev/input/event14
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dev.Close()
		events := dev.Read()
		setListening(true)
		keyStore := &KeyStor{complete: listener}
		for e := range events {
			if e.Code == keyboard.SHUTDOWN {
				setListening(false)
				break
			}
			handleKeyEvent(e, keyStore)
		}
	}()
}

func handleKeyEvent(e keyboard.InputEvent, keystore *KeyStor) {
	switch e.Type {
	case keyboard.EvKey:
		//logKeyPress(e)
		onKeyRelease(e, keystore)
		break
	}

}
