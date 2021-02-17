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

Nice linter to work with shell scripts: `shellcheck`

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
my_array=(1 apple "long string" $(date))
echo "${my_array[1]}"
my_array[1]=orange
echo "${my_array[1]}"
echo "${my_array[2]}"
echo "${my_array[3]}"
echo "${my_array[@]}"             // Refers to all the array values
echo "${#my_array[@]}"            // Refers to the size of the array

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

Some string operations 
```sh
STRING="Hello World!"
${#STRING}                                # Get the lenght of a string, in this case 12

SUBSTRING="ld"
echo expr index "$STRING" "$SUBSTRING"    # Finds first match of any letter of SUBSTRING in STRING
temp=${STRING%%SUBSTRING*}                # Mac alternative https://stackoverflow.com/a/17615946/3364845
LEN=$((${#temp} + 1))

POS=1
LEN=3
echo ${STRING:$POS:$LEN}                  # Substring of STRING, it starts counting at 0 -> ell
echo ${STRING:$POS}                       # If LEN omitted then till the end

STRING="to be or not to be"
echo ${STRING[@]/be/eat}                  # Replaces first match of be w/ eat -> to eat or not to be
echo ${STRING[@]//be/eat}                 # Replaces all matches of be w/ eat -> to eat or not to eat
echo ${STRING[@]// not/}                  # Deletes all matches of not -> to be or to be
echo ${STRING[@]/#to be/eat now}          # Replaces if at beggining of string -> eat now or not to be
echo ${STRING[@]/%be/eat}                 # Replaces if at end of string -> to be or not to eat
echo ${STRING[@]/%be/be on $(date +%Y-%m-%d)}   
```

## Decision making

Syntax if-else:
```sh
NAME="George"
if [ "$NAME" = "John" ]; then
  echo "John Lennon"
elif [ "$NAME" = "George" ]; then
  echo "George Harrison"
else
  echo "This leaves us with Paul and Ringo"
fi
```

Syntax switch:
```sh
mycase=1
case $mycase in
    1) echo "You selected bash";;
    2) echo "You selected perl";;
    3) echo "You selected phyton";;
    4) echo "You selected c++";;
    5) exit
esac
```

Numeric comparisons:
```sh
comparison    Evaluated to true when
$a -lt $b      $a < $b
$a -gt $b      $a > $b
$a -le $b      $a <= $b
$a -ge $b      $a >= $b
$a -eq $b      $a is equal to $b
$a -ne $b      $a is not equal to $b
```

String comparisons:
```sh
comparison     Evaluated to true when
"$a" = "$b"     $a is the same as $b
"$a" == "$b"    $a is the same as $b
"$a" != "$b"    $a is different from $b
-z "$a"         $a is empty
```

## Looping
```sh
FRUITS=(Apple Banana Pear Grape)          # Creates array
for N in "${FRUITS[@]}" ; do              # Iterates
  if [ "${N:0:1}" == "A" ]; then          # Check for first letter
    echo "It is an $N"
  else
    echo "It is a $N"
  fi
done
```
```sh
COUNT=4
while [ $COUNT -gt 0 ]; do
  echo "Value of count is: $COUNT"
  COUNT=$(($COUNT - 1))
done
```

It executes the loop while the condition is false
```sh
COUNT=1
until [ $COUNT -gt 5 ]; do
  echo "Value of count is: $COUNT"
  COUNT=$(($COUNT + 1))
done
```

`continue` to skip current interation
`break` to skip entire rest of the loop

## Functions

Examples:
```sh
function function_B {
  echo "Function B."
}
function function_A {
  echo "$1"
}
function adder {
  echo "$(($1 + $2))"
}

function_A "Function A."     
function_B                   
adder 12 56   
```

## Commands
grep, sort, uniq, cut wc, sed, strinsg, head, tail, awk
ps, fg, bg, jobs, kill,
cd, cat, cp, rm, ls, mv, ln, file, chmod, chown, du, mkdir, mkfifo,
man, help, echo, apropos, tee, test, xargs


### grep

The command `grep` is used to search for Patterns in each File.

Flags:
* `-e`



### which

The command `which` tells you location of another command. Eg:
```sh
which bash
```