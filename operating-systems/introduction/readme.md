# Introduction

## Pre class exercises

**Difference between `man` and `help`?**
- `help` is a built-in command in the bash shell(and only that shell) that documents some of the builtin commands and keywords in a particular shell.
- `man` is a system-wide documentation system for individual commands, API functions, concepts, ...

We can use `type` to find out the type of a command. Eg:
- `type help`: shell builtin 
- `type man`: hashed

`man` has **nine** sections that can be access by providing a number as a second argument. Eg: `man 2 printf`. Some of the sections, rest can be found with `man man`:
1. Executable programs or shell commands
2. System calls (functions provided by the kernel)
3. Library calls(function withing program libraries like C)


## In class exercises

**Find number of executable files**

```bash
find / -executable -type f | wc -l
```

**What types of files are executable on this system?**

```bash
find / -executable -type f | wc -l
```



 
## Additional resources

* https://askubuntu.com/questions/20752/how-can-i-search-within-a-manpage
* https://vim.rtorr.com/
* https://explainshell.com/
