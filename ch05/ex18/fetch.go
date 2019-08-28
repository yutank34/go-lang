package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println(fetch(os.Args[1]))
}

// urlをダウンロードしてローカルファイルの名前と長さを返す
func fetch(url string) (fileName string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	fileName = path.Base(resp.Request.URL.Path)
	if fileName == "/" {
		fileName = "index.html"
	}
	f, err := os.Create(fileName)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return fileName, n, err
}
