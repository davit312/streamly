package main

import (
	"log"
	"os"
	"github.com/davit312/streamly"
	"time"
)

func main() {

	var last [8]string
	f, err := os.Create("out")
	if err != nil {
		log.Println(err)
		return
	}

	next := time.Tick (15 * time.Second)
	for {
		list := streamly.Parse("http://amtv1.livestreamingcdn.com/am1abr/tracks-v2a1/index.m3u8", &last)
		streamly.Write(f, "http://amtv1.livestreamingcdn.com/am1abr/tracks-v2a1/", list)
		<-next
	}
}
