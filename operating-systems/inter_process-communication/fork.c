#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

int main(void)
{
  // short
  pid_t pid;
  // this variable is NOT shared but rather copied from parent to any child
  int rv;

  switch (pid = fork())
  {
  case -1:
    perror("fork"); /* something went wrong */
    exit(1);        /* parent exits */

  case 0:
    printf(" CHILD: This is the child process!\n");
    printf(" CHILD: My PID is %d\n", getpid());
    printf(" CHILD: My parent's PID is %d\n", getppid());
    printf(" CHILD: Enter my exit status (make it small): ");
    scanf(" %d", &rv);
    printf(" CHILD: I'm outta here!\n");
    exit(rv); // exit and return value

  default:
    printf("PARENT: This is the parent process!\n");
    printf("PARENT: My PID is %d\n", getpid());
    printf("PARENT: My child's PID is %d\n", pid);
    printf("PARENT: I'm now waiting for my child to exit()...\n");
    wait(&rv);  // wait for child to exit and read value
    printf("PARENT: My child's exit status is: %d\n", WEXITSTATUS(rv)); // WEXITSTATUS macro to extract the returned value from
    printf("PARENT: I'm outta here!\n");
  }

  return 0;
}