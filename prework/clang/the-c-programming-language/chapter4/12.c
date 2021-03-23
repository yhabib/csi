/* 
 * The C Programming Language
 *
 * Exercise 4-12
 *
 * Adapt the ideas of printd to write a recursive version of itoa; that is, convert an integer 
 * into a string by calling a recursive routine.
*/

#include <stdio.h>
#include <ctype.h>
#include <string.h>

void itoa(int, char[]);
// To make it work like this we need this extnernal variable to keep track of n recursively
int next = 0;

int main(void)
{
  printf("4.12 Recursive itoa \n");
  char s[3];
  itoa(123, s);
  printf("%s\n", s);
  return 1;
}

void itoa(int n, char s[])
{
  if (n <= 0)
    return;
  itoa(n / 10, s);
  s[next++] = (n % 10) + '0';
}

void reverse(char s[])
{
  for (int i = 0; i < strlen(s) / 2; i++)
  {
    int j = strlen(s) - i - 1;
    if (i != j)
    {
      char temp = s[i];
      s[i] = s[j];
      s[j] = temp;
    }
  }
}