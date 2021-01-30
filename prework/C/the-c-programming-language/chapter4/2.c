/* 
 * The C Programming Language
 *
 * Exercise 4-2
 *
 * Extend atof to handle scientific notation of the form
*/

#include <stdio.h>
#include <ctype.h>

double atof(char[]);

int main(void)
{
  printf("4.2 Extend atof to handle scientific notation of the form \n");
  printf("123.45e−6 => %.20f\n", atof("123.45e-6"));
  printf("123.45e−10 => %.20f\n", atof("123.45e-10"));

  // Provided by stdlib.h
  return 1;
}

/* atof:  convert string s to double */
double atof(char s[])
{
  double val, power;
  int i, sign, esign, epower;

  for (i = 0; isspace(s[i]); i++) /* skip white space */
    ;

  sign = (s[i] == '-') ? -1 : 1;

  if (s[i] == '+' || s[i] == '-')
    i++;

  for (val = 0.0; isdigit(s[i]); i++)
    val = 10.0 * val + (s[i] - '0');

  if (s[i] == '.')
    i++;

  for (power = 1.0; isdigit(s[i]); i++)
  {
    val = 10.0 * val + (s[i] - '0');
    power *= 10.0;
  }

  val = sign * val / power;

  if (s[i] == 'e' || s[i] == 'E')
  {
    esign = (s[++i] == '-') ? -1 : 1;
    if (s[i] == '-' || s[i] == '+')
      i++;

    for (epower = 0; s[i] != '\0'; i++)
      epower = 10 * epower + (s[i] - '0');

    for (int k = 0; k < epower; k++)
      val = esign < 0 ? val / 10 : val * 10;
  }
  return val;
}