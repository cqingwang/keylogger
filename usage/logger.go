package usage

import (
	"fmt"
	"usb_keyboard/keyboard"
)

func onKeyRelease(e keyboard.InputEvent, keystore *KeyStore) {
	if !e.KeyRelease() {
		return
	}

	//logEvent(e)
	keystore.Append(e)
}

func logEvent(e keyboard.InputEvent) {
	fmt.Println("")
	fmt.Println("time:", e.UnixTime(), ", type:", e.Type, ", code:", e.Code)
	fmt.Println("keyRelease", e.KeyString())
}
