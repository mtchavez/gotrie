gotrie
======

Trie in Go with a benchmark vs Go Map

## Benchmark

```shell
$ go test --bench=.*
testing: warning: no tests to run
PASS
Benchmark_Add_1000	2000000000	         0.01 ns/op
Benchmark_Add_10000	2000000000	         0.14 ns/op
Benchmark_Add_100000	       1	3190024146 ns/op
Benchmark_Add_1000000	       1	47723452295 ns/op
BenchmarkMap_Add_1000	2000000000	         0.02 ns/op
BenchmarkMap_Add_10000	2000000000	         0.02 ns/op
BenchmarkMap_Add_100000	2000000000	         0.03 ns/op
BenchmarkMap_Add_1000000	1000000000	         0.55 ns/op
ok  	github.com/mtchavez/trie	131.854s
```
