/* 
 * The C Programming Language
 *
 * Exercise 5-1
 *
 * As written, getint treats a + or - not followed by a digit as a valid representation of 
 * zero. Fix it to push such a character back on the input.
*/

#include <stdlib.h>
#include <stdio.h>
#include "getint.h"

int main(void)
{
  int i, r;
  r = getint(&i);
  if (r > 0)
    printf("%d\n", i);
  else if (r == 0)
    printf("not a number\n");
  else if (r == EOF)
    printf("end of file\n");
  else
    printf("whot??\n");
  
  return EXIT_SUCCESS;
}