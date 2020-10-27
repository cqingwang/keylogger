package main

import (
	"fmt"
	"os"
	"time"
	"usb_keyboard/usage"
)

func main() {
	fmt.Println("App starting....")
	usage.Watch(func(self *usage.KeyStor) {
		fmt.Println("time:", self.Get()[0].UnixTime(), "code:", self.ToString())
	})

	go func() {
		for {
			time.Sleep(time.Second * 3)
			fmt.Println("keyboard injected:", usage.HasListening())
		}
	}()

	time.Sleep(time.Second * 30)
	fmt.Println("force exit")
	os.Exit(0)
}
