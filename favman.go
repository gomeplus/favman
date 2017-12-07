package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path"
)

const (
	TIMEOUT   int    = 10
	MODE      int    = 1 // 0:file, 1:text
	FILE_PATH string = "favorites"
	TEXT_PATH string = "favorites.txt"
)

func main() {
	fmt.Printf("TIMEOUT: %d\n", TIMEOUT)
	fmt.Printf("MODE: %d\n", MODE)

	switch MODE {

	case 0:
		break
	case 1:
		loadTextFavorites()
		break
	}

}

func loadTextFavorites() {
	file, err := os.Open(path.Join(path.Dir(os.Args[0]), TEXT_PATH))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var line []byte
	err = nil

	for err == nil {
		line, _, err = reader.ReadLine()

		if len(line) != 0 {
			fmt.Printf("%s : ", string(line))

			req, _ := http.NewRequest("GET", string(line), nil)
			req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")

			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
			//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
			req.Header.Set("Accept-Language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7")
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Connection", "keep-alive")
			req.Header.Set("Pragma", "no-cache")
			req.Header.Set("Upgrade-Insecure-Requests", "1")

			client := &http.Client{}
			resp, _ := client.Do(req)

			fmt.Printf("%d", resp.StatusCode)
			if resp.StatusCode != http.StatusOK {
				fmt.Print("\n")
			}

			fmt.Print("\n")
		}
	}
}
