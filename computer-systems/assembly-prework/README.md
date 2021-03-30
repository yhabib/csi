# Assembly

To assemble the code in BigSur:
```sh
nasm -fmacho64 <file>.asm && ld -L /Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/lib -lSystem <file>.o
```

## Registers

* `rdi` 1st argument
* `rax` temporal reg, for return and `syscal`


## Debugging w/ lldb
1. `lldb <binary>`
2. `breakpoint set --file <.asm> --line <line>`
3. `run`
4. `thread continue | step-in | step-over | step-out`
5. `register read`



## Resources
- [LLDB](https://developer.apple.com/library/archive/documentation/IDEs/Conceptual/gdb_to_lldb_transition_guide/document/lldb-command-examples.html)