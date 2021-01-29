/* 
 * The C Programming Language
 *
 * Exercise 3-1
 *
 * Our binary search makes two tests inside the loop, when one would suffice (at the price of more tests outside). 
 * Write a version with only one test inside the loop and measure the difference in run-time.
*/

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

int main(void)
{

  // Provided by stdlib.h
  return EXIT_SUCCESS;
}

int binsearch(int x, int v[], int n)
{
  int low, high, mid;
  low = 0;
  high = n - 1;

  while (low <= high)
  {
    mid = (low + high) / 2;

    if (x < v[mid])
      high = mid - 1;
    else
      low = mid + 1;
  }
  return v[low] == x ? low : -1;
}

int originaBinsearch(int x, int v[], int n)
{
  int low, high, mid;
  const NOT_FOUND = -1;

  low = 0;
  high = n - 1;

  while (low <= high)
  {
    mid = (low + high) / 2;

    if (x < v[mid])
      high = mid - 1;
    else if (x > v[mid])
      low = mid + 1;
    else
      return mid;
  }

  return NOT_FOUND;
}