package usage

import (
	"usb_keyboard/keyboard"
)

type KeyStore struct {
	keys []keyboard.InputEvent
	done func(self *KeyStore)
}

func (i *KeyStore) ToString() string {
	rst := ""
	for _, e := range i.keys {
		rst += e.KeyString()
	}
	return rst
}

func (i *KeyStore) Append(key keyboard.InputEvent) *KeyStore {
	if key.Code == keyboard.Key_ENTER {
		i.done(i)
		i.pop()
		return i
	}
	i.keys = append(i.keys, key)
	return i
}

func (i *KeyStore) Read() []keyboard.InputEvent {
	return i.keys
}

func (i *KeyStore) pop() []keyboard.InputEvent {
	i.keys = []keyboard.InputEvent{}
	return i.keys
}
