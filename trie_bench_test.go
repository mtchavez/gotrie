package trie

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

func getKeys(total int) []string {
	keys := make([]string, total)
	for i := range keys {
		sha := sha1.New()
		io.WriteString(sha, fmt.Sprintf("hello world %d", i))
		keys[i] = fmt.Sprintf("%x", string(sha.Sum(nil)))
	}
	return keys
}

func Benchmark_Add_1000(b *testing.B)    { benchAdd(b, 1000) }
func Benchmark_Add_10000(b *testing.B)   { benchAdd(b, 10000) }
func Benchmark_Add_100000(b *testing.B)  { benchAdd(b, 100000) }
func Benchmark_Add_1000000(b *testing.B) { benchAdd(b, 1000000) }

func BenchmarkMap_Add_1000(b *testing.B)    { benchMapAdd(b, 1000) }
func BenchmarkMap_Add_10000(b *testing.B)   { benchMapAdd(b, 10000) }
func BenchmarkMap_Add_100000(b *testing.B)  { benchMapAdd(b, 100000) }
func BenchmarkMap_Add_1000000(b *testing.B) { benchMapAdd(b, 1000000) }

func benchAdd(b *testing.B, total int) {
	t := New()
	keys := getKeys(total)
	b.ResetTimer()
	for i := 0; i < len(keys); i++ {
		t.Add(keys[i])
	}
}

func benchMapAdd(b *testing.B, total int) {
	m := make(map[string]interface{})
	keys := getKeys(total)
	b.ResetTimer()
	for i := 0; i < len(keys); i++ {
		m[keys[i]] = i
	}
}
