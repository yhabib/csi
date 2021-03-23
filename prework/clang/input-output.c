#include <stdio.h>


main()
{
  long nc;

  while (getchar() != EOF)
  {
    ++nc;
  }
  printf("%ld\n", nc);
}