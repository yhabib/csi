# CSI

## Links
- [VIM basics](https://opensource.com/article/19/3/getting-started-vim)

## Prework

### C

Compile program: `cc <namme>.c`
Run program: `./<name>.out`
> https://stackoverflow.com/questions/23024016/cant-run-c-program-a-out-command-not-found

It's able to check if escape sequences are right or not still compiles

> integer division truncates: any fractional part is discarded.

* `EOF` terminal character: `CTRL + D`
* Null character: `'\0'` used as convention to indicate end of things. Eg: `char s[] = ['h', 'e', 'l', 'l', 'o', '\0']`
* Most arguments in functions are treated as **value** but not arrays, those always as **reference**.
* [getchar and putchar](https://stackoverflow.com/questions/17552458/theory-behind-getchar-and-putchar-functions)
* [c - '0'](https://stackoverflow.com/a/7403877/3364845)
* If multiple file program: `cc main.c getint.c getch.c`
* [Compiler vs Linker](https://stackoverflow.com/questions/3831312/what-are-the-differences-between-a-compiler-and-a-linker/3831354)
* [Pointer size](https://stackoverflow.com/questions/38822692/why-to-specify-a-pointer-type/38822710)
* [char_list vs char pointer](https://stackoverflow.com/questions/1011455/is-it-possible-to-modify-a-string-of-char-in-c)
* [Undefined behaviour](https://en.wikipedia.org/wiki/Undefined_behavior)
* With arrays memory allocation is "random": memory can be alloacted in different every time, so the behaviour of the program in this situations is undefined.
* Convention case: snake_case


#### DataTypes

|   Type|   Size(bits)|   min|   max|   u max|
|---|---|---|---|---|
|   char|8   |  -128 |  127 |   255|
| short  |  16 |  -2 *15 -1  | 2 *15 - 1 | 2 *15 - 1 |
| int  |  32 |  -2 *31 - 1 | 2 *31 - 1 |  2 *32 - 1 |
| long  |  64 |  -2 * 63 - 1 | 2 * 63 - 1 |  2 * 64 - 1 |


#### Format specifiers

|   Format Specifier|   Description| 
|---|---|
|%d |	Integer |
|%f |	Float |
|%c |	Character |
|%s |	String |
|%u |	Unsigned Integer |
|%ld |	Long Int |

