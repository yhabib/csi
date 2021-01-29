/* 
 * The C Programming Language
 *
 * Exercise 2-1
 *
 * Write a program to determine the ranges of char, short, int, and long variables, both signed 
 * and unsigned, by printing appropriate values from standard headers and by direct computation. 
 * Harder if you compute them: determine the ranges of the various floating-point types.
*/

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

int main(void)
{
  // Important format specifiers
  printf("CHAR: \n");
  printf("Size of char: %d bits\n", CHAR_BIT);
  printf("Maximum value of char: %d \n", CHAR_MAX);
  printf("Minimum value of char: %d \n", CHAR_MIN);
  printf("Maximum value of signed char: %d \n", SCHAR_MAX);
  printf("Minimum value of signed char: %d \n", SCHAR_MIN);
  printf("Maximum value of unsigned char: %d \n", UCHAR_MAX);
  // NO UCHAR_MIN as its zero

  printf("\nINT: \n");
  printf("Maximum value of int(signed): %d \n", INT_MAX);
  printf("Minimum value of int(signed): %d \n", INT_MIN);
  printf("Maximum value of unsigned int: %u \n", UINT_MAX);

  printf("\nLONG: \n");
  printf("Maximum value of long: %ld\n", LONG_MAX);
  printf("Minimum value of long: %ld\n", LONG_MIN);
  printf("Maximum value of unsigned long: %lu\n", ULONG_MAX);

  printf("\nSHORT:\n");
  printf("Maximum value of short: %d\n", SHRT_MAX);
  printf("Minimum value of short: %d\n", SHRT_MIN);
  printf("Maximum value of unsigned short: %u\n", USHRT_MAX);

  // Provided by stdlib.h
  return EXIT_SUCCESS;
}