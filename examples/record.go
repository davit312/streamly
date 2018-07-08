package main

import (
	"os"
	"streamly"
	"time"
)

func main() {

	f, e := os.Create("out")
	if e != nil {
		return
	}
	list := streamly.ParseRecord("http://amtv1.livestreamingcdn.com/am1abr/tracks-v1a1/index-1531037700-3900.m3u8", 35*time.Minute)
	streamly.WriteAll(f, "http://amtv1.livestreamingcdn.com/am1abr/tracks-v1a1/", list)
}
