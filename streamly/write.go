package streamly

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Write(file *os.File, prefix string, list []string) {

	length := len(list)
	data := make([]*io.ReadCloser, length)
	fatched := make(chan bool, length+1)

	fetch := func(index int, address string) {
		defer func() { fatched <- true }()
		conn, err := http.Get(address)
		if err != nil {
			println(777)
			log.Println(err)
			data[index] = nil
			return
		}
		data[index] = &conn.Body
	}

	for i, url := range list {
		fetch(i, prefix+url)
	}

	done := 0
	for _ = range fatched {
		done++
		if done == length {
			close(fatched)
			break
		}
	}

	for _, chunk := range data {
		io.Copy(file, *chunk)
		(*chunk).Close()
	}
}

func WriteAll(file *os.File, prefix string, list []string) {
	length := len(list)

	for i := 4; i < length; i += 4 {
		current := list[i-4 : i]
		Write(file, prefix, current)
	}
}
