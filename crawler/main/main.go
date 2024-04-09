package main

import (
	"crawler/qidian"
	"crawler/structs"
)

func main() {
	// rankList, _ := qidian.GetTodayPopularNovels(
	// 	"https://www.qidian.com/rank/readindex/",
	// 	"https://www.qidian.com/rank/readindex/page{index}/",
	// 	"www.qidian.com",
	// 	"{index}",
	// 	5,
	// )

	// utils.SaveAsJson(rankList.Name+rankList.Time.Format("20060102150405"), rankList)

	// rankList2, _ := qidian.GetTodayPopularNovels(
	// 	"https://www.qidian.com/rank/newfans/",
	// 	"https://www.qidian.com/rank/newfans/page{index}/",
	// 	"www.qidian.com",
	// 	"{index}",
	// 	5,
	// )

	// utils.SaveAsJson(rankList2.Name+rankList2.Time.Format("20060102150405"), rankList2)

	// var rankNames []string = []string{
	// 	"readindex",
	// 	"newfans",
	// 	"yuepiao",
	// 	"hotsales",
	// 	"newfans",
	// 	"recom",
	// 	"collect",
	// 	"vipup",
	// 	"vipcollect",
	// }

	placeHolder := "{index}"

	// qidian.GetViableRanks(structs.Ranks, placeHolder)
	qidian.GetViableRankSpecifyClasses(structs.Ranks, structs.TagsPageName, placeHolder)
}

// func test_yousuu() {
// 	test, err := yousuu.GetTodayPopularNovels()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// 将结构体实例序列化为 JSON 格式
// 	jsonData, err := json.MarshalIndent(test, "", "    ")
// 	if err != nil {
// 		fmt.Println("序列化 JSON 失败:", err)
// 		return
// 	}

// 	// 将 JSON 数据写入本地文件
// 	err = os.WriteFile("ranking_list.json", jsonData, 0644)
// 	if err != nil {
// 		fmt.Println("写入 JSON 文件失败:", err)
// 		return
// 	}

// 	fmt.Println("JSON 文件保存成功！")
// }
