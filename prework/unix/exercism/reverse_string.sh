#!/usr/bin/env bash

# Notes:
# ${#variable} returns the length of the variable
# "${variable:position:length}" string formatter/accessor
# echo -n omits the newline at the end, which works for this solution well

main () {
  input="$1"
  len=${#input}
  for((i=len-1;i>=0;i--)); do echo -n "${input:$i:1}"; done
}

main "$@"
