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
// https://gist.github.com/RabaDabaDoba/145049536f815903c79944599c6f952a
#define WHITE_COLOR "\e[0;37m"
#define RED_COLOR "\e[0;31m"
#define CYAN_COLOR "\e[0;36m"

#define DEFAULT_DIR "."
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