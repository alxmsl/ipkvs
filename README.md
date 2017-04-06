# ipkvs

In-process KV-storage. Look at [docs](docs/) for more info

## Examples

get/set example

```
	c := ipkvs.NewIpKvs()
	c.Logger = log.Warn

	v1 := Data{ int16(rand.Intn(265)) }
	err := c.SetMsgPack("aaa", v1)
	fmt.Println(err)

	v2 := Data{}
	err = c.GetMsgPack("aaa", &v2)
	fmt.Println(err)
	fmt.Println(v1)
	fmt.Println(v2)
```

## Benchmarks

Allocations benchmark

```
$: go test -bench=. -benchmem benches/allocation/alloc_bench_test.go
BenchmarkAllocate6-4    	50000000	        32.1 ns/op	      64 B/op	       1 allocs/op
BenchmarkAllocate7-4    	30000000	        46.3 ns/op	     128 B/op	       1 allocs/op
BenchmarkAllocate8-4    	20000000	        66.1 ns/op	     256 B/op	       1 allocs/op
BenchmarkAllocate9-4    	10000000	       151 ns/op	     512 B/op	       1 allocs/op
BenchmarkAllocate10-4   	 5000000	       292 ns/op	    1024 B/op	       1 allocs/op
BenchmarkAllocate11-4   	 2000000	       726 ns/op	    2048 B/op	       1 allocs/op
BenchmarkAllocate12-4   	 2000000	      1222 ns/op	    4096 B/op	       1 allocs/op
BenchmarkAllocate13-4   	 1000000	      1873 ns/op	    8192 B/op	       1 allocs/op
BenchmarkAllocate14-4   	  500000	      3821 ns/op	   16384 B/op	       1 allocs/op
BenchmarkAllocate15-4   	  200000	      7758 ns/op	   32768 B/op	       1 allocs/op
BenchmarkAllocate16-4   	  100000	     19485 ns/op	   65536 B/op	       1 allocs/op
BenchmarkAllocate17-4   	   50000	     33839 ns/op	  131072 B/op	       1 allocs/op
BenchmarkAllocate18-4   	   20000	     57927 ns/op	  262144 B/op	       1 allocs/op
BenchmarkAllocate19-4   	   20000	     88757 ns/op	  524288 B/op	       1 allocs/op
BenchmarkAllocate20-4   	   10000	    179909 ns/op	 1048576 B/op	       1 allocs/op
BenchmarkAllocate21-4   	    5000	    313409 ns/op	 2097152 B/op	       1 allocs/op
BenchmarkAllocate22-4   	    2000	    595710 ns/op	 4194304 B/op	       1 allocs/op
BenchmarkAllocate23-4   	    1000	   1115752 ns/op	 8388608 B/op	       1 allocs/op
BenchmarkAllocate24-4   	     500	   2067225 ns/op	16777216 B/op	       1 allocs/op
BenchmarkAllocate25-4   	     500	   2101925 ns/op	33554432 B/op	       1 allocs/op
BenchmarkAllocate26-4   	     300	   3911006 ns/op	67108864 B/op	       1 allocs/op
BenchmarkAllocate27-4   	     100	  16517353 ns/op	134217728 B/op	       1 allocs/op
BenchmarkAllocate28-4   	     100	  22168094 ns/op	268435456 B/op	       1 allocs/op
BenchmarkAllocate29-4   	      30	 283001564 ns/op	536870912 B/op	       1 allocs/op
BenchmarkAllocate30-4   	      10	 642595832 ns/op	1073741824 B/op	       1 allocs/op
PASS
ok  	command-line-arguments	57.751s
```

GC overhead on maps
 
```
$: go run benches/gc_map/gc_map_bench.go --show-types
boolmap is map[int]bool
pintmap is map[int]*int
uintptrmap is map[int]unsafe.Pointer
emptystructmap is map[int]struct{} (struct is empty here)
dynamicstructmap is map[int]struct{} (struct has pointer)
intstringmap is map[string]int
intmap is map[int]int
unsafepointermap is map[int]uintptr
staticstructmap is map[int]struct{} (struct has fixed length)
$: 
$: go run benches/gc_map/gc_map_bench.go -t=staticstructmap -s=20000000
6.041232ms
1.1 GiB
```
