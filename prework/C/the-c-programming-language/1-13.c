/* 
 * The C Programming Language
 *
 * Exercise 1-13, page 24
 *
 * Write a program to print a histogram of the lengths of words in it's
 * input. It is easy to draw a histogram with the bars horizontal; a
 * vertical orientation is more challenging.
*/

#include <stdio.h>

#define MAX_SIZE 10
#define IN 1
#define OUT 0

void renderHHistogram(int[]);

int main(void)
{
  // Without the +1 it prints always: [1]    4610 abort      ./a.out
  int ndigit[MAX_SIZE + 1];
  int c, state, size;

  printf("Exercise 1.13: \n");

  for (int i = 0; i <= MAX_SIZE; ++i)
    ndigit[i] = 0;

  state = OUT;
  while ((c = getchar()) != EOF)
  {
    if (c == ' ' || c == '\n' || c == '\t')
    {
      if (state == IN)
      {
        if (size <= MAX_SIZE)
          ++ndigit[size];
      }
      state = OUT;
    }
    else
    {
      if (state == OUT)
      {
        size = 0;
        state = IN;
      }
      ++size;
    }
  }

  renderHHistogram(ndigit);

  return 0;
}

void renderHHistogram(int data[])
{
  printf("\nHorizontal Histogram: \n\n");
  for (int i = 1; i <= MAX_SIZE; ++i)
  {
    if (i != MAX_SIZE)
      printf(" %2d: ", i);
    else
      printf(">%d: ", MAX_SIZE);

    for (int j = 0; j < data[i]; ++j)
      putchar('#');

    putchar('\n');
  }
  printf("\n");
}