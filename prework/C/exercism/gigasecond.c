#include <time.h>

// Notes:
//  time_t type is an integral value representing the number of seconds elapsed since 00:00 hours, Jan 1, 1970 UTC.
//  it is NOT standarized
//  Alternative to 1000000000 is 1e9
//  More elaborate alternative: https://exercism.io/tracks/c/exercises/gigasecond/solutions/a3722e65a9944f3a8becce8e4f067e36

const long int GIGASECOND = 1000000000;

time_t gigasecond_after(time_t current)
{
  return current + GIGASECOND;
}
