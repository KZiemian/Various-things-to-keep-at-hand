#include <stdio.h>

int main(int argc, char *argv[]) {
  int i = 0;

  for (i = 0; i < argc; ++i) {
    printf("i = %i,  param = %s\n", i, argv[i]);
    printf("i = %i,  param = %s\n\n", i, *(argv + i));
  }



  return 0;
}
