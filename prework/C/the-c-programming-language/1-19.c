/*
 * The C Programming Language, second edition,
 *
 * Exercise 1-19
 *
 * Write a function reverse(s) that reverses the character string s. 
 * Use it to write a program that reverses its input a line at a time.
 *
*/

#include <stdio.h>

#define MAX_LINE 1000

int size(char[]);
void reverse(char[]);
int getLine(char[], int);

int main(void)
{
  int l = 0;
  char line[MAX_LINE];

  while((l = getLine(line, MAX_LINE)) > 0) {
    reverse(line);
    printf("%s", line);
  }

  return 0;
}

void reverse(char s[])
{
  int sizeOfWord = size(s);
  for (int i = 0; i < sizeOfWord / 2; i++)
  {
    char temp = s[sizeOfWord - i];
    s[sizeOfWord - i] = s[i];
    s[i] = temp;
  }
}

int size(char s[])
{
  int i;
  for (i = 0; s[i] != '\0'; i++)
    ;
  return i - 1;
}

int getLine(char line[], int lim)
{
  int i, c;
  for (i = 0; i < lim && (c = getchar()) != '\n' && c != EOF; i++)
    line[i] = c;

  line[i] = '\0';

  return i;
}