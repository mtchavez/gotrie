package trie

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

var testKeys []string

func getKeys(total int) []string {
	size := len(testKeys)
	if size < total {
		for i := size; i < total; i++ {
			sha := sha1.New()
			io.WriteString(sha, fmt.Sprintf("hello world %d", i))
			testKeys = append(testKeys, fmt.Sprintf("%x", string(sha.Sum(nil))))
		}
		return testKeys
	} else {
		return testKeys[:total]
	}
}

func Benchmark_Add_1000(b *testing.B)    { benchAdd(b, 1000) }
func Benchmark_Add_10000(b *testing.B)   { benchAdd(b, 10000) }
func Benchmark_Add_100000(b *testing.B)  { benchAdd(b, 100000) }
func Benchmark_Add_1000000(b *testing.B) { benchAdd(b, 1000000) }

func BenchmarkMap_Add_1000(b *testing.B)    { benchMapAdd(b, 1000) }
func BenchmarkMap_Add_10000(b *testing.B)   { benchMapAdd(b, 10000) }
func BenchmarkMap_Add_100000(b *testing.B)  { benchMapAdd(b, 100000) }
func BenchmarkMap_Add_1000000(b *testing.B) { benchMapAdd(b, 1000000) }

func Benchmark_Find_1000(b *testing.B)    { benchFind(b, 1000) }
func Benchmark_Find_10000(b *testing.B)   { benchFind(b, 10000) }
func Benchmark_Find_100000(b *testing.B)  { benchFind(b, 100000) }
func Benchmark_Find_1000000(b *testing.B) { benchFind(b, 1000000) }

func Benchmark_MapFind_1000(b *testing.B)    { benchMapFind(b, 1000) }
func Benchmark_MapFind_10000(b *testing.B)   { benchMapFind(b, 10000) }
func Benchmark_MapFind_100000(b *testing.B)  { benchMapFind(b, 100000) }
func Benchmark_MapFind_1000000(b *testing.B) { benchMapFind(b, 1000000) }

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

func benchFind(b *testing.B, total int) {
	t := New()
	keys := getKeys(total)
	for i := 0; i < len(keys); i++ {
		t.Add(keys[i])
	}
	b.ResetTimer()
	for i := 0; i < len(keys); i++ {
		t.Find(keys[i])
	}
}

func benchMapFind(b *testing.B, total int) {
	m := make(map[string]interface{})
	keys := getKeys(total)
	for i, key := range keys {
		m[key] = i
	}
	b.ResetTimer()
	for i := 0; i < len(keys); i++ {
		_, ok := m[keys[i]]
		if !ok {
			panic("oops")
		}
	}
}
