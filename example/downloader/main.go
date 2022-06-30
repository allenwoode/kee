package main

import (
	"eeyoo.com/downloader/downloader"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	uri := os.Args[1]

	n := runtime.NumCPU()

	t1 := time.Now()
	downloader.NewDownloader(uri).Download(n)
	elapsed := time.Since(t1)
	fmt.Println("total elapsed:", elapsed)
}
