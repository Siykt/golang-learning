package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func copyAndCheck(dest string, src io.Reader, job string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	buf := make([]byte, 1*1024*1024)
	n, err := io.CopyBuffer(out, src, buf)
	if err != nil {
		return err
	}

	if n < 1*1024*1024 {
		newPath := filepath.Join("G:/HG/", job+".xml")
		err = os.Rename(dest, newPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadFile(url string, dest string, bar *mpb.Bar, job string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	total := resp.ContentLength

	if total < 1*1024 { // Skip files smaller than 1KB
		return nil
	}

	proxyReader := bar.ProxyReader(resp.Body)

	err = copyAndCheck(dest, proxyReader, job)
	if err != nil {
		return err
	}

	bar.SetTotal(total, true)
	return nil
}

func worker(id int, jobs <-chan string, p *mpb.Progress, wg *sync.WaitGroup) {
	defer wg.Done()

	bar := p.AddBar(0, // Initialize with dummy total
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("Worker %02d ", id), decor.WC{W: 10, C: decor.DidentRight}),
			decor.EwmaSpeed(decor.UnitKiB, "%6.2f KiB/s ", 60, decor.WC{W: 15}),
			decor.CountersKibiByte("%.2f / %.2f", decor.WC{W: 20}),
		),
		mpb.AppendDecorators(
			decor.Percentage(decor.WC{W: 5}),
		),
	)

	for job := range jobs {
		dest := filepath.Join("G:/HG/", job+".rar")
		url := "https://mifeng.oss-cn-hangzhou.aliyuncs.com/" + job + "/GMAE.rar"

		bar.SetTotal(0, false) // Reset bar for new job
		bar.SetCurrent(0)      // Reset current progress
		err := downloadFile(url, dest, bar, job)
		if err != nil {
			continue
		}
	}
}

func main() {
	const numWorkers = 10
	jobs := make(chan string, numWorkers)
	var wg sync.WaitGroup

	p := mpb.New(mpb.WithWaitGroup(&wg))

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, p, &wg)
	}

	go func() {
		for id := 23117; id <= 30000; id++ {
			jobs <- strconv.Itoa(id)
		}
		close(jobs)
	}()

	wg.Wait() // Wait for all workers to finish

	p.Wait() // Wait for all progress bars to complete

	fmt.Println("All jobs completed")
}
