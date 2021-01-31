/* 
 * The C Programming Language
 *
 * Exercise 5-3
 *
 * Write a pointer version of the function strcat that we showed in 
 * Chapter 2: strcat(s,t) copies the string t to the end of s.
*/

#include <stdio.h>
#include <stdlib.h>

// Even though I'm not loading <string.h> recognizes that a function strrcat exists and complains
char *strrcat(char *, char *);

int main(void)
{
  char a[] = "hello,";
  char b[] = " world!";
  printf("5.3 strcat(s,t) \n");
  // printf("%s", strrcat("hello", "hello")); It doens't work like this because it creates this variables in executable as constants
  printf("%s", strrcat(a, b));

  printf("\n");
  return EXIT_SUCCESS;
}

char *strrcat(char *s, char *t)
{
  char *original = s;
  while (*s)
    s++;
  while ((*s++ = *t++))
    ;

  return original;
}