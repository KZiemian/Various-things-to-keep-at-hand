#include <stdio.h>
#include <string.h>

int main() {
  char bufferA[40] = "Instytucja";
  char bufferB[40] = "Adres";

  char *const ptrA = bufferA;
  char *const ptrB = bufferB;

  printf("ptrA: %s, ptrA: %p\n", ptrA, ptrA);
  printf("ptrB: %s, ptrB: %p\n\n", ptrB, ptrB);

  strcpy(bufferA, "Poczta");
  strcpy(bufferB, "ul. Witosa 14/2");

  printf("ptrA: %s, ptrA: %p\n", ptrA, ptrA);
  printf("ptrB: %s, ptrB: %p\n", ptrB, ptrB);

  

  return 0;
}
