#include <signal.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <unistd.h>

int START = 2000000000, END = 2000100000;
char *TESTS[] = {"brute_force", "brutish", "miller_rabin"};
// char *TESTS[] = {"miller_rabin", "brute_force", "brutish"};
int num_tests = sizeof(TESTS) / sizeof(char *);

int main(int argc, char *argv[])
{
  int proc_num = get_nprocs();
  printf("%d", proc_num);
  int testfds[num_tests][2];
  int resultfds[num_tests][2];
  int result, i;
  long n;
  pid_t pid;

  for (i = 0; i < num_tests; i++)
  {
    pipe(testfds[i]);
    pipe(resultfds[i]);

    pid = fork();

    if (pid == -1)
    {
      fprintf(stderr, "Failed to fork\n");
      exit(-1);
    }

    if (pid == 0)
    {
      // we are the child, connect the pipes correctly and exec!
      close(testfds[i][1]);
      close(resultfds[i][0]);
      dup2(testfds[i][0], STDIN_FILENO);
      dup2(resultfds[i][1], STDOUT_FILENO);
      execl("primality", "primality", TESTS[i], (char *)NULL);
    }

    // we are the parent
    close(testfds[i][0]);
    close(resultfds[i][1]);
  }

  // for each number, run each test
  for (n = START; n <= END; n++)
  {
    for (i = 0; i < num_tests; i++)
    {

      // we are the parent, so send test case to child and read results
      write(testfds[i][1], &n, sizeof(n));
      n += 1;
      write(testfds[i][1], &n + 1, sizeof(n));
      read(resultfds[i][0], &result, sizeof(result));
      read(resultfds[i][0], &result, sizeof(result));
      printf("%15s says %ld %s prime\n", TESTS[i], n, result ? "is" : "IS NOT");
    }
  }
}
