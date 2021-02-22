#!/usr/bin/env bash

# Notes:
# $#: Gives number of arguments
# $1 != ?(-)+([0-9]): 
# For logic operators in if statements: [[ A || B ]] or [ A ] || [ B ]
# +([0-9]) -> Pattern matching


main () {
  year="$1"
  if [[ $# -ne 1 || $year != +([0-9]) ]]; then
    echo 'Usage: leap.sh <year>'
    exit 1
  elif [ $((year % 4)) -eq 0 ] && [ $((year % 100)) -ne 0  ] || [ $((year % 400)) -eq 0 ]; then 
    echo true
  else 
    echo false 
  fi
}

main "$@"