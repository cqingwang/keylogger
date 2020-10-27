package main

import (
	"fmt"
	"github.com/cqingwang/usb_keyboard/core"
	"os"
	"time"
)

func main() {
	fmt.Println("App starting....")
	keyboard.Watch(func(self *keyboard.KeyStor) {
		fmt.Println("time:", self.Get()[0].UnixTime(), "code:", self.ToString())
	})

	go func() {
		for {
			time.Sleep(time.Second * 3)
			fmt.Println("keyboard injected:", keyboard.HasListening())
		}
	}()

	time.Sleep(time.Second * 30)
	fmt.Println("force exit")
	os.Exit(0)
}
