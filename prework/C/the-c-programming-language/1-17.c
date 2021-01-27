/*
 * The C Programming Language, second edition,
 *
 * Exercise 1-17
 *
 * Write a program to print all input lines that are longer than 80 characters.
 *
 * 1. Get line
 * 2. If line bigger than max print
*/
#include <stdio.h>

#define MIN_SIZE 10

int getLine(char[], int);

int main(void)
{
  int len;
  char line[MIN_SIZE + 1];
  int isLongWord = 0;

  while ((len = getLine(line, MIN_SIZE)) > 0)
    if (line[len - 1] != '\n')
    {
      printf("%s", line);
      isLongWord = 1;
    }
    else if (isLongWord)
    {
      printf("%s", line);
      isLongWord = 0;
    }

  return 0;
}

int getLine(char line[], int lim)
{
  int i, c;
  for (i = 0; i < lim && (c = getchar()) != '\n' && c != EOF; i++)
    line[i] = c;

  if (c == '\n')
  {
    line[i] = c;
    ++i;
  }
  line[i] = '\0';

  return i;
}