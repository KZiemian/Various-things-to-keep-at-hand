#include <stdio.h>

void funTest(int arrayInt[5]) {
  int i = 0;
  
  printf("Poczatek funkcji funTest.\n");

  for (i = 0; i < 5; ++i) {
    printf("*arrayInt = %d\n", *arrayInt);
    ++arrayInt;
  }

  arrayInt += -2;
  *arrayInt = -3;  
}

int main() {
  int i = 0;
  static int arrayInt[5] = { 0, 1, 2, 3, 4 };

  for (i = 0; i < 5; ++i) {
    printf("*(arrayInt + %d) = %d\n", i, *(arrayInt + i));
  }

  funTest(arrayInt);

  printf("\n\nPo powrocie z funkcji funTest.\n\n");

  for (i = 0; i < 5; ++i) {
    printf("*(arrayInt + %d) = %d\n", i, arrayInt[i]);
  }
    
  


  return 0;
}
