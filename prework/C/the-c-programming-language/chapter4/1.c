/* 
 * The C Programming Language
 *
 * Exercise 4-1
 *
 * Write the function strrindex(s,t), which returns the position of the rightmost occurrence of t in s, 
 * or −1 if there is none.
*/

#include <stdio.h>
#include <stdlib.h>

int strrindex(char[], char[]);

int main(void)
{
  printf("4.1 strrindex \n");
  printf("Last match index: %d\n", strrindex("a hello world, hello", "hello"));
  printf("Last match index: %d\n", strrindex("a helloa world, hellow", "hello"));
  printf("Last match index: %d\n", strrindex(" not sure whot should be here", "not"));
  printf("Last match index: %d\n", strrindex(" not sure whot should be here", "what"));

  printf("\n");
  // Provided by stdlib.h
  return EXIT_SUCCESS;
}

int strrindex(char s[], char t[])
{
  int i, j, lastIndex;
  lastIndex = -1;
  for (i = 0; s[i] != '\0'; i++)
  {
    for (j = 0; s[i + j] == t[j]; j++)
    {
      if (t[j + 1] == '\0')
      {
        // Got to the end of t with all matches
        lastIndex = i;
      }
    }
  }
  return lastIndex;
}