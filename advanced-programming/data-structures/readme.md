## Data Structures

### Notes
- string interning
- all bytes depend on how we interpret them
    - compare string exercises → maps the unsafe.Pointer to a custom struct with pointer a nd len
- the pointers to the arguments represent the location of the heap where the data is stored ~38min
- blum filter
- map iteration order is not guaranteed
- links:
    - [https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics](https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics)
    - [ Some insights on Maps in Golang | Hacker Noon](https://www.hackernoon.com/some-insights-on-maps-in-golang-rm5v3ywh)
    - [ A comprehensive analysis of Golang's map design](https://www.fatalerrors.org/a/a-comprehensive-analysis-of-golang-s-map-design.html)
    - go-delve


### Questions
- Array exercise -> to go through the array +8 ?? 64b / 8b = 8 bytes -> every 8 bytes I have a another number, so what do I have what do I have every byte then? random data? padding?
- > By converting a *float64 pointer to a *uint64, for instance, we can inspect the bit pattern of a floating-point variable ?
- Difference between:
```go
uint64(unsafe.Pointer(&num))      // Compiler error as it can't convert a pointer to an integer
*(*uint64)(unsafe.Pointer(&num))  // Converts first from pointer to pointer uint64 and then to integer
```