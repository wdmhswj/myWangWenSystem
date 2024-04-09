package parse

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func Parse_html(file_path string) {
	// 打开本地 HTML 文件
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	// 解析 HTML 文件
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println("解析 HTML 文件失败:", err)
		return
	}

	// 打印 HTML 文件中的标题
	title := findTitle(doc)
	fmt.Println("标题:", title)

}

// 辅助函数，用于在 HTML 文档中查找标题
func findTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title := findTitle(c); title != "" {
			return title
		}
	}
	return ""
}
