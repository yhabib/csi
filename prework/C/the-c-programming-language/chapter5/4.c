/* 
 * The C Programming Language
 *
 * Exercise 5-4
 *
 * Write the function strend(s,t), which returns 1 if the string t
 * occurs at the end of the string s, and zero otherwise.
*/

#include <stdio.h>
#include <stdlib.h>

int streend(char *, char *);

int main(void)
{
  char a[] = "hello,";
  char b[] = "lo";
  char c[] = "o,";
  printf("5.4 strend(s,t) \n");
  printf("%d\n", streend(a, b));
  printf("%d\n", streend(a, c));

  return EXIT_SUCCESS;
}

int streend(char *s, char *t)
{
  while (*s != *t && *s != '\0')
    s++;
  while (*s++ == *t++ && *t != '\0')
    ;
  return (*s == '\0' && *t == '\0') ? 1 : 0;
}