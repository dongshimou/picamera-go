package main

import (
	"flag"
	"log"

	"picamera-go/pilib/invoke/image"
	"picamera-go/pilib/invoke/video"
)

func imageOb() {
	ob := image.NewObserver()
	if err := ob.Start(); err != nil {
		panic(err)
	}
	if err := ob.Shoot(); err != nil {
		log.Println(err)
	}
}

func videoOb() {
	ob := video.NewObserver()
	if err := ob.Start(); err != nil {
		panic(err)
	}
	if err := ob.Shoot(); err != nil {
		log.Println(err)
	}
}

func main() {
	flag.Parse()
	imageOb()
}
