/* 
 * The C Programming Language
 *
 * Exercise 2-3
 *
 * Write the function htoi(s), which converts a string of hexadecimal digits 
 * (including an optional 0x or 0X) into its equivalent integer value. 
 * The allowable digits are 0 through 9, a through f, and A through F.
*/

// https://stackoverflow.com/a/7403877/3364845

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

#define LIM 10
#define NEW_LINE '\n'

int htoi(char[]);

int main(void)
{
  printf("0x1F => 31 ?? %d\n", htoi("0x1F"));
  printf("0xF => 15 ?? %d\n", htoi("0xF"));
  printf("10 => 16 ?? %d\n", htoi("10"));
  printf("0x100 => 256 ?? %d\n", htoi("100"));
  return EXIT_SUCCESS;
}

int htoi(char s[])
{
  int c, i, n;
  
  n = 0;
  for (i = 0; (c = s[i]) != '\0'; i++)
  {
    n *= 16;
    // Checks that it does not start by 0x or 0X
    if (i == 0 && c == '0')
    {
      c = s[++i];
      if (c != 'x' && c != 'X')
        i--;
    }
    else if (c >= '0' && c <= '9')
      n += c - '0';
    else if (c >= 'a' && c <= 'f')
      n += 10 + (c - 'a');
    else if (c >= 'A' && c <= 'F')
      n += 10 + (c - 'A');
    else
      /* invalid input */
      return n;
  }
  
  return n;
}
