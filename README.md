# Benchmark test result of queue implementions in go.
```
2023-05-24T15:53:18.7161362Z goos: linux
2023-05-24T15:53:18.7162164Z goarch: amd64
2023-05-24T15:53:18.7162879Z pkg: queue
2023-05-24T15:53:18.7163268Z cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
2023-05-24T15:53:18.7163547Z BenchmarkQueue
2023-05-24T15:53:21.4781736Z BenchmarkQueue-2                       	1000000000	         2.502 ns/op	       0 B/op	       0 allocs/op
2023-05-24T15:53:21.4839794Z BenchmarkQueueWithGrow
2023-05-24T15:53:33.0756010Z BenchmarkQueueWithGrow-2               	1000000000	        10.43 ns/op	      32 B/op	       0 allocs/op
2023-05-24T15:53:33.0756599Z BenchmarkRingBuffer
2023-05-24T15:53:37.4783552Z BenchmarkRingBuffer-2                  	1000000000	         3.994 ns/op	       0 B/op	       0 allocs/op
2023-05-24T15:53:37.4784826Z BenchmarkRingBufferWithGrow
2023-05-24T15:53:54.5513634Z BenchmarkRingBufferWithGrow-2          	1000000000	        15.53 ns/op	      31 B/op	       0 allocs/op
2023-05-24T15:53:54.5514464Z BenchmarkRingBufferChecked
2023-05-24T15:53:59.9026587Z BenchmarkRingBufferChecked-2           	1000000000	         4.858 ns/op	       0 B/op	       0 allocs/op
2023-05-24T15:53:59.9027083Z BenchmarkRingBufferCheckedWithGrow
2023-05-24T15:54:18.6227623Z BenchmarkRingBufferCheckedWithGrow-2   	1000000000	        16.99 ns/op	      31 B/op	       0 allocs/op
2023-05-24T15:54:18.6228348Z BenchmarkSlice
2023-05-24T15:54:23.5377368Z BenchmarkSlice-2                       	1000000000	         4.445 ns/op	      15 B/op	       0 allocs/op
2023-05-24T15:54:23.5377880Z BenchmarkSliceWithGrow
2023-05-24T15:54:36.3134541Z BenchmarkSliceWithGrow-2               	1000000000	        11.58 ns/op	      44 B/op	       0 allocs/op
```