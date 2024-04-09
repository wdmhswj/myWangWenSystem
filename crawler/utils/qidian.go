package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveAsJson(filename string, entity any) {
	filename += ".json"
	// fmt.Println(filename)
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("相同文件名称的JSON文件已存在！")
	} else {

		// 将结构体实例序列化为 JSON 格式
		jsonData, err := json.MarshalIndent(entity, "", "    ")
		if err != nil {
			fmt.Println("序列化 JSON 失败:", err)
			return
		}

		// 将 JSON 数据写入本地文件
		err = os.WriteFile(filename, jsonData, 0644)
		if err != nil {
			fmt.Println("写入 JSON 文件失败:", err)
			return
		}

		fmt.Println("JSON 文件保存成功！")
	}
}
