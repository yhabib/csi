# Assembly

To assemble the code in BigSur:
```sh
nasm -fmacho64 <file>.asm && ld -L /Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/lib -lSystem <file>.o
```

## Registers

- `rax` temporal reg, for return and `syscal`
- `rdi` 1st argument
- `rsi`, `rdx`, `rcx`
- `rsp` and `rbp` stack pointers
- `flags`

## Stack
- Operations stack and pop handle inc/dec of rsp
- At the end of a function, the stack must contain exactly the same number of elements as at the beginning 

## Debugging w/ lldb
1. `lldb tests`ˇ
2. `breakpoint set --file <.asm> --line <line>`
3. `run`
4. `thread continue | step-in | step-over | step-out`
5. `register read`
6. `exit` 
### Shortcuts
- `c` continue
- `re r` register read
- `re r <reg>` reads specific register
- `n` step over
- `br s -f <file> -l <line>` for the win.
- `b <file>:<line>` for the win.
- `br l`
- `br del <num_of_breakpoint>` delete breakpoint



## Resources
- [LLDB](https://developer.apple.com/library/archive/documentation/IDEs/Conceptual/gdb_to_lldb_transition_guide/document/lldb-command-examples.html)
- [Instruction set](https://www.felixcloutier.com/x86/)
- [FLAGS](https://en.wikipedia.org/wiki/FLAGS_register)
- [ASCII](https://en.wikipedia.org/wiki/ASCII#/media/File:USASCII_code_chart.png)