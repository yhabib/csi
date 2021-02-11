/* 
 * ls - Yusef Habib
 *  
 *  1. Get all flags  √
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
#define MAX_SIZE_PATH 255

typedef struct
{
  int a;
  int h;
  int l;
} flags;

// Definitions
flags get_flags(char **, int);
void print_dir(char *, flags);
int get_number_of_paths(char **, int);

int main(int argc, char **argv)
{
  // Order of these two matters
  int number_of_paths = get_number_of_paths(argv, argc);
  flags f = get_flags(argv, argc);

  // printf("Total: %d\n", number_of_paths);
  while (--argc > 0 && (*++argv)[0] != '-')
  {
    if (argc < number_of_paths - 1)
      printf("\n");

    if (number_of_paths > 1)
      printf("%s:\n", *argv);

    print_dir(*argv, f);
  }
  return EXIT_SUCCESS;
}

// Declarations

int get_number_of_paths(char **argv, int argc)
{
  int count = 0;
  while (--argc > 0 && (*++argv)[0] != '-')
  {
    count++;
  }
  return count;
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
  // printf("{ a: %d, h: %d, l: %d }\n", f.a, f.h, f.l);

  return f;
}

int filter_hidden_files(const struct dirent *entry)
{
  return *entry->d_name != '.';
}

void print_dir(char *path, flags f)
{
  struct dirent *dir;
  struct dirent **dirs;
  int n, i;
  i = 0;
  printf("Analyzing: %s\n", path);
  n = scandir(path, &dirs, f.a == 0 ? filter_hidden_files : NULL, alphasort);
  if (n == -1)
  {
    perror("scandir");
    exit(EXIT_FAILURE);
  }

  if (f.l == 0)
  {
    while (i < n)
    {
      printf("%s ", dirs[i]->d_name);
      free(dirs[i]);
      i++;
    }
    printf("\n");
  }
  else
  {
    while (i < n)
    {
      char f_path[MAX_SIZE_PATH];
      struct stat *statbuf;
      printf("i: %d\n", i);

      strcpy(f_path, path);
      printf("name: %s\n", dirs[i]->d_name);
      strcat(f_path, dirs[i]->d_name);
      printf("f_path: %s\n", f_path);
      if (stat("./", statbuf))
      {
        perror("stat");
        exit(EXIT_FAILURE);
      }

      printf("File: %s ", dirs[i]->d_name);
      // printf("%s%.2f KB", NORMAL_COLOR, statbuf->st_size / 1000.0);
      // printf("%s%.2ld s", NORMAL_COLOR, statbuf->st_mtimespec.tv_sec);
      printf("%s%s\n", dir->d_type == DT_DIR ? DIR_COLOR : FILE_COLOR, dir->d_name);

      free(dirs[i]);
      i++;
    }
  }
  // else
  // {
  //   while ((dir = readdir(d)) != NULL)
  //   {
  //     struct stat *statbuf;
  //     char f_path[MAX_SIZE_PATH];

  //     strcpy(f_path, "");
  //     strcat(f_path, dir->d_name);

  //     int result = stat(f_path, statbuf);
  //     // Define max size per column
  //     // Rounding && apply unit based on size: B or KB
  //     printf("%s%.2f KB", NORMAL_COLOR, statbuf->st_size / 1000.0);
  //     printf("%s%.2ld s", NORMAL_COLOR, statbuf->st_mtimespec.tv_sec);
  //     printf("%s%s\n", dir->d_type == DT_DIR ? DIR_COLOR : FILE_COLOR, dir->d_name);
  free(dirs);
}