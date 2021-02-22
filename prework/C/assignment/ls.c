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

#include "ls.h"

int main(int argc, char **argv)
{
  int number_of_paths = get_number_of_paths(argv, argc);
  int i = 0;
  flags f = get_flags(argv, argc);

  // Skips the program name arg
  argv++;
  if (number_of_paths == 0)
  {
    print_dir(DEFAULT_DIR, f);
  }
  else
  {

    while (i < number_of_paths)
    {
      if ((*++argv)[0] == '-')
        continue;

      if (number_of_paths == 0)

        if (number_of_paths > 1)
          printf("%s:\n", *argv);
      if (i >= 1)
        printf("\n");

      print_dir(*argv, f);
      i++;
    }
  }
  return EXIT_SUCCESS;
}

// Declarations

int get_number_of_paths(char **argv, int argc)
{
  int count = 0;
  while (--argc > 0)
  {
    if ((*++argv)[0] != '-')
      count++;
  }
  return count;
}

int calculate_max_size_of_name(struct dirent **dirs, int size)
{
  int max, i;

  max = 0;
  i = 0;
  while (i < size)
  {
    int size = strlen(dirs[i]->d_name);
    if (size > max)
      max = size;
    i++;
  }
  return max;
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
  int i = 0;
  int n = scandir(path, &dirs, f.a == 0 ? filter_hidden_files : NULL, alphasort);
  int max_size_of_name = calculate_max_size_of_name(dirs, n) + 1;

  if (n == -1)
  {
    perror("scandir");
    exit(EXIT_FAILURE);
  }
  if (f.l == 0)
  {
    while (i < n)
    {
      if (i == n - 1)
      {
        printf("%s%s\n", dirs[i]->d_type == DT_DIR ? CYAN_COLOR : WHITE_COLOR, dirs[i]->d_name);
      }
      else
      {
        printf("Type: %d\n", dirs[i]->d_type);
        printf("%s%-*.*s", dirs[i]->d_type == DT_DIR ? CYAN_COLOR : WHITE_COLOR, max_size_of_name, max_size_of_name, dirs[i]->d_name);
      }
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
      printf("%s%s\n", dir->d_type == DT_DIR ? CYAN_COLOR : WHITE_COLOR, dir->d_name);

      free(dirs[i]);
      i++;
    }
  }
  free(dirs);
}