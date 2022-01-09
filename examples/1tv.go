package main

import (
	"os"
	"streamly"
	"time"
	"flag"
	"log"
)

func main() {

	playlist := flag.String("list", "", "playlist url")
	startAt := flag.String("start", "0", "start record at minute")
	prefix := flag.String("prefix", "http://amtv1.livestreamingcdn.com/am1abr/tracks-v1a1/", "source to download video chunks")
	fileName := flag.String("out", "out.ts", "output file address")
    flag.Parse()

	if(len(*playlist) < 1){
		log.Fatal("Error: no playlist.")
	}


	skip, err := time.ParseDuration(*startAt)
	if err != nil {
		log.Fatal(err)
	}

	
	f, e := os.Create(*fileName)
	if e != nil {
		log.Fatal(e)
	}
	list := streamly.ParseRecord(*playlist, skip)
	streamly.WriteAll(f, *prefix, list)
}
