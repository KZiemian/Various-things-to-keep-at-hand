#include <stdio.h>
/* #include <stdlib.h> */
#include <unistd.h>

int main(int argc, char *argv[]) {
  int i = 0;
  int opt = 0;


  printf("argc = %i\n", argc);

  for (i = 0; i < argc; ++i) {
    printf("*argv[%i] = %s\n", i, argv[i]);
  }
  printf("\n");

  opt = getopt(argc, argv, "a:b:c");

  printf("opt = %i\n", opt);
  /* while ( */



  return 0;
}
