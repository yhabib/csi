# Memory Hierarchy

## For loop

Below are the metrics for each version after running `valgrind --tool=cachegrid ./<version_name>`
```
Version 1
==6510==
==6510== I   refs:      240,216,156
==6510== I1  misses:          1,071
==6510== LLi misses:          1,058
==6510== I1  miss rate:        0.00%
==6510== LLi miss rate:        0.00%
==6510==
==6510== D   refs:      112,070,983  (96,055,049 rd   + 16,015,934 wr)
==6510== D1  misses:      1,003,189  (     2,528 rd   +  1,000,661 wr)
==6510== LLd misses:      1,002,697  (     2,095 rd   +  1,000,602 wr)
==6510== D1  miss rate:         0.9% (       0.0%     +        6.2%  )
==6510== LLd miss rate:         0.9% (       0.0%     +        6.2%  )
==6510==
==6510== LL refs:         1,004,260  (     3,599 rd   +  1,000,661 wr)
==6510== LL misses:       1,003,755  (     3,153 rd   +  1,000,602 wr)
==6510== LL miss rate:          0.3% (       0.0%     +        6.2%  )
```

```
==6494== I   refs:      240,216,156
==6494== I1  misses:          1,070
==6494== LLi misses:          1,057
==6494== I1  miss rate:        0.00%
==6494== LLi miss rate:        0.00%
==6494==
==6494== D   refs:      112,070,983  (96,055,049 rd   + 16,015,934 wr)
==6494== D1  misses:     16,003,189  (     2,528 rd   + 16,000,661 wr)
==6494== LLd misses:      1,002,697  (     2,095 rd   +  1,000,602 wr)
==6494== D1  miss rate:        14.3% (       0.0%     +       99.9%  )
==6494== LLd miss rate:         0.9% (       0.0%     +        6.2%  )
==6494==
==6494== LL refs:        16,004,259  (     3,598 rd   + 16,000,661 wr)
==6494== LL misses:       1,003,754  (     3,152 rd   +  1,000,602 wr)
==6494== LL miss rate:          0.3% (       0.0%     +        6.2%  )
```
> On a modern machine, an L1 miss will typically cost around 10 cycles, an LL miss can cost as much as 200 cycles, 

V1 has a better D1 miss rate (0.3% vs 14.3%) and requires also less LL(last-level) writes that v2. 
My supposition here is that v1 makes far better use of memory location, by accessin subsequent memory slots

* Second function should take longer to run because needs to use more often the LL cache.
* From assembly point of view number of instructions should be the same
* Number of misses also make sense, the factor 16 is due the fact that _cache lines_ are 64 bytes, and because we use integers(4B) means that for each access in the second version it will use only one element out of the 16 in the cache.


## Matrix multiply
Optimizations:
* Loop unrolling -> The more we do it the better specially because of the next two optimizations
* Cache leveraging of `b[k][j]`?
* Parallelism

Best result:
```
Naive: 0.743s
Fast: 0.298s
2.50x speedup
```

## Tools
### Multipass
* `multipass launch --name <name>`
* `multipass shell <id>`
* `multipass mount $HOME/<path> <id>:<path>`
* `multipass start <id>`
* `multipass stop <id>`
* `multipass delete <id>`
* `multipass purge`

### Valgrind
* Compile w/ loggin informtion: `gcc -g <file> -o <dest_file>` also could be beneficial with low level optimization: `gcc -g -O0 <file> -o <dest_file>`
* `valgrind --tool=cachegrind ./<file>`
* `valgrind --tool=cachegrind --branch-sim=yes`
* `cg_annotate <valgrind_output>`