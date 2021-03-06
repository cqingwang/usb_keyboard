package main

import (
	"fmt"
	"github.com/cqingwang/usb_keyboard/usage"
	"os"
	"time"
)

func main() {
	fmt.Println("App starting....")
	usage.Watch(func(self *usage.KeyStor) {
		fmt.Println("time:", self.Get()[0].UnixTime(), "code:", self.ToString())
	})

	time.Sleep(time.Second * 30)
	fmt.Println("force exit")
	os.Exit(0)
}
