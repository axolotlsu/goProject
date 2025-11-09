package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/axolotlsu/goProject/learn/word_counter/wordcounter"
)

func main() {
	// 支持命令行选项
	filePath := flag.String("file", "", "text file to read (if empty, read from stdin). You can also provide filename as positional argument.")
	topN := flag.Int("top", 0, "show top N words (0 means show all)")
	jsonOut := flag.Bool("json", false, "output as JSON (an array of {word, count})")
	flag.Parse()

	// 优先使用 -file，其次使用第一个位置参数（如果提供了）
	if *filePath == "" && flag.NArg() > 0 {
		*filePath = flag.Arg(0)
	}

	// 打开文件或使用 stdin
	var f *os.File
	var err error
	if *filePath == "" {
		f = os.Stdin
	} else {
		f, err = os.Open(*filePath)
		if err != nil {
			log.Fatalf("failed to open file %q: %v", *filePath, err)
		}
		defer f.Close()
	}

	// 统计词频（使用我们实现的包）
	counts, err := wordcounter.CountWords(f)
	if err != nil {
		log.Fatalf("CountWords error: %v", err)
	}

	// 把 map 转为排序切片
	type kv struct {
		Word  string `json:"word"`
		Count int    `json:"count"`
	}
	list := make([]kv, 0, len(counts))
	for w, c := range counts {
		list = append(list, kv{w, c})
	}

	// 按 Count 降序，Count 相同时按 Word 升序
	sort.Slice(list, func(i, j int) bool {
		if list[i].Count == list[j].Count {
			return list[i].Word < list[j].Word
		}
		return list[i].Count > list[j].Count
	})

	limit := *topN
	if limit <= 0 || limit > len(list) {
		limit = len(list)
	}
	list = list[:limit]

	// 输出（JSON 或 人类可读表格）
	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(list); err != nil {
			log.Fatalf("failed to encode json: %v", err)
		}
		return
	}

	// 表格输出：count then word
	for _, kv := range list {
		fmt.Printf("%6d  %s\n", kv.Count, kv.Word)
	}
}
