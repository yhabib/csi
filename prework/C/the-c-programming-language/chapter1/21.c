/*
 * The C Programming Language
 *
 * Exercise 1-21, page 34
 *
 * Write a program entab that replaces strings of blanks with the
 * minimum number of tabs and blanks to achieve the same spacing. Use
 * the same stops as for detab. When either a tab or a single blank
 * would suffice to reach a tab stop, which should be given preference?
*/

// Note that \t means "tabulate", or in other words to pad to the nearest multiple of N position, where N is usually 8.
// https://stackoverflow.com/questions/38720331/why-the-first-t-doesnt-work

#include <stdio.h>

#define TAB_SIZE 8

int main(void)
{

  int c, spaceCount, col;

  spaceCount = 0;
  col = 0;

  while ((c = getchar()) != EOF)
  {
    if (c == ' ')
    {
      ++spaceCount;
      ++col;
    }
    else
    {
      if (spaceCount == 1)
        putchar(' ');
      else
      {
        int numOfTabs = c / TAB_SIZE - (c - spaceCount) / TAB_SIZE;
        for (int i = 0; i < numOfTabs; ++i)
          putchar('\t');
        if (numOfTabs >= 1)
          spaceCount = TAB_SIZE - col % TAB_SIZE;
        for (int i = 0; i < spaceCount; ++i)
          putchar(' ');
      }
      putchar(c);
      spaceCount = 0;
      ++col;
    }
  }

  return 0;
}
