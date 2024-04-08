package utils

import (
	. "crawler/structs" // 直接使用包内导出物
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func fmt2Entity(s string) ListEntity_yousuu {
	attrs := strings.Split(s, " ")

	rank, err := strconv.Atoi(attrs[0])
	if err != nil {
		fmt.Println(err.Error())
		rank = -1
	}
	wordNum, err := wordNumTrans(attrs[2])
	if err != nil {
		fmt.Println(err.Error())
	}

	updateTime, err := timeTrans(attrs[3], attrs[4])
	if err != nil {
		fmt.Println(err.Error())
	}

	grade, gradersNum, err := gradeTrans(attrs[6])
	if err != nil {
		fmt.Println(err.Error())
	}

	res := ListEntity_yousuu{
		Rank:       rank,
		Name:       attrs[1],
		WordNum:    wordNum,
		UpdateTime: updateTime,
		State:      attrs[5],
		Grade:      grade,
		GradersNum: gradersNum,
		Tags:       attrs[10],
	}

	return res
}

func wordNumTrans(s string) (float32, error) {

	// 先将 "万字" 替换为空格，然后进行字符串拆分
	parts := strings.Split(strings.ReplaceAll(s, "万字", ""), " ")
	// 解析数字部分
	num, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return 0, err
	}

	return float32(num), nil
}

func timeTrans(num string, unit string) (int, error) {
	res, err := strconv.Atoi(num)
	if err != nil {
		return -1, nil
	}

	switch unit {
	case "个月":
		res *= 30
	case "年":
		res *= 365
	}
	return res, nil
}

func gradeTrans(s string) (float32, int, error) {
	// 使用正则表达式匹配数字和人数
	re := regexp.MustCompile(`([\d.]+)\((\d+)人\)`)

	// 使用正则表达式提取评分和人数
	matches := re.FindStringSubmatch(s)

	if len(matches) == 3 {
		score := matches[1]
		numGraders_str := matches[2]
		grade, err := strconv.ParseFloat(score, 32)
		if err != nil {
			return 0, 0, err
		}
		numGraders, err := strconv.Atoi(numGraders_str)
		if err != nil {
			return 0, 0, err
		}
		return float32(grade), numGraders, nil
	} else {
		return 0, 0, errors.New("未找到匹配的评分和评分人数")
	}

}

var Fmt2Entity = fmt2Entity
var WordNumTrans = wordNumTrans
var TimeTrans = timeTrans
var GradeTrans = gradeTrans
