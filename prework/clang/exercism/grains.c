#include "grains.h"
#include <stdio.h>

// Notes:
//  Initially I was doing: bs |= 1 << (index - 1); but this had undefined behavior
//  Two's complement: 8 - 00001000 -> -8 - 11110111 + 1 = 11111000
//  lu: is at least32bits  while llu is at least 64bits, both could be equally big
//  Long prefix can be also capitalized to make it more legible. Eg: 1LLU
//
// From reviewer: ~0 => The type of the literal 0 is int.
// The tilde inverts all bits of that int, resulting in an int where all bits are set.
// Since all modern platforms implement int as Two's Complement the value of that int is -1.
// The return statement implicitly converts that -1 to a uint64_t which results in the correct value.

// In C shifting an n-bit integer by n or more bits to the left invokes undefined behavior. So 1ULL << 64 is only
// valid on platforms where unsigned long long has more than 64 bits. But you could either use slightly
//
// Different operands for that shift, or you could split that one 64-bit shift into two smaller shifts.
//  return 18446744073709551615;
//  return 0xFFFFFFFFFFFFFFFF;
//  return (1ULL << 63) * 2 - 1;
//  return (1ULL << 32 << 32) - 1;
//  return (2ULL << 63) - 1;
//  return UINT64_MAX;
//  return ~UINT64_C(0);
//  return 2 * square(64) - 1;
//
// and because the return type is uint64_t
//  return -1;
//  return ~0ULL;

uint64_t square(uint8_t index)
{
  if (index < 1 || index > 64)
  {
    return 0;
  }
  return 0 | 1LLU << (index - 1);
}

uint64_t total(void)
{
  return ~0;
  // return (1LLU << 63) + (1LLU << 63) - 1;
}
