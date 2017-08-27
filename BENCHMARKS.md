# Benchmarks
Although the goal was to get a balanced library for speed/memory/features, it's still interesting to see some numbers. So here are results:

### results for getting value from random cell
purpose - get a cell's value at random column and row. It can be outside of current dimension of sheet or inside of it
```
BenchmarkRandomGet/excelize-4            1000000              1411 ns/op             161 B/op         20 allocs/op
BenchmarkRandomGet/tealeg-4             10000000               124 ns/op               0 B/op          0 allocs/op
BenchmarkRandomGet/xlsx-4                3000000               514 ns/op              64 B/op          6 allocs/op
```
tealeg is fastest here, although this library took **second place**



### results for setting value to random cell
purpose - set a random value to a cell's at random column and row. It can be outside of current dimension of sheet or inside of it
```
BenchmarkRandomSet/excelize-4             500000              2426 ns/op             416 B/op         25 allocs/op
BenchmarkRandomSet/tealeg-4              5000000               311 ns/op              16 B/op          1 allocs/op
BenchmarkRandomSet/xlsx-4                3000000               595 ns/op              80 B/op          7 allocs/op
```
tealeg is fastest here, although this library took **second place**



### results for settings style to random cell 
purpose - set a style to a random cell at random column and row. It can be outside of current dimension of sheet or inside of it
```
BenchmarkRandomSetStyle/excelize-4           200           7758241 ns/op         1276697 B/op     135760 allocs/op
BenchmarkRandomSetStyle/tealeg-4        10000000               120 ns/op               0 B/op          0 allocs/op
BenchmarkRandomSetStyle/xlsx-4           3000000               465 ns/op              64 B/op          6 allocs/op
```
tealeg is fastest here, although this library took **second place**



### results for reading quite big file - 2.8mb xlsx with about 18k rows
purpose - reading value for each row. [Download](https://www.dropbox.com/s/u27pjfzmyu1vbmx/example_big.xlsx?dl=0)
```
BenchmarkReadBigFile/excelize-4                1        10691191886 ns/op       2624469112 B/op 55370203 allocs/op
BenchmarkReadBigFile/tealeg-4                  1        3212188432 ns/op        687663552 B/op  19086008 allocs/op
BenchmarkReadBigFile/xlsx-4                    1        3084188730 ns/op        505871936 B/op  16668862 allocs/op
```
this library is **fastest** here, also it took less memory than any other lib.



### results for updating quite big file - 2.8mb xlsx with about 18k rows
purpose - reading value for each row and save file. [Download](https://www.dropbox.com/s/u27pjfzmyu1vbmx/example_big.xlsx?dl=0)
```
BenchmarkUpdateBigFile/excelize-4              1        13112628571 ns/op       3103088296 B/op 57420973 allocs/op
BenchmarkUpdateBigFile/tealeg-4                1        6769859126 ns/op        1119744888 B/op 24478898 allocs/op
BenchmarkUpdateBigFile/xlsx-4                  1        5426188223 ns/op        752162640 B/op  22417594 allocs/op
```
this library is **fastest** here, also it took less memory than any other lib.



### results for reading quite huge file - 73mb xlsx with about 1m rows
purpose - reading value for each row. [Download](https://www.dropbox.com/s/7zpqf0qw1yawviv/example_huge.xlsx?dl=0)
```
BenchmarkReadHugeFile/xlsx-4                   1        106438125232 ns/op      13631668472 B/op        408849745 allocs/op
```
It's a great pity, but **only** this library could open and read this file



### results for updating quite huge file - 73mb xlsx with about 1m rows
purpose - reading value for each row and save file. [Download](https://www.dropbox.com/s/7zpqf0qw1yawviv/example_huge.xlsx?dl=0)
```
BenchmarkUpdateHugeFile/xlsx-4                 1        226916121547 ns/op      20354324704 B/op        526517282 allocs/op
```
It's a great pity, but **only** this library could open, read and save this file