# Benchmark test result of queue implementions in go.
```
BenchmarkQueue-2                        	1000000000	         2.724 ns/op	       0 B/op	       0 allocs/op
BenchmarkQueueWithGrow-2                	1000000000	        10.25 ns/op	      32 B/op	       0 allocs/op
BenchmarkRingBuffer-2                   	1000000000	         4.057 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferWithGrow-2           	1000000000	        13.34 ns/op	      31 B/op	       0 allocs/op
BenchmarkRingBufferWithCheck-2          	1000000000	         5.832 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferWithCheckAndGrow-2   	1000000000	        16.27 ns/op	      31 B/op	       0 allocs/op
BenchmarkSlice-2                        	1000000000	         4.329 ns/op	      15 B/op	       0 allocs/op
BenchmarkSliceWithGrow-2                	1000000000	        10.81 ns/op	      44 B/op	       0 allocs/op
```