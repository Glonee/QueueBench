# Benchmark test result of queue implementions in go.
```
BenchmarkTwoStagesQueue-2        	1000000000	         4.777 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBuffer-2            	1000000000	         4.735 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferWithCheck-2   	1000000000	         6.753 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice-2                 	1000000000	        30.31 ns/op	       8 B/op	       1 allocs/op
```