#include <stdio.h>

float funTest(float x, float y) {
  return x * y;
}

int main() {
  float aFloat = 0.0;
  float bFloat = 6.0;
  float cFloat = 5.0;

  
  printf("aFloat = %f\n", aFloat);
  printf("funTest: %p\n\n", funTest);

  aFloat = (*funTest)(bFloat, cFloat);

  printf("aFloat = %f\n", aFloat);


  
  return 0;
}
