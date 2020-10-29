package main

import (
	"flag"
	"github.com/gen2brain/beeep"
)

var (
	title = flag.String("title","Notification", "notification title")
	message = flag.String("message","", "notification message")
	appIcon = flag.String("appIcon","", "app icon")
)

func main() {
	flag.Parse()
	err := beeep.Alert(*title,*message,*appIcon)
	if err != nil {
		panic(err)
	}
}
