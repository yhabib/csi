#!/bin/bash
# In this exercise, you will need to write a function called ENGLISH_CALC which can process sentences

function ENGLISH_CALC {
  case $2 in
    plus) echo $(($1 + $3));;
    minus) echo $(($1 - $3));;
    times) echo $(($1 * $3));;
    5) exit
esac
}

# testing code
ENGLISH_CALC 3 plus 5
ENGLISH_CALC 5 minus 1
ENGLISH_CALC 4 times 6