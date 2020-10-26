package main

import (
	"fmt"
	"usb_keyboard/usage"
)

func main() {
	fmt.Println("App starting....")
	devices, err := usage.DevicesFind()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	usage.DeviceBind(err, devices[0])
}
