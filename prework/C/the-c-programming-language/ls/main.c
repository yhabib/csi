/* 
 * ls 
 *
 * ls program
 * 1st iteration
 *  1. Read input: path if none then use default dir "./"
 *  2. Go to path: opendir
 *  3. Read path: readdir
 *  4. Show stats: stat
 *  5. Print error: perror
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

// opendir&readdir
#include <sys/types.h>
#include <dirent.h>

// stat
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>

// styles
#define NORMAL_COLOR "\x1B[0m"
#define DIR_COLOR "\x1b[36m"
#define FILE_COLOR "\x1b[35m"

#define DEFAULT_DIR "./"
void printDir(char *path);

int main(int argc, char **argv)
{
  printf("%sls program\n\n", NORMAL_COLOR);

  char *path = argc == 1 ? DEFAULT_DIR : argv[1];
  printf("%sflags\tsize\t\tname\n", NORMAL_COLOR);
  printf("%s-------------------------------\n", NORMAL_COLOR);
  printDir(path);

  return EXIT_SUCCESS;
}

void printDir(char *path)
{
  struct dirent *dir;
  DIR *d = opendir(path);

  if (d)
  {
    while ((dir = readdir(d)) != NULL)
    {
      struct stat *statbuf;
      char f_path[255];
      // Check if size of path is bigger in that case break
      strcpy(f_path, path);
      strcat(f_path, dir->d_name);
      int result = stat(f_path, statbuf);

      // Rounding && apply unit based on size: B or KB
      printf("%s%u\t", NORMAL_COLOR, statbuf->st_flags);
      printf("%s%.2f KB\t", NORMAL_COLOR, statbuf->st_size / 1000.0);
      printf("\t%s%s\n", dir->d_type == DT_DIR ? DIR_COLOR : FILE_COLOR, dir->d_name);
    }
  }
  else
    perror("Something wrong happened opening the directory");

  closedir(d);
}