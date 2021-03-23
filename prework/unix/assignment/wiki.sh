#!/usr/local/bin/bash

# ---------------------------------------------------------------------------------------------------------------------------
# Notes:
# $1 might not have first letter capital -> ${1^} upper cases the first letter, ${1^^} all letters ${1,} first lower and ${1,,} all lower
# grep "<pattern>" -> Returns all lines that match the pattern
# cut -d"<delimiter>" -f<index-of-element> -> Splits string by delimiter and take element/elements form the split
# [ -n "$var" ] -> Checks if variable is set
# grep -A1 "$SUBTOPIC" "$FILE_NAME".txt | head -2 | tail -1 -> Gets the second line from a match of two lines
#   -A1 to get also next line
# [[:space:]] official `\s` value 
# grep -n print line number
#  grep -m 1 -> find first match
# printf works better for escaping characs
#  sed "s/<[^>]*>//g -> For matching innerText in HTML NODE by deleting the tags
# ---------------------------------------------------------------------------------------------------------------------------

## Todos:
# * Extract some functions?

TOPIC=${1^}
SUBTOPIC=${2^}
FILE_NAME=${1}

CONTENT=$(curl -s https://en.wikipedia.org/wiki/"$TOPIC")
echo "$CONTENT" > "$FILE_NAME".html

SHORT_DESCRIPTION=$(echo "$CONTENT" | grep "shortdescription" | cut -d">" -f2 | cut -d"<" -f1)
SHORT_DESCRIPTION="A $TOPIC is a ${SHORT_DESCRIPTION,}"
echo "$SHORT_DESCRIPTION" > "$FILE_NAME".txt
echo "----------------------------" >> "$FILE_NAME".txt

HEADINGS=$(echo "$CONTENT" | grep "<h2" | grep "mw-headline" | cut -d">" -f3 | cut -d"<" -f1)

i=1
echo "$HEADINGS" | while read -r line ; do echo "$i. ${line}" >> "$FILE_NAME".txt; i=$((i+1)); done 

if [ -n "$SUBTOPIC" ]; then
  NEXT_SUB_SECTION=$(grep -A1 "$SUBTOPIC" "$FILE_NAME".txt | head -2 | tail -1 | cut -d"." -f2 | sed -e 's/^[[:space:]]*//' | sed -e "s/[[:space:]]/_/")
  SUB_SECTION_START=$(echo "$CONTENT" | grep -n "id=\"$SUBTOPIC\"" | cut -d":" -f1)
  SUB_SECTION_END=$(echo "$CONTENT" | grep -n "id=\"$NEXT_SUB_SECTION\"" | cut -d":" -f1)
  SUB_SECTION=$(echo "$CONTENT" | sed -n "$SUB_SECTION_START,$((SUB_SECTION_END - 1))p")
  FIRST_P=$(echo "$SUB_SECTION" | grep -m 1 "<p" | cut -d">" -f2 | cut -d"." -f1)
  printf "\n" >> "$FILE_NAME".txt
  echo "The $SUBTOPIC of a $TOPIC: ${FIRST_P,}" >> "$FILE_NAME".txt
  SUB_HEADINGS=$(echo "$SUB_SECTION" | grep "<h3" | sed "s/<[^>]*>//g")
  i=1
  echo "$SUB_HEADINGS" | while read -r line ; do echo "$i. ${line}" >> "$FILE_NAME".txt; i=$((i+1)); done 
fi
