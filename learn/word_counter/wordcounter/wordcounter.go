package wordcounter

import (
	"io"
	"regexp"
	"strings"
)

func CountWords(r io.Reader) (map[string]int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// 匹配字母或数字序列（允许内部有撇号）
	// 这里允许数字开头，也会抓 "123"
	re := regexp.MustCompile(`[\p{L}\p{N}']+`)

	words := re.FindAllString(string(data), -1)
	counts := make(map[string]int, len(words))

	for _, w := range words {
		w = strings.ToLower(w)
		// 可选：去掉所有撇号，统一 dont / don't
		w = strings.ReplaceAll(w, "'", "")
		counts[w]++
	}

	return counts, nil
}
