#include <stdio.h>

int main() {
  FILE *ptrFile = fopen("dane.txt", "r");
  char c = 'a';

  while (feof(ptrFile) != 1) {
    c = getc(ptrFile);

    printf("%c, %d, feof(ptrFile) = %d\n", c, c, feof(ptrFile));
  }



  return 0;
}
