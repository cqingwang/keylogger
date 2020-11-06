package usage

import (
	"fmt"
	"github.com/cqingwang/usb_keyboard/keyboard"
	"sync"
	"time"
)

func Watch(complete func(self *KeyStor)) {
	go func() {
		for {
			devices := keyboard.FindAllKeyboardDevices()
			if len(devices) <= 0 {
				time.Sleep(time.Second * 3)
				continue
			}

			for _, dev := range devices {
				val, ok := listenDevices.Load(dev)
				if ok && val.(bool) {
					continue
				}
				fmt.Println("keyboard.Bind:", dev)
				DeviceBind(dev, complete)
			}
			time.Sleep(time.Second * 3)
		}
	}()

}

var listenDevices = sync.Map{}

func DeviceBind(devPath string, listener func(*KeyStor)) {
	go func() {
		dev, err := keyboard.New(devPath) ///dev/input/event14
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dev.Close()
		events := dev.Read()
		listenDevices.Store(devPath, true)
		keyStore := &KeyStor{complete: listener}
		for e := range events {
			if e.Code == keyboard.SHUTDOWN {
				fmt.Println("keyboard.UnBind:", devPath)
				listenDevices.Store(devPath, false)
				break
			}
			handleKeyEvent(e, keyStore)
		}
	}()
}

func handleKeyEvent(e keyboard.InputEvent, keystore *KeyStor) {
	switch e.Type {
	case keyboard.EvKey:
		//logKeyPress(e)
		onKeyRelease(e, keystore)
		break
	}

}
