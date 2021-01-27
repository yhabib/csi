/*
 * The C Programming Language, second edition,
 *
 * Exercise 1-12, page 21
 *
 * Rewrite the temperature conversion program of Section 1.2 to use a function for conversion.
*/
#include <stdio.h>

#define LOWER -20
#define UPPER 140
#define STEP 20

double FtoC(double temp);
double CtoF(double temp);

int main(void)
{
  printf("-- Fahrenheit to Celsius --\n");
  for (double fahr = LOWER; fahr <= UPPER; fahr += STEP)
    printf("%3.0f   %3.1f\n", fahr, FtoC(fahr));

  printf("\n\n-- Celsius to Fahrenheit --\n");
  for (double celsius = LOWER; celsius <= UPPER; celsius += STEP)
    printf("%3.0f   %3.1f\n", celsius, CtoF(celsius));
}

double FtoC(double temp)
{
  return (5.0 / 9.0) * (temp - 32.0);
}

double CtoF(double temp)
{
  return (9.0 / 5.0) * temp + 32.0;
}
