/* 
 * The C Programming Language
 *
 * Exercise 2-2
 *
 * Write a loop equivalent to the for loop above without using && or ¦¦
 * 
 * -------------------------------------------------------------------
 *   for (i=0; i<lim−1 && (c=getchar()) != ′\n′ && c != EOF; ++i)
 *     s[i] = c;
*/

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

#define LIM 10
#define NEW_LINE '\n'

int main(void)
{

  int i;
  char c;

  i = 0;
  while (i < LIM - 1)
  {
    c = getchar();
    if (c == NEW_LINE)
      i = LIM;
    if (c == EOF)
      i = LIM;;
    printf("s[i]=c");
  }

  return EXIT_SUCCESS;
}