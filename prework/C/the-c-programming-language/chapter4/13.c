/* 
 * The C Programming Language
 *
 * Exercise 4-13
 *
 * Write a recursive version of the function reverse(s), which reverses the string s in place.
*/

#include <stdio.h>
#include <string.h>

void reverse(char[]);

int main(void)
{
  printf("4.13 Recursive reverse \n");
  char s[] = "hello";
  reverse(s);
  printf("%s\n", s);
  return 1;
}

void swap(char s[], int l, int r)
{
  char temp = s[l];
  s[l] = s[r];
  s[r] = temp;
}

void doReverse(char s[], int left)
{
  int len = strlen(s);
  if (left >= len / 2)
    return;

  doReverse(s, left + 1);
  swap(s, left, len - 1 - left);
}

void reverse(char s[])
{
  int left = 0;
  doReverse(s, left);
}