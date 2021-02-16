#!/bin/bash

BUFFETT="Life is like a snowball. The important thing is finding wet snow and a really long hill."
ISSAY=$BUFFETT

# Change1: replace the first occurrence of 'snow' with 'foot'. 
ISSAY=${ISSAY[@]/snow/foot}
echo "$ISSAY"
# Change2: delete the second occurrence of 'snow'. 
ISSAY=${ISSAY[@]/snow/}
echo "$ISSAY"
# Change3: replace 'finding' with 'getting'. 
ISSAY=${ISSAY[@]/finding/getting}
echo "$ISSAY"
# Change4: delete all characters following 'wet'. 
# Tip: One way to implement Change4, if to find the index of 'w' in the word 'wet' and then use substring extraction.
# POSITION=$(index "$ISSAY" wet) doesn't work in MAC: https://stackoverflow.com/a/17615946/3364845
temp=${ISSAY%%wet*}
LEN=$((${#temp} + 3))
ISSAY=${ISSAY:0:$LEN}
echo "$ISSAY"