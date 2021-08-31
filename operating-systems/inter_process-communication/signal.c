#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <signal.h>

void sigint_handler(int sig)
{
  // it can't call printf() as it is not async safe
  write(0, "Ahhh! SIGINT!\n", 14);
}

int main(void)
{
  void sigint_handler(int sig);
  char s[200];
  struct sigaction sa;

  sa.sa_handler = sigint_handler;
  // sa.sa_flags = 0; in this case if the signal is send the child process fgets will be exit
  sa.sa_flags = SA_RESTART; // now the handler will restart
  sigemptyset(&sa.sa_mask);// block none

  if (sigaction(SIGINT, &sa, NULL) == -1)
  {
    perror("sigaction");
    exit(1);
  }

  printf("Enter a string:\n");

  if (fgets(s, sizeof s, stdin) == NULL)
    // fgets was interrupted by a signal, in this case SIGINT
    perror("fgets");
  else
    printf("You entered: %s\n", s);

  return 0;
}