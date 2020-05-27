package main

import (
	"log"
	"time"
)

func main() {
	log.Println("container host is up")

	for {
		time.Sleep(time.Second * 1)
	}
}
