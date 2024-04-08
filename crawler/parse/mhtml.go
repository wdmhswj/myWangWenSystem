package parse

import (
	"fmt"
	"os"
)

func Parse_mhtml(file_path string) {
	// 读取 .mhtml 文件内容
	content, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	// 将字节切片转换为字符串
	mhtmlContent := string(content)

	// 打印 .mhtml 文件内容
	fmt.Println(mhtmlContent)
}
