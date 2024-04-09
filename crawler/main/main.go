package main

import (
	"crawler/qidian"
	"crawler/utils"
	"crawler/yousuu"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	rankList, _ := qidian.GetTodayPopularNovels()
	utils.SaveAsJson(rankList.Name, rankList)
}

func test_yousuu() {
	test, err := yousuu.GetTodayPopularNovels()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 将结构体实例序列化为 JSON 格式
	jsonData, err := json.MarshalIndent(test, "", "    ")
	if err != nil {
		fmt.Println("序列化 JSON 失败:", err)
		return
	}

	// 将 JSON 数据写入本地文件
	err = os.WriteFile("ranking_list.json", jsonData, 0644)
	if err != nil {
		fmt.Println("写入 JSON 文件失败:", err)
		return
	}

	fmt.Println("JSON 文件保存成功！")
}
