# Inter-Process Communication

[Content](https://beej.us/guide/bgipc/html/multi/index.html)

File descriptiors are _ints_ that can be use with system calls like `open()`, `creat()`, `close` and `write()`. Ex:

- `stdin` is "0"
- `stdout` is "1"
- `stderr` is "2"

## API

### Fork

`fork()`

- 0: If it returns 0, you are the child process. You can get the parent's PID by calling `getppid()`. Of course, you can get your own PID by calling getpid().

- -1: If it returns -1, something went wrong, and no child was created. Use `perror()` to see what happened. You've probably filled the process table—if you turn around you'll see your sysadmin coming at you with a fireaxe.

- else: Any other value returned by `fork()` means that you're the parent and the value returned is the PID of your child. This is the only way to get the PID of your child, since there is no getcpid() call (obviously due to the one-to-many relationship between parents and children.)
  
## Examples

### Pipes

**Implete `ls | wc -l` in C**

