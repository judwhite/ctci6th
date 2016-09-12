> Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

Approach:
- Use a `uint32` as the 4 byte pixel.
- Take a `[][]uint32` to represent the matrix.
- Some hand rolling:
  - Assume 3x3 matrix. We'll call `[0][0]` the upper left, the first dimension will be the rows.
```
1 2 3      7 4 1
4 5 6  ->  8 5 2
7 8 9      9 6 3
```
  - This case may be too simple, let's try 4x4.
```
0 1 2 3      C 8 4 0
4 5 6 7  ->  D 9 5 1
8 9 A B      E A 6 2
C D E F      F B 7 3
```
  - Visually this is easy to verify. Using the rotated matrix, start in the upper right, read down, then move back to the top and left 1 column. Both matrices read `0123456789ABCDEF`.
  - Let's look at what happened to `0`. It shifted down `0`, its current X position, and shifted left `3`, its current distance from the bottom (N=4, x = N-1-y = 4-1-0 = 3).
  - Let's look at D. It moved from (1, 3) to (0, 1). Y became the old X. X became its current distance from the bottom, `0`.
  - Let's look at 9. It moved from (1, 2) to (1, 1). Y became the old X. X became its current distance from the bottom, `1`.
  - Let's look at 5. It moved from (1, 1) to (2, 1). Y became the old X. X became its current distance from the bottom, `2`.
  - Let's look at 6. It moved from (2, 1) to (2, 2). Y became the old X. X became its current distance from the bottom, `2`.
  - Let's look at A. It moved from (2, 2) to (1, 2). Y became the old X. X became its current distance from the bottom, `1`.
  - `let y2 = x1, x2 = N-1-y1`.
  - For each "ring", we'll need to do 4 operations N-(row*2)-1 times.
  - We'll start at col=row and work until N-1-row
  - We only need to do N/2 rows

Walking through it manually helped. [The code](https://github.com/judwhite/ctci6th/blob/master/pg91-1.7-rotatematrix/rotatematrix.go) is easier to comprehend.

```
BenchmarkRotate/0x0-8         	500000000	         3.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkRotate/1x1-8         	500000000	         3.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkRotate/2x2-8         	100000000	        10.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkRotate/3x3-8         	100000000	        15.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkRotate/4x4-8         	50000000	        25.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkRotate/5x5-8         	50000000	        39.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkRotate/6x6-8         	20000000	        60.1 ns/op	       0 B/op	       0 allocs/op
```
