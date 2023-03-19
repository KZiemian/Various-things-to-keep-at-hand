#include <stdio.h>

void funTest(float **fPPtr) {
  printf("\nPoczatek funkcji funTest.\n");
  printf("*fPPtr = %p, fPPtr = %p, &fPPtr = %p\n\n", *fPPtr, fPPtr, &fPPtr);

  *fPPtr = NULL;

  printf("\nKoniec funkcji funTest.\n");
  printf("*fPPtr = %p, fPPtr = %p, &fPPtr = %p\n\n", *fPPtr, fPPtr, &fPPtr);
}

int main() {
  float fVar = 0.0;
  float *fPtr = &fVar;

  printf("Poczatek funkcji main.\n");
  printf("fPtr = %p, &fPtr = %p\n\n", fPtr, &fPtr);

  funTest(&fPtr);

  printf("Koniec funkcji main.\n");
  printf("fPtr = %p, &fPtr = %p.\n", fPtr, &fPtr);



  return 0;
}
