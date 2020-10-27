package main

import (
	"fmt"
	"time"
	"usb_keyboard/usage"
)

func main() {
	fmt.Println("App starting....")
	devices, err := usage.DevicesFind()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	usage.DeviceBind(err, devices[0], func(self *usage.KeyStor) {
		fmt.Println("time:", self.Get()[0].UnixTime(), "code:", self.ToString())
	})

	for {
		time.Sleep(2000)
		fmt.Println("keyboard listen:", usage.HasListening())
	}
}
