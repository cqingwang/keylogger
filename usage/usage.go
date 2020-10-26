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
		//onPress(e)
		onRelease(e)
		break
	}
}

func onRelease(e keyboard.InputEvent) {
	if e.KeyRelease() {
		logEvent("keyRelease", e)
	}
}

func onPress(e keyboard.InputEvent) {
	if e.KeyPress() {
		logEvent("keyPress", e)
	}
}

func logEvent(title string, e keyboard.InputEvent) {
	fmt.Println("")
	fmt.Println("time:", e.UnixTime(), ", type:", e.Type, ", code:", e.Code)
	fmt.Println(title, e.KeyString())
}
