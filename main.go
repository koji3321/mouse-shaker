package main

import (
	"fmt"
	"time"

	"github.com/MarinX/keylogger"
	"github.com/go-vgo/robotgo"
)

var asd = true

func mouseMove() {
	for asd {
		robotgo.MoveRelative(5, 0)
		time.Sleep(100 * time.Millisecond)
		robotgo.MoveRelative(-5, 0)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go mouseMove()

	var dev []string = keylogger.FindAllKeyboardDevices()

	fmt.Println(dev)
	listener, err := keylogger.New(dev[0])
	if err != nil {
		fmt.Println("hata")
		return
	}

	for event := range listener.Read() {
		if event.Type == keylogger.EvKey {
			if event.KeyString() == "ESC" {
				fmt.Println("çıkış yap")
				asd = false
				return
			}
		}
	}
}
