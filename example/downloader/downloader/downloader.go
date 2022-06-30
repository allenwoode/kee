package downloader

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Downloader struct {
	Uri                   string
	ConnectTimeout        int
	ReadTimeout           int
	ContentLength         int
	AlreadyDownloadLength int
}

func NewDownloader(uri string) *Downloader {
	return &Downloader{
		Uri: uri,
	}
}

func (d *Downloader) Download(num int) {
	fmt.Println("download thread num:", num)

	resp, err := http.Head(d.Uri)
	if err != nil {
		log.Fatal(err)
		return
	}

	if resp.StatusCode == http.StatusOK &&
		resp.Header.Get("Accept-Ranges") == "bytes" {
		log.Println("支持部分请求")
	}

	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(index int) {
			time.Sleep(1000 * time.Millisecond)
			wg.Done()
			fmt.Println(index, "done")
		}(i)
	}

	wg.Wait()
}
