#include <stdio.h>

void funTest(int arrayInt[5]) {
  arrayInt[3] = -3;
}

int main() {
  int i = 0;
  static int arrayInt[5] = { 0, 1, 2, 3, 4 };

  for (i = 0; i < 5; ++i) {
    printf("arrayInt[%d] = %d\n", i, arrayInt[i]);
  }

  printf("\n");

  for (i = 0; i < 5; ++i) {
    printf("(arrayInt + %d) = %d\n", i, *(arrayInt + i));
  }

  funTest(arrayInt);

  printf("\nPo powrocie z funkcji funTest\n\n");

  for (i = 0; i < 5; ++i) {
    printf("*(arrayInt + %i) = %d\n", i, *(arrayInt + i));
  }



  return 0;
}
