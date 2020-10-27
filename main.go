package main

import (
	"fmt"
	"time"
	"usb_keyboard/usage"
)

func main() {
	fmt.Println("App starting....")
	usage.Watch(func(self *usage.KeyStor) {
		fmt.Println("time:", self.Get()[0].UnixTime(), "code:", self.ToString())
	})

	for {
		time.Sleep(time.Second * 3)
		fmt.Println("keyboard listen:", usage.HasListening())
		if !usage.HasListening() {
			fmt.Println("keyboard maybe reject")
			break
		}
	}
}
