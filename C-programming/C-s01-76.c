#include <stdio.h>

struct alpha {
  char line[100];
  float floatArray[100];
};

int main() {
  struct alpha kos;


  printf("sizeof(kos) = %ld\n", sizeof(kos));
  printf("sizeof(kos.floatArray) = %ld\n", sizeof(kos.floatArray));



  return 0;
}
