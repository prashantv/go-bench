## Argwriter Results

Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkWriteIndirectVal_Val|20000000|       109 ns/op
BenchmarkWriteIndirectVal_Ptr|10000000|       133 ns/op
BenchmarkWriteIndirectPtr|10000000|       194 ns/op
BenchmarkWriteDirectVal_Val|20000000|       110 ns/op
BenchmarkWriteDirectVal_Ptr|20000000|       112 ns/op
BenchmarkWriteDirectPtr|10000000|       186 ns/op
BenchmarkNoWriter|20000000|       127 ns/op

Generated using go version go1.4.2 darwin/amd64
