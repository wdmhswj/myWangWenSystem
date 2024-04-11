package utils

import (
	"crawler/structs"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func SaveAsJson(filename string, entity structs.RankingList_qidian, dir string) {
	// 若不存在data目录则创建
	ok, err := FileDirExist(dir)
	if !ok {
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = os.Mkdir(dir, 0750)
		if err != nil {
			fmt.Println("failed to mkdir:", err.Error())
			return
		}
	}
	filename += ".json"
	filename = filepath.Join(dir, filename)
	// filename = dir + filename + ".json"
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

func LoadJsonAsStruct(filename string, dir string) structs.RankingList_qidian {
	if ok, _ := FileDirExist(dir); !ok {
		fmt.Println(dir + "目录return")
		return structs.RankingList_qidian{}
	}
	filename = filepath.Join(dir, filename)

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

func FileDirExist(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func GetAllFileName(dirname string) ([]string, error) {
	// 打开当前目录
	dir, err := os.Open(dirname)
	if err != nil {
		fmt.Println("打开目录失败:", err)
		return nil, errors.New("打开目录失败: " + err.Error())
	}
	defer dir.Close()

	// 读取目录中的文件名
	files, err := dir.Readdirnames(-1) // -1 表示读取所有文件名
	if err != nil {
		fmt.Println("读取文件名失败:", err)
		return nil, errors.New("读取文件名失败: " + err.Error())
	}

	return files, nil
}

func MoveFiles(sourceDir, targetDir string) error {
	// 打开源目录
	source, err := os.Open(sourceDir)
	if err != nil {
		return err
	}
	defer source.Close()

	// 读取源目录下的所有文件和子目录
	entries, err := source.Readdir(0)
	if err != nil {
		return err
	}

	// 遍历文件和子目录
	for _, entry := range entries {
		// if entry.IsDir() {
		// 	// 如果是子目录，则递归调用 moveFiles 函数
		// 	err := moveFiles(filepath.Join(sourceDir, entry.Name()), targetDir)
		// 	if err != nil {
		// 		return err
		// 	}
		// } else {
		if !entry.IsDir() {
			// 如果是文件，则移动到目标目录
			sourceFile := filepath.Join(sourceDir, entry.Name())
			targetFile := filepath.Join(targetDir, entry.Name())
			err := os.Rename(sourceFile, targetFile)
			if err != nil {
				return err
			}
			fmt.Printf("移动文件 %s 到 %s\n", sourceFile, targetFile)
		}
	}

	return nil
}
