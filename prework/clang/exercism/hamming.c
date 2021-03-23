#include "hamming.h"

// Notes:
//  First iteration I was calculating the len of each string -> O(3)
//  But at the end of the loop pointers will let me know if they have the same length or not  -> O(n)

int compute(const char *lhs, const char *rhs)
{
  int count = 0;
  while (*lhs != '\0' && *rhs != '\0')
  {
    if (*lhs != *rhs)
    {
      count++;
    }
    lhs++;
    rhs++;
  }
  if (*lhs != *rhs)
    return -1;
  return count;
}
