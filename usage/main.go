package main

import (
	"fmt"
	"github.com/cqingwang/usb_keyboard"
	"os"
	"time"
)

func main() {
	fmt.Println("App starting....")
	usb_keyboard.Watch(func(self *usb_keyboard.KeyStor) {
		fmt.Println("time:", self.Get()[0].UnixTime(), "code:", self.ToString())
	})

	go func() {
		for {
			time.Sleep(time.Second * 3)
			fmt.Println("keyboard injected:", usb_keyboard.HasListening())
		}
	}()

	time.Sleep(time.Second * 30)
	fmt.Println("force exit")
	os.Exit(0)
}
