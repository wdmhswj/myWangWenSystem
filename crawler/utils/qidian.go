package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveAsJson(filename string, entity any) {
	// 将结构体实例序列化为 JSON 格式
	jsonData, err := json.MarshalIndent(entity, "", "    ")
	if err != nil {
		fmt.Println("序列化 JSON 失败:", err)
		return
	}

	// 将 JSON 数据写入本地文件
	err = os.WriteFile(filename+".json", jsonData, 0644)
	if err != nil {
		fmt.Println("写入 JSON 文件失败:", err)
		return
	}

	fmt.Println("JSON 文件保存成功！")
}
