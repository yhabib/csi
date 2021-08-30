#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

// implement "ls | wc -l" in C.


int main(void)
{
  int pfds[2];

  // pfds[0] is for reading, pfds[1] is for writing
  pipe(pfds);

  if (!fork())
  // fork returned 0 -> child process
  {
    close(1);       /* close normal stdout as we need it for the pipe*/
    dup(pfds[1]);   /* duplicate the writing fd to be the first available one "1", stdout */
    close(pfds[0]); /* close reading fd */
    execlp("ls", "ls", "-a", NULL);
  }
  else
  // fork returned != 0 -> parent process
  {
    close(0);       
    dup(pfds[0]);   
    close(pfds[1]); 
    execlp("wc", "wc", "-l", NULL);
  }

  return 0;
}