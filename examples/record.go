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
	list := streamly.ParseRecord("http://amtv1.livestreamingcdn.com/am1abr/tracks-v2a1/index-1531414800-7800.m3u8", 10*time.Minute)
	streamly.WriteAll(f, "http://amtv1.livestreamingcdn.com/am1abr/tracks-v2a1/", list)
}
