package usage

import (
	"fmt"
	"usb_keyboard/keyboard"
)

type KeyStor struct {
	keys     []keyboard.InputEvent
	complete func(self *KeyStor)
}

func onKeyRelease(e keyboard.InputEvent, keystore *KeyStor) {
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

func (i *KeyStor) ToString() string {
	rst := ""
	for _, e := range i.keys {
		rst += e.KeyString()
	}
	return rst
}

func (i *KeyStor) Append(key keyboard.InputEvent) *KeyStor {
	if key.Code == keyboard.KEY_DOWN_80 || key.Code == keyboard.KEY_DOWN_108 {
		return i
	}

	if key.Code == keyboard.KEY_ENTER {
		done(i)
		i.popReset()
		return i
	}

	i.keys = append(i.keys, key)
	return i
}

func done(i *KeyStor) {
	if len(i.keys) > 0 {
		i.complete(i)
	}
}

func (i *KeyStor) Get() []keyboard.InputEvent {
	return i.keys
}

func (i *KeyStor) popReset() []keyboard.InputEvent {
	i.keys = []keyboard.InputEvent{}
	return i.keys
}
