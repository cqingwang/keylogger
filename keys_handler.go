package usb_keyboard

import (
	"fmt"
)

type KeyStor struct {
	keys     []InputEvent
	complete func(self *KeyStor)
}

func onKeyRelease(e InputEvent, keystore *KeyStor) {
	if !e.KeyRelease() {
		return
	}
	//logEvent(e)
	keystore.Append(e)
}

func logEvent(e InputEvent) {
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

func (i *KeyStor) Append(key InputEvent) *KeyStor {
	if key.Code == KEY_DOWN_80 || key.Code == KEY_DOWN_108 {
		return i
	}

	if key.Code == KEY_ENTER {
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

func (i *KeyStor) Get() []InputEvent {
	return i.keys
}

func (i *KeyStor) popReset() []InputEvent {
	i.keys = []InputEvent{}
	return i.keys
}
