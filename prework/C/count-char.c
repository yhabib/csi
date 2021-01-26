#include <stdio.h>

main()
{
  long nc;
  nc = 0;
  printf("%ld\n", nc);
  while (getchar() != EOF)
    ++nc;

  printf("%ld\n", nc);
}