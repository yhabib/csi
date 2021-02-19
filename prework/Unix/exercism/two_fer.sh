#!/usr/bin/env bash

# Notes:
# ${1:-you} defines a default value for variable $1. More on this here: https://www.debuntu.org/how-to-bash-parameter-expansion-and-default-values
# Alternative would be: echo "One for ${1:-you}, one for me."

function getName () {
  if [ -z "$1" ]; then echo "you"; else echo "$1"; fi
}

main () {
  echo "One for $(getName "$1"), one for me."
}

main "$@"