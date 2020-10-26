package usage

import "usb_keyboard/keyboard"

type KeyStore struct {
	keys []*keyboard.InputEvent
	done func(keys []*keyboard.InputEvent)
}

func (i *KeyStore) Append(key *keyboard.InputEvent) *KeyStore {
	i.keys = append(i.keys, key)
	if key.Code == keyboard.Key_ENTER {
		i.done(i.pop())
	}
	return i
}

func (i *KeyStore) Read() []*keyboard.InputEvent {
	return i.keys
}

func (i *KeyStore) pop() []*keyboard.InputEvent {
	i.keys = []*keyboard.InputEvent{}
	return i.keys
}
