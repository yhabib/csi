/*
 * The C Programming Language, second edition,
 *
 * Exercise 1-12
 *
 * Write a program that prints its input one word per line.
 */

#include <stdio.h>

#define TRUE 1
#define FALSE 0

int main(void)
{
  int c, prevblank;

  prevblank = FALSE;
  while ((c = getchar()) != EOF)
  {
    if (c == ' ' || c == '\t' || c == '\n')
    {
      if (prevblank == FALSE)
        putchar('\n');
      prevblank = TRUE;
    }
    else
    {
      prevblank = FALSE;
      putchar(c);
    }
  }
  return 0;
}