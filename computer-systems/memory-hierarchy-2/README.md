## Evolution

Original:
```
goos: darwin
goarch: amd64
pkg: metrics
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMetrics/Average_age-12                      342           3544423 ns/op
BenchmarkMetrics/Average_payment-12                   44          26793051 ns/op
BenchmarkMetrics/Payment_stddev-12                    21          53453941 ns/op
PASS
ok      metrics 6.781s
```


int uint8
```
goos: darwin
goarch: amd64
pkg: metrics
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMetrics/Average_age-12                      308           3438952 ns/op
BenchmarkMetrics/Average_payment-12                   38          27188489 ns/op
BenchmarkMetrics/Payment_stddev-12                    21          54160311 ns/op
PASS
ok      metrics 7.143s
```


types for average age
```
goos: darwin
goarch: amd64
pkg: metrics
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMetrics/Average_age-12                      916           1112333 ns/op
BenchmarkMetrics/Average_payment-12                   44          27123512 ns/op
BenchmarkMetrics/Payment_stddev-12                    22          52901466 ns/op
PASS
ok      metrics 5.812s
````

maps
```
goos: darwin
goarch: amd64
pkg: metrics
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMetrics/Average_age-12                      940           1151987 ns/op
BenchmarkMetrics/Average_payment-12                  100          11261466 ns/op
BenchmarkMetrics/Payment_stddev-12                    51          23030894 ns/op
PASS
ok      metrics 6.314s
```

arrays everywhere
```
goos: darwin
goarch: amd64
pkg: metrics
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMetrics/Average_age-12                     1993            553905 ns/op
BenchmarkMetrics/Average_payment-12                  214           5614393 ns/op
BenchmarkMetrics/Payment_stddev-12                   169           6976766 ns/op
PASS
ok      metrics 5.585s
```

Leverage the fact we now the limits of the progra:
```
goos: darwin
goarch: amd64
pkg: metrics
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMetrics/Average_age-12                    20980             56746 ns/op
BenchmarkMetrics/Average_payment-12                 1683            650710 ns/op
BenchmarkMetrics/Payment_stddev-12                   631           1851020 ns/op
PASS
ok      metrics 4.778s
```
