package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func downloadFile(url string, dest string, p *mpb.Progress) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	total := resp.ContentLength

	if total < 1*1024 { // Skip files smaller than 1KB
		_, err = io.Copy(out, resp.Body) // Still need to download to handle renaming
		return err
	}

	parts := strings.Split(url, "/")
	desiredPart := parts[len(parts)-2] + ".rar"

	bar := p.AddBar(total,
		mpb.PrependDecorators(
			decor.Name(desiredPart+" ", decor.WC{W: len(desiredPart) + 1, C: decor.DidentRight}),
			decor.CountersKibiByte("%.2f / %.2f"),
		),
		mpb.AppendDecorators(
			decor.Percentage(decor.WC{W: 5}),
		),
	)

	proxyReader := bar.ProxyReader(resp.Body)

	_, err = io.Copy(out, proxyReader)
	if err != nil {
		return err
	}

	bar.SetTotal(total, true)
	return nil
}

func worker(jobs <-chan string, p *mpb.Progress, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		dest := filepath.Join("G:/HG/", job+".rar")
		url := "https://mifeng.oss-cn-hangzhou.aliyuncs.com/" + job + "/GMAE.rar"

		err := downloadFile(url, dest, p)
		if err != nil {
			continue
		}

		fileInfo, err := os.Stat(dest)
		if err != nil {
			continue
		}

		if fileInfo.Size() < 1*1024*1024 {
			newPath := filepath.Join("G:/HG/", job+".xml")
			err := os.Rename(dest, newPath)
			if err != nil {
				continue
			}
		}
	}
}

func main() {
	const numWorkers = 2
	jobs := make(chan string, numWorkers)
	var wg sync.WaitGroup

	p := mpb.New(mpb.WithWaitGroup(&wg))

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(jobs, p, &wg)
	}

	go func() {
		for id := 28329; id < 30000; id++ {
			jobs <- strconv.Itoa(id)
		}
		close(jobs)
	}()

	wg.Wait() // Wait for all workers to finish

	p.Wait() // Wait for all progress bars to complete

	fmt.Println("All jobs completed")
}
