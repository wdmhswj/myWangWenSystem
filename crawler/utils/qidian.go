package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
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

// 替换
func ReplacePlaceholer(placeHodler string, target string, replacer string) string {

	// 定义替换函数
	replacer_func := func(match string) string {
		switch match {
		case placeHodler:
			return replacer
		default:
			return match
		}
	}

	// 定义正则表达式
	re := regexp.MustCompile(`\{([^}]+)\}`)

	// 使用正则表达式进行替换
	output := re.ReplaceAllStringFunc(target, replacer_func)

	return output
}
