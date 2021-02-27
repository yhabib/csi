#ifndef ISOGRAM_H
#define ISOGRAM_H

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

bool is_isogram(const char phrase[]);

#endif


// #include "isogram.h"

// Notes
//  bitset is a 32 bit to check occurrences of each letter
//

bool is_isogram(const char phrase[])
{
  if (phrase == NULL)
    return false;

  int32_t bitset = 0;
  int code;

  while ((code = *phrase++) != '\0')
  {
    int normalized_capital_code = (code & 0xdf) - 'A';
    if (normalized_capital_code < 0 || normalized_capital_code > 25)
      continue;

    int32_t temp_bitset = bitset;
    bitset |= 1 << normalized_capital_code;
    if (bitset == temp_bitset)
      return false;
  }
  return true;
}
