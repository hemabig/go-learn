package main

import (
	"fmt"
	"net/http"

	// "net/url"
	"crypto/tls"
	"log"
	"os"
	// "time"
)

func main() {
	// 创建自定义 Client
	sslKeyLogfile := "log.txt"
	w, err := os.OpenFile(sslKeyLogfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("Could not create keylogger: ", err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				KeyLogWriter:       w,
				InsecureSkipVerify: true, // test server certificate is not trusted.
				MinVersion:         tls.VersionTLS12,
				MaxVersion:         tls.VersionTLS13,
			},
		},
	}

	// 创建请求
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		fmt.Println("创建请求错误:", err)
		return
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求错误:", err)
		return
	}
	defer resp.Body.Close()
	// 处理响应
	fmt.Println("响应状态码:", resp.StatusCode)
	// 其他操作...
}
