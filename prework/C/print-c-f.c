#include <stdio.h>

/* print Celsius-Fahrenheit table
*/

#define LOWER -20
#define UPPER 50
#define STEP 10

main()
{
  float fahr, celsius;

  printf("-- Celsius to Fahrenheit --\n");
  for (celsius = LOWER; celsius <= UPPER; celsius += STEP)
  {
    fahr = (9.0 / 5.0) * celsius + 32.0;
    printf("%3.0f   %3.1f\n", celsius, fahr);
  }
}