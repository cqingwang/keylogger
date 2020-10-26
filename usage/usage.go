package usage

import (
	"errors"
	"fmt"
	"usb_keyboard/keyboard"
)

func DeviceBind(err error, device string) {
	k, err := keyboard.New(device) ///dev/input/event14
	if err != nil {
		fmt.Println(err)
		return
	}

	defer k.Close()
	events := k.Read()
	keyStore := &KeyStore{done: func(self *KeyStore) {
		fmt.Println("done:", self.ToString())
	}}
	for e := range events {
		handleKeyEvent(e, keyStore)
	}
}

func DevicesFind() ([]string, error) {
	devices := keyboard.FindAllKeyboardDevices()
	if len(devices) <= 0 {
		return nil, errors.New("not one keyboard")
	}
	fmt.Println("devices:", devices)
	return devices, nil
}

func handleKeyEvent(e keyboard.InputEvent, keystore *KeyStore) {
	switch e.Type {
	case keyboard.EvKey:
		//logKeyPress(e)
		logKeyRelease(e, keystore)
		break
	}

}
