#include <stdio.h>

int main() {
  int i = 0;
  float floatArray[10];


  for (i = 0; i < 10; ++i) {
    floatArray[i] = 0.0;
  }

  for (i = 0; i < 10; ++i) {
    printf("floatArray[%d] = %f\n", i, floatArray[i]);
  }

  printf("\n");


  *floatArray = 1.0;

  printf("floatArray[0] = %f\n\n", floatArray[0]);

  *(floatArray) = 1.5;

  printf("floatArray[0] = %f\n\n", floatArray[0]);

  *(floatArray + 3) = 0.25;

  for (i = 0; i < 10; ++i) {
    printf("floatArray[%d] = %f\n", i, floatArray[i]);
  }



  return 0;
}

