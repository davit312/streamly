package streamly

import (
	"bufio"
	"log"
	"net/http"
	"strings"
)

func Parse(addr string) []string {
	data, err := http.Get(addr)
	if err != nil {
		log.Print(err)
		return nil
	}

	result := make([]string, 8)
	reader := bufio.NewReader(data.Body)
	for i := 0; i < 8; {
		r, _, e := reader.ReadLine()
		if e != nil {
			log.Print(e)
			break
		}
		s := string(r)
		if !strings.ContainsRune(s, '#') {
			result[i] = s
			i++
		}
	}
	return result
}
