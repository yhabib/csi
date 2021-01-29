# CSI


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


#### DataTypes

|   Type|   Size(bits)|   min|   max|   u max|
|---|---|---|---|---|
|   char|8   |  -128 |  127 |   255|
| short  |  16 |  -2 *15 -1  | 2 *15 - 1 | 2 *15 - 1 |
| int  |  32 |  -2 *31 - 1 | 2 *31 - 1 |  2 *32 - 1 |
| long  |  64 |  -2 * 63 - 1 | 2 * 63 - 1 |  2 * 64 - 1 |

## Extra

### VIM
Three modes:
* normal -> `ESC` goes to normal mode
* insert  -> `i` goes to insert mode
* command line -> `:` goes to command line mode

* Exit: `:q` or `:q!`
* Save `:w`
* Save only if changes `:x`
* Can be combined into: `:wq`


#### Navigating

**In Norma Mode**

* Show numbers: `:set number`
* Jump to line: `:5`
* Jump to end of file: `:$`
* Jump to end of line: `$`

<!-- https://opensource.com/article/19/3/getting-started-vim -->


#### Editing

**In Norma Mode**
* Deleting line: `dd`
* Undo: `u`
* Select mode: `v`
* To copy after selecting(yank): `y`
* New line below: `o`
* Paste: `p`


#### Searching:

* Search forward: `/<term>` -> move to next: `n` or to previous: `N`
* Search backward: `?<term>`


#### Bonus

* Split mode: `:split <file2>` or `:vsplit <file2>`
* Toggle between panels: `ww`