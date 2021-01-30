/* 
 * The C Programming Language
 *
 * Exercise 4-14
 *
 * Define a macro swap(t,x,y) that interchanges two arguments of type t. (Block structure will help.)
*/
#include <stdio.h>
#define swap(t, x, y) {t temp = x; x = y; y = temp;}

int main(void)
{
  int x = 10;
  int y = 20;
  printf("x: %d and y: %d\n",x, y);
  swap(int, x, y);
  printf("x: %d and y: %d\n",x, y);

  return 1;
}
