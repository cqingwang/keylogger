package usage

import (
	"fmt"
	"usb_keyboard/keyboard"
)

func logKeyRelease(e *keyboard.InputEvent, keystore *KeyStore) {
	if e.KeyRelease() {
		logEvent("keyRelease", e)
		keystore.Append(e)
	}
}

func logKeyPress(e *keyboard.InputEvent) {
	if e.KeyPress() {
		logEvent("keyPress", e)
	}
}

func logEvent(title string, e *keyboard.InputEvent) {
	fmt.Println("")
	fmt.Println("time:", e.UnixTime(), ", type:", e.Type, ", code:", e.Code)
	fmt.Println(title, e.KeyString())
}
