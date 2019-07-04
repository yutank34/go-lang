package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	f, err := os.Open(os.Args[1])
	urls := getUrls(f)
	for _, url := range urls {
		go fetch(url, ch) //start goroutin
	}
	file, err := os.OpenFile(`./file`, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {

	}
	for range urls {
		// fmt.Println(<-ch) // receive from chanel
		file.Write(([]byte)(<-ch))
	}
	file.Close()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send	 to chanel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

func getUrls(f *os.File) []string {
	var v []string
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := strings.Split(input.Text(), ",")
		url := "https://" + s[1]
		v = append(v, url)
	}
	return v
}
