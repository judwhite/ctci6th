> Partition: Write code to partition a linked list around a value x, such that all nodes less than x come before all nodes greater than or equal to x. If x is contained within the list, the value of x only need to be after the elements less than x (see below). The partition element x can appear anywhere in the "right partition"; it does not need to appear between the left and right partitions.
>
> EXAMPLE  
> Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 (partition = 5)  
> Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8

```
BenchmarkPartition/[3_5_8_5_10_2_1]-8                    5000000           286 ns/op
BenchmarkPartition/[3_5_8_5_10_2_1_1]-8                  5000000           324 ns/op
BenchmarkPartition/[1_1_1_1_1]-8                        10000000           195 ns/op
BenchmarkPartition/[5_5_5]-8                            20000000           118 ns/op
BenchmarkPartition/[10_5_4]-8                           10000000           120 ns/op
BenchmarkPartition/[10_5_4_3_2_1]-8                      5000000           247 ns/op
BenchmarkPartition/[2_1_5_5_9_9_0_3]-8                   5000000           333 ns/op
BenchmarkPartition/[9_1_1]-8                            10000000           124 ns/op
BenchmarkPartition/[9_9_1]-8                            10000000           121 ns/op
BenchmarkPartition/[1_1_9]-8                            20000000           121 ns/op
BenchmarkPartition/[1_9_9]-8                            20000000           118 ns/op
BenchmarkPartition/[1_9_1]-8                            10000000           127 ns/op
BenchmarkPartition/[9_1_9]-8                            20000000           125 ns/op
BenchmarkPartition/[2_1_5_5_9_9_0_3_2_1...]-8             100000         13438 ns/op

BenchmarkPartitionStable/[3_5_8_5_10_2_1]-8              5000000           297 ns/op
BenchmarkPartitionStable/[3_5_8_5_10_2_1_1]-8            5000000           342 ns/op
BenchmarkPartitionStable/[1_1_1_1_1]-8                  10000000           210 ns/op
BenchmarkPartitionStable/[5_5_5]-8                      10000000           127 ns/op
BenchmarkPartitionStable/[10_5_4]-8                     10000000           126 ns/op
BenchmarkPartitionStable/[10_5_4_3_2_1]-8                5000000           251 ns/op
BenchmarkPartitionStable/[2_1_5_5_9_9_0_3]-8             5000000           332 ns/op
BenchmarkPartitionStable/[9_1_1]-8                      10000000           125 ns/op
BenchmarkPartitionStable/[9_9_1]-8                      10000000           125 ns/op
BenchmarkPartitionStable/[1_1_9]-8                      10000000           125 ns/op
BenchmarkPartitionStable/[1_9_9]-8                      10000000           125 ns/op
BenchmarkPartitionStable/[1_9_1]-8                      10000000           125 ns/op
BenchmarkPartitionStable/[9_1_9]-8                      10000000           124 ns/op
BenchmarkPartitionStable/[2_1_5_5_9_9_0_3_2_1...]-8       200000          7700 ns/op
```
