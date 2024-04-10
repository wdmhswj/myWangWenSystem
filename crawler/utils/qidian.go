package utils

import (
	"crawler/structs"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

func SaveAsJson(filename string, entity structs.RankingList_qidian) {
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

func LoadJsonAsStruct(filename string) structs.RankingList_qidian {
	// fmt.Println(filename)
	if _, err := os.Stat(filename); err == nil {
		// 打开 JSON 文件
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("打开 JSON 文件失败:", err)
			return structs.RankingList_qidian{}
		}
		defer file.Close()

		// 读取 JSON 数据
		jsonData, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("读取 JSON 数据失败:", err)
			return structs.RankingList_qidian{}
		}

		// 解析 JSON 数据到结构体实例
		var entity structs.RankingList_qidian

		err = json.Unmarshal(jsonData, &entity)
		if err != nil {
			fmt.Println("反序列化 JSON 失败:", err)
			return structs.RankingList_qidian{}
		}

		fmt.Println("反序列化成功")
		return entity

	} else {
		fmt.Println(filename + "文件不存在！")
		return structs.RankingList_qidian{}
	}
}
