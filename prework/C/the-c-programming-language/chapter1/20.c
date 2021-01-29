/*
 * The C Programming Language
 *
 * Exercise 1-20
 *
 * Write a program detab that replaces tabs in the input with the proper
 * number of blanks to space to the next tab stop. Assume a fixed set of
 * tab stops, say every n columns. Should n be a variable or a symbolic
 * parameter?
*/
// stdin is line buffered so until we hit enter it keeps reading, in is then when does stdout

#include <stdio.h>

#define TAB_SIZE 8

int main(void)
{

  int c, col, size;
  col = 0;

  while ((c = getchar()) != EOF)
  {
    if (c == '\t')
    {
      size =  TAB_SIZE - col % TAB_SIZE;
      for (int i = 0; i < size; i++)
      {
        putchar(' ');
        ++col;
      }
    }
    else
    {
      putchar(c);
      ++col;
    }
  }
  return 0;
}
