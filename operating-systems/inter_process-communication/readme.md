# Inter-Process Communication

[Content](https://beej.us/guide/bgipc/html/multi/index.html)

The **Man pages** have different levels:

- 1: General Commands
- 2: System Calls
- 3: Library functions, covering in particular the C standard library
- 4: Special files (usually devices, those found in /dev) and drivers
- 5: File formats and conventions
- 6: Games and screensavers
- 7: Miscellanea
- 8: System administration commands and daemons

## API

### Fork

`fork()`

- 0: If it returns 0, you are the child process. You can get the parent's PID by calling `getppid()`. Of course, you can get your own PID by calling getpid().

- -1: If it returns -1, something went wrong, and no child was created. Use `perror()` to see what happened. You've probably filled the process table—if you turn around you'll see your sysadmin coming at you with a fireaxe.

- else: Any other value returned by `fork()` means that you're the parent and the value returned is the PID of your child. This is the only way to get the PID of your child, since there is no getcpid() call (obviously due to the one-to-many relationship between parents and children.)

### File Descriptors

File descriptiors are _ints_ that can be use with system calls like `open()`, `creat()`, `close` and `write()`. Ex:

- `stdin` is "0"
- `stdout` is "1"
- `stderr` is "2"


## Examples

### Pipes

**Implete `ls | wc -l` in C**

