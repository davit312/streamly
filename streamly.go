package streamly

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"strings"
)

func Parse(addr string, f io.Writer) []string {
	data, err := http.Get(addr)
	if err != nil {
		log.Print(err)
		return nil
	}

	reader := bufio.NewReader(data.Body)
	for {
		r, _, e := reader.ReadLine()
		if e != nil {
			log.Print(e)
			return nil
		}

		s := string(r)

		if !strings.ContainsRune(s, '#') {
			data, er := http.Get("http://amtv1.livestreamingcdn.com/am1abr/tracks-v2a1/" + s)
			if er != nil {
				continue
			}
			io.Copy(f, data.Body)
		}
	}

	return nil
}
