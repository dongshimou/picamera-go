package main

import (
	"log"

	"picamera-go/pilib/invoke/image"
)

func main() {
	ob := image.NewObserver()
	if err := ob.Start(); err != nil {
		panic(err)
	}
	if err := ob.Shoot(); err != nil {
		log.Println(err)
	}
}
