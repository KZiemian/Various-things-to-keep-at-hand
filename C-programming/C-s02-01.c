#include <stdio.h>

int main() {
  int i = 0;
  char charVar = 'a';
  FILE *ptrFile = fopen("dane.txt", "r");

  while (feof(ptrFile) == 0) {
    charVar = getc(ptrFile);

    printf("%c, %d, feof(ptrFile) = %d, ferror(ptrFile) = %d\n", charVar, charVar, feof(ptrFile), ferror(ptrFile));
  }



  return 0;
}
