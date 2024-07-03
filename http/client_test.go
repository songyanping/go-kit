package http

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestHttpClient(t *testing.T) {
	cxt := context.Background()
	c := NewClient()

	bodyByte, _ := readJSONFile("car.json")
	body := string(bodyByte)
	c.RequestWithBody(cxt, "https://open.feishu.cn/open-apis/bot/v2/hook/4c686ae1-77c7-4b90-9510-52944d1fa476", "POST", body)
}

// 定义从JSON文件中读取数据的函数，使用 os.ReadFile 替换 ioutil.ReadFile
func readJSONFile(filePath string) ([]byte, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %v", err)
	}
	return fileContent, nil
}
