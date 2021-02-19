/*
*
* An Armstrong number is a number that is the sum of its own digits each raised to the power of the number of digits.
* For example:
*  9 is an Armstrong number, because 9 = 9^1 = 9
*  10 is not an Armstrong number, because 10 != 1^2 + 0^2 = 1
*  153 is an Armstrong number, because: 153 = 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
*  154 is not an Armstrong number, because: 154 != 1^3 + 5^3 + 4^3 = 1 + 125 + 64 = 190
*  Write some code to determine whether a number is an Armstrong number.
*
*/

#include <stdlib.h>
#include <math.h>
#include <stdbool.h>

bool is_armstrong_number(int);

int main(void)
{
  printf("%d\n", is_armstrong_number(153));
  printf("%d\n", is_armstrong_number(9));

  return EXIT_SUCCESS;
}

bool is_armstrong_number(int candidate)
{
  int number_of_digits = 0;

  for (int i = candidate; i > 0; i /= 10)
    number_of_digits++;

  int sum = 0;
  for (int i = candidate; i > 0; i /= 10)
    sum += pow(i % 10, number_of_digits);

  return sum == candidate;
}

// Feedback:
// It's a best practice to try to avoid uninitialized variables. I would go even one step further and say that you should define variables as late as
// possible and in the narrowest possible scope. Often that turns unintended uses into errors (which is a good thing). It also reduces the mental burden
//  of the reader because they don't have to remember the type and value of a variable for several lines.