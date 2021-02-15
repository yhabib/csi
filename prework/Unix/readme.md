# Unix

First line of a shell script file begins with a `#`, and it tells the OS that the file is a set of commands to be fed into the interpreted indicated by its path. Eg:
`#!/bin/bash`

The common extension for this type of files is: `.sh`

After the first line any text followed by `#` is considered a comment.

To execute it:
```sh
./<name_of_script.sh>
```

If yuo get an error like `permission denied: ./arrays.sh` means that we are not allow to run this file yet, so lets give it *executable* permissions for us.

`chmod +x FILENAME`

In mac some commands are different, to use UNIX ones:
`brew install coreutils`

All GNU ones are accesible by prefixing them with `g`, eg: `gdate`

## Variables

A variable can contain a number, a character or a string of characters.
It is case sensitive and can consist of a combination of letters and the `_`.
**No space is permitted around the `=` sign**

```sh
PRICE_BALL=5
greeting='Hello    world'
character=ABC
```

To print the variable, either `$variable` or `${variable}`:
```sh
echo $variable
echo "${variable} yey yo! "
```

We can also assign commands to variables by wrapping them with either `\`\`` or `$()`, this is referred as *substitution*. Eg:
```sh
file_name=./file_$(date +%Y-%m-%d).txt   # Output: ./file_2021-02-15.txt
```

### Arguments

Arguments passed to the script can be referenced with `#`. Eg:

test.sh
```sh
echo $1
echo $2
echo $3
echo $#     # Prints the amount of arguments pased
```

```sh
./test.sh hello 2 "ma name is"
```

### Arrays
Very similar to variable initialization

```sh
my_array=(1 apple "long string" `date`)
echo $my_array[1]
my_array[1]=orange
echo $my_array[1]
echo $my_array[2]
echo $my_array[3]
```

## Operations

### Arithmetic
Arithmetic expression `$((expression))`, available operators are: `+,-, *, /, %, **`

Ex:
```sh
#!/bin/bash
COST_APPLE=5
COST_BANANA=10
APPLES=3
BANANAS=2
TOTAL=$((APPLES * COST_APPLE + BANANAS * COST_BANANA))
echo "$TOTAL"
```

### String

To know length of string: `${#<variable_holding_string>}

## Commands
grep, sort, uniq, cut wc, sed, strinsg, head, tail, awk
ps, fg, bg, jobs, kill,
cd, cat, cp, rm, ls, mv, ln, file, chmod, chown, du, mkdir, mkfifo,
man, help, echo, apropos, tee, test, xargs

`shellcheck`

### grep

The command `grep` is used to search for Patterns in each File.

Flags:
* `-e`



### which

The command `which` tells you location of another command. Eg:
```sh
which bash
```