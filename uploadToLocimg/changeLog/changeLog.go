package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 打开 JSON 文件
	filePath := `C:\www\golang-learning\upload_results.json`
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer jsonFile.Close()

	// 读取文件内容
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// 定义一个 map 来存储解析的 JSON 数据
	var data map[string]string
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// 定义一个新的 map 来存储转换后的数据
	transformedData := make(map[string][]string)

	// 遍历原始数据，构造新的 URL
	for key, _ := range data {
		// 分割 key，获取文件名和类别编号
		parts := strings.Split(key, "\\")
		category := parts[0] // 类别编号（如 "1", "2"）
		filename := parts[1] // 文件名（如 "IPAdapter_06391_.png"）

		// 构造新的 URL，包含类别编号
		newURL := fmt.Sprintf("https://miss-ribbit.pages.dev/%s/%s", category, filename)

		// 将新的 URL 添加到新的 map 中
		transformedData[category] = append(transformedData[category], newURL)
	}

	// 将转换后的数据转成 JSON 输出
	transformedJSON, err := json.MarshalIndent(transformedData, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling transformed JSON:", err)
		return
	}

	// 打印转换后的 JSON
	fmt.Println(string(transformedJSON))
}
