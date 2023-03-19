#include <stdio.h>
#include <unistd.h>

int main() {
  int i = 0;
  char bufferChar[10] = "Karakorum";
  FILE *ptrFile = fopen("data_1.dat", "w+");

  if (ptrFile == NULL) {
    printf("Blad otwarcia pliku.\n");
  }

  fprintf(ptrFile, "Eureka...\n");
  fprintf(ptrFile, "Eureka, drugi raz...\n");

  fflush(ptrFile);

  printf("fileno(ptrFile) = %d\n", fileno(ptrFile));

  write(fileno(ptrFile), bufferChar, 7);
  fprintf(ptrFile, "\n");



  return 0;
}
