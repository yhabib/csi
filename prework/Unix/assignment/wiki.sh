#!/usr/local/bin/bash

# if $1 doesnt have first letter capital then it gets 301 and fails on my side
CONTENT=$(curl -s https://en.wikipedia.org/wiki/"${1^}")
FILE_NAME=${1,}
echo "$CONTENT" > "$FILE_NAME".html

SHORT_DESCRIPTION=$(echo "$CONTENT" | grep "shortdescription" | cut -d">" -f2 | cut -d"<" -f1)
echo "$SHORT_DESCRIPTION" > "$FILE_NAME".txt
echo "----------------------------" >> "$FILE_NAME".txt

HEADINGS=$(echo "$CONTENT" | grep "<h2" | grep "mw-headline" | cut -d">" -f3 | cut -d"<" -f1)
i=1
echo "$HEADINGS" | while IFS= read -r line ; do echo "$i. ${line}" >> "$FILE_NAME".txt; i=$((i+1)); done 
