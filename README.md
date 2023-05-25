# Benchmark test result of queue implementions in go. Recent test result can be found in GitHub Actions.
```
BenchmarkTwoStagesQueue                	1000000000	         4.432 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoStagesQueueBuffered        	1000000000	         2.636 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBuffer                    	1000000000	         4.022 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferBuffered            	1000000000	         3.710 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferWithCheck           	1000000000	         5.721 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferWithCheckBuffered   	1000000000	         6.872 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice                         	1000000000	        24.47 ns/op	       8 B/op	       1 allocs/op
BenchmarkSliceBuffered                 	1000000000	         7.242 ns/op	      16 B/op	       0 allocs/op
```
