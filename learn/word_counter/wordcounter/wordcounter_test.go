package wordcounter

import (
	"strings"
	"testing"
)

func TestCountWords_Basic(t *testing.T) {
	input := "Hello, world! Hello 123 hello-world don't DONT"

	// 我们的 CountWords 实现会：
	//  - 小写化
	//  - 去掉撇号 (so "don't" -> "dont")
	//  - 把数字作为词 ("123")
	//  - 把 hello-world 拆成 "hello" 和 "world"
	want := map[string]int{
		"hello": 3, // Hello, Hello, hello-world -> hello x3
		"world": 2, // "world" + part of "hello-world"
		"123":   1,
		"dont":  2, // "don't" + "DONT" -> both normalized to "dont"
	}

	got, err := CountWords(strings.NewReader(input))
	if err != nil {
		t.Fatalf("CountWords error: %v", err)
	}

	// 比较 want 中的每个键是否和 got 匹配
	for k, wantCount := range want {
		if got[k] != wantCount {
			t.Errorf("word %q: want %d, got %d", k, wantCount, got[k])
		}
	}
}
