#!/usr/bin/env bash

# Notes:
#  in Shell everything is a string 
#  bc math library to handle floating point numbers and their comparison -> bc -1 to load the library so scale is big enough for decimal part
#  to check if an argument is a number || [[ $1 != *[[:digit:]]* ]]  https://stackoverflow.com/a/3029040/3364845 
#  alternative: re='^[0-9.-]+$' ; ! [[ $1 =~ $re && $2 =~ $re ]]
#  to handle errors, exit with 1. 
#  non-error programs are exited with 0

main () {
  if [[ $# -ne 2 ]] || [[ $1 != *[[:digit:]]* ]] || [[ $2 != *[[:digit:]]* ]] ; then
    echo "Error"
    exit 1
  fi

  d=$(echo "($1 * $1) + ($2 * $2)" | bc -l)  
  if (( $(echo "$d <= 1.0" | bc -l) )); then 
    echo 10
  elif (( $(echo "$d <= 25.0" | bc -l) )); then 
    echo 5
  elif (( $(echo "$d <= 100.0" | bc -l) )); then 
    echo 1
  else echo 0
  fi
}

# call main with all of the positional arguments
main "$@"

