package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// 上传的结果结构体
type UploadResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
	Data struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
		Size int64  `json:"size"`
		Mime string `json:"mime"`
		Sha1 string `json:"sha1"`
		Md5  string `json:"md5"`
	} `json:"data"`
}

func uploadImage(filePath string) (*UploadResponse, error) {
	url := "https://www.locimg.com/upload/upload.html"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile := os.Open(filePath)
	if errFile != nil {
		return nil, errFile
	}
	defer file.Close()

	part, errFile := writer.CreateFormFile("image", filepath.Base(filePath))
	if errFile != nil {
		return nil, errFile
	}
	_, errFile = io.Copy(part, file)
	if errFile != nil {
		return nil, errFile
	}

	// 其他上传字段
	_ = writer.WriteField("fileId", filepath.Base(filePath))
	_ = writer.WriteField("initialPreview", "[]")
	_ = writer.WriteField("initialPreviewConfig", "[]")
	_ = writer.WriteField("initialPreviewThumbTags", "[]")

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("dnt", "1")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("Cookie", "PHPSESSID=ecgaciq8etdgh4tn8t7g7nh0tf")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// 解析上传返回的 JSON
	var uploadResp UploadResponse
	err = json.Unmarshal(body, &uploadResp)
	if err != nil {
		return nil, fmt.Errorf("解析 JSON 出错: %v, 响应内容: %s", err, string(body))
	}

	return &uploadResp, nil
}

// worker 用于并发处理上传任务
func worker(id int, jobs <-chan string, results chan<- map[string]string, wg *sync.WaitGroup, rootDir string) {
	defer wg.Done()

	for filePath := range jobs {
		fmt.Printf("Worker %d 正在上传文件: %s\n", id, filePath)
		resp, err := uploadImage(filePath)
		if err != nil {
			fmt.Printf("Worker %d 上传失败: %v\n", id, err)
			continue
		}

		// 获取相对于 rootDir 的相对路径，包含子目录
		relPath, err := filepath.Rel(rootDir, filePath)
		if err != nil {
			fmt.Printf("Worker %d 获取相对路径失败: %v\n", id, err)
			continue
		}

		// 保存上传结果，使用相对路径作为键
		results <- map[string]string{relPath: resp.Data.URL}
	}
}

func main() {
	rootDir := "D:\\Downloads\\images\\111"
	jobs := make(chan string, 10)
	results := make(chan map[string]string, 10)
	var wg sync.WaitGroup

	// 启动多个 worker
	numWorkers := 5 // 可以根据需要调整 worker 数量
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg, rootDir)
	}

	// 遍历目录，将所有图片路径放入 jobs channel
	go func() {
		err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// 判断是否是文件，并且是图片文件
			if !info.IsDir() && (filepath.Ext(path) == ".png" || filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg") {
				jobs <- path // 将任务推入 jobs 通道
			}
			return nil
		})

		if err != nil {
			fmt.Println("遍历目录时出错:", err)
			close(jobs) // 关闭 jobs 通道
			return
		}
		close(jobs) // 所有任务已推入，关闭 jobs 通道
	}()

	// 等待所有 worker 完成任务
	go func() {
		wg.Wait()
		close(results) // 关闭结果通道
	}()

	// 收集结果
	finalResults := make(map[string]string)
	for result := range results {
		for key, value := range result {
			finalResults[key] = value
		}
	}

	// 保存结果为 JSON 文件
	outputFile := "upload_results.json"
	resultData, err := json.MarshalIndent(finalResults, "", "  ")
	if err != nil {
		fmt.Println("生成 JSON 出错:", err)
		return
	}

	err = ioutil.WriteFile(outputFile, resultData, 0644)
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		return
	}

	fmt.Printf("所有文件上传完成，结果已保存到 %s\n", outputFile)
}
