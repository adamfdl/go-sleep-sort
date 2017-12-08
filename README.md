Go Sleep Sort
---

Sleep Sort is very funny way to sort an array of non-negative numbers. The basic
idea is, for number `x` wait for time proportional to `x` (example `x` seconds)
and then print it.

`go-sleep-sort` is implementation of Sleep Sort algorithm in Golang.

## How to run
```
go run sleepsort.go 3 1 2

[3]     Sleep time:  ------- 1s ------- 2s ------- 3s
[1]     Sleep time:  ------- 1s
[2]     Sleep time:  ------- 1s ------- 2s
1 2 3

======STATS======

1]Sorted array: [1 2 3]
2]Num of goroutine spawned: 3 (4 if main Goroutine is included)
3]Time spent running sleep sort: 3.000919649s
```

## Source
Code was cloned from https://github.com/arpitbbhayani/go-sleep-sort