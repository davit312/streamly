package streamly

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func getM3u(url string) *bufio.Reader {
	conn, err := http.Get(url)
	if err != nil {
		return nil
	}
	return bufio.NewReader(conn.Body)
}

func Parse(addr string, last *[8]string) []string {

	i := 0
	l := 0
	isNew := false

	reader := getM3u(addr)
	result := make([]string, 8)

	for i < 8 {
	begin:
		r, _, e := reader.ReadLine()
		if e != nil {
			break
		}
		s := string(r)
		if !strings.ContainsRune(s, '#') {
			if !isNew {
				isNew = true
				for _, item := range last {
					if item == s {
						isNew = false
						goto begin
					}
				}
			}

			result[i] = s
			i++

			if l == 8 {
				l = 0
			}
			last[l] = s
			l++
		}
	}
	return result[:i]
}

func ParseRecord(addr string, start time.Duration) []string {
	reader := getM3u(addr)
	result := make([]string, 0)
	i := 0
	now := time.Duration(0 * time.Second)
	skip := false

	for i < 10500 {
		r, _, e := reader.ReadLine()
		if e != nil {
			break
		}
		s := string(r)

		if start != 0 && strings.Contains(s, "#EXTINF") {
			skip = false

			var t float32
			fmt.Sscanf(s, "#EXTINF:%f,", &t)
			now += time.Duration(t) * time.Second

			if now < start {
				skip = true
			}

		} else if !strings.ContainsRune(s, '#') && !skip {
			result = append(result, s)
			i++
		}
	}
	return result[:i]
}
