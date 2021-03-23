#include "pangram.h"

// Notes:
//   It uses a 32 bit as a mask to map all seen letters to 1, where lower bit maps to 'a' and 25th bit to 'z'
//   It will return true if the mask has the lowest 26bits to 1 => 0x3ffffff
//   To find out the position on the mask of current letter takes advantage of ASCII structure
//   where only the five lower bits are relevant to know which letter is(case insensitive) => 'a' 'A' -> 1, 'b' 'B' -> 2, ....
//   Eg: 'a' == 1100001 & 31 == 1100001 & 11111 = 1
//       'A' == 1000001 & 31 == 1000001 & 11111 = 1
//       'f' == 1100110 & 31 == 1100110 & 11111 = 6
//       'F' == 1000110 & 31 == 1000110 & 11111 = 6
//

bool is_pangram(const char *sentence)
{
  if (sentence == NULL)
    return false;

  int32_t bit_mask = 0;
  while (*sentence != '\0')
  {
    int shift_amount = (*sentence & 31) - 1;
    if (shift_amount >= 0 && shift_amount <= 26)
    {
      bit_mask |= 1 << shift_amount;
    }

    sentence++;
  }

  return bit_mask == 0x3ffffff;
}
