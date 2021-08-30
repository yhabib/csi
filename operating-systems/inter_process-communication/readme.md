# Inter-Process Communication

[Content](https://beej.us/guide/bgipc/html/multi/index.html)

File descriptiors are _ints_ that can be use with system calls like `open()`, `creat()`, `close` and `write()`. Ex:

* `stdin` is "0"
* `stdout` is "1"
* `stderr` is "2"

## Examples

### Pipes

**Implete `ls | wc -l` in C**

