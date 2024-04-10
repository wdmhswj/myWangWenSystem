package analysis

import (
	"crawler/structs"
	"sort"
)

// 投票法
func GetComprehensiveRank(rankingLists []structs.RankingList_qidian) []string {
	record := make(map[string]int)

	for _, list := range rankingLists {
		// 识别榜单，确定榜单权重
		weight := structs.RankWeight[list.Name]
		entites := list.Entities
		for _, entity := range entites {
			value, ok := record[entity.Name] // 是否存在
			if ok {
				value += ((100 - entity.Rank) * weight)
				record[entity.Name] = value
			} else {
				record[entity.Name] = ((100 - entity.Rank) * weight)
			}
		}
	}

	return mysort(record)
}

func findBook(list structs.RankingList_qidian, bookName string) bool {
	entites := list.Entities
	for _, entity := range entites {
		if entity.Name == bookName {
			return true
		}
	}
	return false
}

func mysort(record map[string]int) []string {
	// 将 map 转换为切片
	keys := make([]string, 0, len(record))
	for key := range record {
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Slice(keys, func(i, j int) bool {
		return record[keys[i]] > record[keys[j]] // 降序
	})

	// // 输出排序后的结果
	// for _, key := range keys {
	// 	fmt.Printf("%s: %d\n", key, record[key])
	// }

	return keys
}
