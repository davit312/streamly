package main

import (
	"log"
	"os"
	"streamly"
	"time"
)

func main() {

	var last [8]string
	f, err := os.Create("out")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		next := time.After(15 * time.Second)
		list := streamly.Parse("http://amtv1.livestreamingcdn.com/am1abr/tracks-v2a1/index.m3u8", &last)
		streamly.Write(f, "http://amtv1.livestreamingcdn.com/am1abr/tracks-v2a1/", list)
		<-next
	}
}