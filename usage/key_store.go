package usage

import (
	"usb_keyboard/keyboard"
)

type KeyStore struct {
	keys     []keyboard.InputEvent
	complete func(self *KeyStore)
}

func (i *KeyStore) ToString() string {
	rst := ""
	for _, e := range i.keys {
		rst += e.KeyString()
	}
	return rst
}

func (i *KeyStore) Append(key keyboard.InputEvent) *KeyStore {
	if key.Code == keyboard.KEY_DOWN {
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

func done(i *KeyStore) {
	if len(i.keys) > 0 {
		i.complete(i)
	}
}

func (i *KeyStore) Read() []keyboard.InputEvent {
	return i.keys
}

func (i *KeyStore) popReset() []keyboard.InputEvent {
	i.keys = []keyboard.InputEvent{}
	return i.keys
}
