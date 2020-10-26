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
	for e := range events {
		handleKeyEvent(e)
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

func handleKeyEvent(e keyboard.InputEvent) {
	switch e.Type {
	case keyboard.EvKey:
		fmt.Println(" - ")
		if e.KeyPress() {
			fmt.Println("time:", e.UnixTime(), ", type:", e.Type, ", code:", e.Code)
			fmt.Println("keyPress: ", e.KeyString())
		}
		if e.KeyRelease() {
			fmt.Println("time:", e.UnixTime(), ", type:", e.Type, ", code:", e.Code)
			fmt.Println("keyRelease:", e.KeyString())
		}
		fmt.Println("")
		break
	}
}
