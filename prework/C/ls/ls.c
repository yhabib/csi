/* 
 * ls - Yusef Habib
 *  
 *  1. Get all flags
 *  2. Get all paths
 *  3. Render:
 *    3.1 As space separated values per path and filtering .hidden files
 *    3.2 -a All files
 *    3.3 -l As list
 *    3.3 -h If list then print nicely: time, size, ...
 *  4. More stats: st_mode: permissions, sm_mtime: time, st_uid: , st_author: author
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

typedef struct
{
  int a;
  int h;
  int l;
} flags;

char **get_paths(char **, int);
flags get_flags(char **, int);
void print_dir(char *);

int main(int argc, char **argv)
{
  flags f = get_flags(argv, argc);
  char **paths = get_path(argv, argc);
  print_dir(*paths);

  return EXIT_SUCCESS;
}

char **get_paths(char **argv, int argc)
{
  // to find each path and then return as an array of paths
}

flags get_flags(char **argv, int argc)
{
  flags f = {0, 0, 0};
  int i = 0;
  while (++i < argc)
  {
    argv++;
    if (**argv == '-')
    {
      while ((*(*argv)++) != '\0')
      {
        switch (**argv)
        {
        case 'a':
          f.a = 1;
          break;
        case 'h':
          f.h = 1;
          break;
        case 'l':
          f.l = 1;
          break;
        }
      }
    }
  }
  printf("{ a: %d, h: %d, l: %d }\n", f.a, f.h, f.l);
  return f;
}

// With no flags should be just list of files(filtering .)
void print_dir(char *path)
{
  printf("%sSize\t\tLast modified\t\tName\n", NORMAL_COLOR);
  printf("%s--------------------------------------------------------------\n", NORMAL_COLOR);
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

      // Define max size per column
      // Rounding && apply unit based on size: B or KB
      printf("%s%.2f KB\t\t", NORMAL_COLOR, statbuf->st_size / 1000.0);
      // Format into date
      printf("%s%.2ld s\t", NORMAL_COLOR, statbuf->st_mtimespec.tv_sec);
      printf("\t%s%s\n", dir->d_type == DT_DIR ? DIR_COLOR : FILE_COLOR, dir->d_name);
    }
  }
  else
    perror("Something wrong happened opening the directory");

  closedir(d);
}