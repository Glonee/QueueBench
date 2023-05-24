# Benchmark test result of queue implementions in go.
```
BenchmarkQueue-2                       	1000000000	         3.055 ns/op	       0 B/op	       0 allocs/op
BenchmarkQueueWithGrow-2               	1000000000	        22.47 ns/op	      32 B/op	       0 allocs/op
BenchmarkRingBuffer-2                  	1000000000	         5.671 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferWithGrow-2          	1000000000	        28.62 ns/op	      31 B/op	       0 allocs/op
BenchmarkRingBufferChecked-2           	1000000000	         8.630 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferCheckedWithGrow-2   	1000000000	        32.38 ns/op	      31 B/op	       0 allocs/op
BenchmarkSlice-2                       	1000000000	        10.30 ns/op	      15 B/op	       0 allocs/op
BenchmarkSliceWithGrow-2               	1000000000	        24.68 ns/op	      44 B/op	       0 allocs/op
```